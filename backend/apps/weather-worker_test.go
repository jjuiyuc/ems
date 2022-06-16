package apps

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/kafka"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/testutils"
)

type WeatherWorkerSuite struct {
	suite.Suite
	seedUtTime    time.Time
	seedUtWeather LatestWeather
	repo          *repository.Repository
	handler       weatherConsumerHandler
}

func Test_WeatherWorker(t *testing.T) {
	suite.Run(t, &WeatherWorkerSuite{})
}

func (s *WeatherWorkerSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	repo := repository.NewRepository(db)
	handler := weatherConsumerHandler{
		cfg:  cfg,
		repo: repo,
	}

	s.repo = repo
	s.handler = handler

	// Truncate data
	_, err := db.Exec("TRUNCATE TABLE weather_forecast")
	s.Require().NoError(err)
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.Require().NoError(err)
	_, err = db.Exec("TRUNCATE TABLE gateway")
	s.Require().NoError(err)
	_, err = db.Exec("TRUNCATE TABLE customer")
	s.Require().NoError(err)
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.Require().NoError(err)
	// Mock seedUtWeather data
	s.seedUtTime = time.Now().UTC()
	s.seedUtWeather = LatestWeather{
		Lat: 24.75,
		Lng: 121.0,
		Alt: 100,
	}
	seedUtValue1 := map[string]interface{}{
		"validDate":   s.seedUtTime.Format(time.RFC3339),
		"acpcpsfc":    0.125,
		"capesfc":     1.0,
		"cpratavesfc": 6.960000064282212e-06,
		"cpratsfc":    2.752000000327825e-05,
		"crainavesfc": 1.0,
		"crainsfc":    0.0,
		"dlwrfsfc":    386.7000122070313,
		"dpt2m":       287.0,
		"dswrfsfc":    15.712000846862791,
		"lftxsfc":     10.458468437194824,
		"lhtflsfc":    85.36327362060547,
		"no4lftxsfc":  1.2618210315704346,
		"prateavesfc": 0.0002532000071369,
		"pratesfc":    8.560000424040481e-05,
		"pressfc":     100225.5234375,
		"pwatclm":     45.57017135620117,
		"rh2m":        88.5999984741211,
		"shtflsfc":    26.11602783203125,
		"lcdclcll":    100.0,
		"mcdcmcll":    100.0,
		"tmpsfc":      288.83892822265625,
		"ulwrfsfc":    402.679443359375,
		"ulwrftoa":    208.6212615966797,
		"uswrfsfc":    1.856000065803528,
		"uswrftoa":    314.656005859375,
	}
	s.seedUtWeather.Values = append(s.seedUtWeather.Values, seedUtValue1)
	seedUtValue2 := make(map[string]interface{})
	for key, value := range seedUtValue1 {
		seedUtValue2[key] = value
	}
	seedUtValue2["validDate"] = s.seedUtTime.Add(+15 * time.Minute).Format(time.RFC3339)
	s.seedUtWeather.Values = append(s.seedUtWeather.Values, seedUtValue2)

	// Mock customer table
	_, err = db.Exec(`
		INSERT INTO customer (id,customer_number,field_number,weather_lat,weather_lng) VALUES
		(1,'A00001','00001',24.75,121),
		(2,'A00001','00002',24.75,121),
		(3,'B00001','00001',24.75,121);
	`)
	s.Require().NoError(err)

	// Mock gateway table
	_, err = db.Exec(`
		INSERT INTO gateway (id,uuid,customer_id) VALUES
		(1,'U00001',1),
		(2,'U00002',1),
		(3,'U00003',2),
		(4,'U00004',3);
	`)
	s.Require().NoError(err)
}

func (s *WeatherWorkerSuite) TearDownSuite() {
	models.Close()
}

func (s *WeatherWorkerSuite) Test_01_GetWeatherData() {
	// Use default seedUtWeather data
	seedUtWeatherJson, err := json.Marshal(s.seedUtWeather)
	s.Require().NoError(err)

	testMsg := s.getMockConsumerMessage(seedUtWeatherJson)

	s.handler.SaveWeatherData(testMsg.Value)
	count, err := s.repo.Weather.GetWeatherForecastCount()
	s.Require().NoError(err)
	s.Equal(2, int(count))
}

func (s *WeatherWorkerSuite) Test_02_GetUpdatedWeatherData() {
	// Modify seedUtWeather data
	for i, value := range s.seedUtWeather.Values {
		switch i {
		case 0:
			value["validDate"] = s.seedUtTime.Add(+15 * time.Minute).Format(time.RFC3339)
		case 1:
			value["validDate"] = s.seedUtTime.Add(+30 * time.Minute).Format(time.RFC3339)
		}
	}
	seedUtWeatherJson, err := json.Marshal(s.seedUtWeather)
	s.Require().NoError(err)

	testMsg := s.getMockConsumerMessage(seedUtWeatherJson)

	s.handler.SaveWeatherData(testMsg.Value)
	count, err := s.repo.Weather.GetWeatherForecastCount()
	s.Require().NoError(err)
	s.Equal(3, int(count))
}

func (s *WeatherWorkerSuite) Test_03_GenerateWeatherSendingInfo() {
	// Mock data
	testLat := s.seedUtWeather.Lat
	testLng := s.seedUtWeather.Lng
	testWeatherData, _ := json.Marshal(s.seedUtWeather)
	testUUIDs := []string{"U00001", "U00002", "U00003", "U00004"}

	weatherData, UUIDs, err := s.handler.GenerateWeatherSendingInfo(testLat, testLng)
	s.Require().NoError(err)
	s.Equal(testWeatherData, weatherData)
	s.Equal(testUUIDs, UUIDs)
}

func (s *WeatherWorkerSuite) Test_04_GetNoValidDateWeatherData() {
	// Modify seedUtWeather data
	const validDate = "validDate"
	for _, value := range s.seedUtWeather.Values {
		delete(value, validDate)
	}
	seedUtWeatherJson, _ := json.Marshal(s.seedUtWeather)
	testMsg := s.getMockConsumerMessage(seedUtWeatherJson)

	_, _, err := s.handler.SaveWeatherData(testMsg.Value)
	s.Equal(e.NewKeyNotExistError(validDate).Error(), err.Error())
}

func (s *WeatherWorkerSuite) getMockConsumerMessage(seedUtWeatherJson []byte) (testMsg *sarama.ConsumerMessage) {
	consumer := mocks.NewConsumer(s.T(), mocks.NewTestConfig())
	defer func() {
		if err := consumer.Close(); err != nil {
			log.WithFields(log.Fields{
				"caused-by": "consumer.Close",
				"err":       err,
			}).Error()
		}
	}()

	var seedUtPartition int32 = 0
	consumer.ExpectConsumePartition(kafka.ReceiveWeatherData, seedUtPartition, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte(seedUtWeatherJson)})

	test, err := consumer.ConsumePartition(kafka.ReceiveWeatherData, seedUtPartition, sarama.OffsetOldest)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "consumer.ConsumePartition",
			"err":       err,
		}).Fatal()
	}
	testMsg = <-test.Messages()
	s.Equal(kafka.ReceiveWeatherData, testMsg.Topic)
	s.Equal(seedUtPartition, testMsg.Partition)
	s.Equal(string(seedUtWeatherJson), string(testMsg.Value))

	log.WithFields(log.Fields{
		"topic":     testMsg.Topic,
		"partition": testMsg.Partition,
		"offset":    testMsg.Offset,
		"value":     string(testMsg.Value),
	}).Info("consuming")
	return
}
