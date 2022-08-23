package apps

import (
	"encoding/json"
	"testing"
	"time"

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
	seedUtTopic   string
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

	s.seedUtTopic = kafka.ReceiveWeatherData
	s.repo = repo
	s.handler = handler

	// Truncate data
	_, err := db.Exec("TRUNCATE TABLE weather_forecast")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec("TRUNCATE TABLE gateway")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec("TRUNCATE TABLE customer")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	// Mock seedUtWeather data
	s.seedUtTime = time.Now().UTC()
	s.seedUtWeather = LatestWeather{
		Lat: 24.75,
		Lng: 121.75,
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
		(1,'00001','001',24.75,121.75),
		(2,'00001','002',24.75,121.75),
		(3,'00002','001',24.75,121.75);
	`)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	// Mock gateway table
	_, err = db.Exec(`
		INSERT INTO gateway (id,uuid,customer_id) VALUES
		(1,'U00001',1),
		(2,'U00002',1),
		(3,'U00003',2),
		(4,'U00004',3);
	`)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
}

func (s *WeatherWorkerSuite) TearDownSuite() {
	models.Close()
}

func (s *WeatherWorkerSuite) Test_01_SaveWeatherData() {
	const validDate = "validDate"

	// Modify seedUtWeather data
	// seedUtDataUpdated
	seedUtDataUpdated := LatestWeather{
		Lat: s.seedUtWeather.Lat,
		Lng: s.seedUtWeather.Lng,
		Alt: s.seedUtWeather.Alt,
	}
	for _, value := range s.seedUtWeather.Values {
		seedUtDataUpdated.Values = append(seedUtDataUpdated.Values, testutils.CopyMap(value))
	}
	for i, value := range seedUtDataUpdated.Values {
		switch i {
		case 0:
			value[validDate] = s.seedUtTime.Add(+15 * time.Minute).Format(time.RFC3339)
		case 1:
			value[validDate] = s.seedUtTime.Add(+30 * time.Minute).Format(time.RFC3339)
		}
	}
	// seedUtDataNoValidDate
	seedUtDataNoValidDate := LatestWeather{
		Lat: s.seedUtWeather.Lat,
		Lng: s.seedUtWeather.Lng,
		Alt: s.seedUtWeather.Alt,
	}
	for _, value := range s.seedUtWeather.Values {
		seedUtDataNoValidDate.Values = append(seedUtDataNoValidDate.Values, testutils.CopyMap(value))
	}
	for _, value := range seedUtDataNoValidDate.Values {
		delete(value, validDate)
	}

	tests := []struct {
		name string
		args LatestWeather
	}{
		{
			name: "saveWeatherData",
			args: s.seedUtWeather,
		},
		{
			name: "saveWeatherDataUpdated",
			args: seedUtDataUpdated,
		},
		{
			name: "saveWeatherDataNoValidDate",
			args: seedUtDataNoValidDate,
		},
		{
			name: "saveWeatherDataEmptyInput",
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.name)
		dataJSON, err := json.Marshal(tt.args)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		msg, err := testutils.GetMockConsumerMessage(s.T(), s.seedUtTopic, dataJSON)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		s.Equalf(s.seedUtTopic, msg.Topic, e.ErrNewMessageNotEqual.Error())

		switch tt.name {
		case "saveWeatherDataNoValidDate":
			_, _, err = s.handler.saveWeatherData(msg.Value)
			s.Equalf(e.ErrNewKeyNotExist(validDate).Error(), err.Error(), e.ErrNewMessageNotEqual.Error())
			continue
		case "saveWeatherDataEmptyInput":
			_, _, err := s.handler.saveWeatherData(nil)
			s.Equalf(e.ErrNewUnexpectedJSONInput.Error(), err.Error(), e.ErrNewMessageNotEqual.Error())
			continue
		}

		currentCount, err := s.repo.Weather.GetWeatherForecastCount()
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		_, _, err = s.handler.saveWeatherData(msg.Value)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		updatedCount, err := s.repo.Weather.GetWeatherForecastCount()
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		switch tt.name {
		case "saveWeatherData":
			s.Equalf(currentCount+2, updatedCount, e.ErrNewMessageNotEqual.Error())
		case "saveWeatherDataUpdated":
			s.Equalf(currentCount+1, updatedCount, e.ErrNewMessageNotEqual.Error())
		}
	}
}

func (s *WeatherWorkerSuite) Test_02_GenerateWeatherSendingInfo() {
	type args struct {
		Lat float32
		Lng float32
	}

	type response struct {
		WeatherData []byte
		UUIDs       []string
	}

	testData := LatestWeather{
		Lat: s.seedUtWeather.Lat,
		Lng: s.seedUtWeather.Lng,
		Alt: s.seedUtWeather.Alt,
	}
	for _, value := range s.seedUtWeather.Values {
		testData.Values = append(testData.Values, testutils.CopyMap(value))
	}
	for i, value := range testData.Values {
		switch i {
		case 0:
			value["validDate"] = s.seedUtTime.Add(+15 * time.Minute).Format(time.RFC3339)
		case 1:
			value["validDate"] = s.seedUtTime.Add(+30 * time.Minute).Format(time.RFC3339)
		}
	}
	testWeatherData, _ := json.Marshal(testData)
	testUUIDs := []string{"U00001", "U00002", "U00003", "U00004"}

	tt := struct {
		name   string
		args   args
		wantRv response
	}{
		name: "generateWeatherSendingInfo",
		args: args{
			Lat: s.seedUtWeather.Lat,
			Lng: s.seedUtWeather.Lng,
		},
		wantRv: response{
			WeatherData: testWeatherData,
			UUIDs:       testUUIDs,
		},
	}

	log.Info("test name: ", tt.name)
	weatherData, uuids, err := s.handler.generateWeatherSendingInfo(tt.args.Lat, tt.args.Lng)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(tt.wantRv.WeatherData, weatherData, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.UUIDs, uuids, e.ErrNewMessageNotEqual.Error())
}
