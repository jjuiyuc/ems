package kafka

import (
	"testing"

	"der-ems/config"
	"der-ems/models"
	"der-ems/testutils"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	deremsmodels "der-ems/models/der-ems"
)

type ConsumerSuite struct {
	suite.Suite
}

func Test_Consumer(t *testing.T) {
	suite.Run(t, &ConsumerSuite{})
}

func (s *ConsumerSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	models.Init()

	// truncate data
	models.GetDB().Exec("truncate table weather_forecast")
}

func (s *ConsumerSuite) Test_ReceiveWeatherData() {
	consumer := mocks.NewConsumer(s.T(), mocks.NewTestConfig())
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Error("err Close: ", err)
		}
	}()

	seedUtMsg := `{
		"lat":24.75,
		"lng":121.0,
		"alt":100,
		"values":[
			{
				"validDate":"2022-04-19T12:00:00Z",
				"acpcpsfc":0.125,
				"capesfc":1.0,
				"cpratavesfc":"6.960000064282212e-06",
				"cpratsfc":"2.752000000327825e-05",
				"crainavesfc":"1.0",
				"crainsfc":"0.0",
				"dlwrfsfc":"386.7000122070313",
				"dpt2m":"287.0",
				"dswrfsfc":"15.712000846862791",
				"lftxsfc":"10.458468437194824",
				"lhtflsfc":"85.36327362060547",
				"no4lftxsfc":"1.2618210315704346",
				"prateavesfc":"0.0002532000071369",
				"pratesfc":"8.560000424040481e-05",
				"pressfc":"100225.5234375",
				"pwatclm":"45.57017135620117",
				"rh2m":"88.5999984741211",
				"shtflsfc":"26.11602783203125",
				"lcdclcll":"100.0",
				"mcdcmcll":"100.0",
				"tmpsfc":"288.83892822265625",
				"ulwrfsfc":"402.679443359375",
				"ulwrftoa":"208.6212615966797",
				"uswrfsfc":"1.856000065803528",
				"uswrftoa":"314.656005859375"
			}
		]
	}`
	var seedUtPartition int32 = 0
	consumer.ExpectConsumePartition(ReceiveWeatherData, seedUtPartition, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte(seedUtMsg)})

	test, err := consumer.ConsumePartition(ReceiveWeatherData, seedUtPartition, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("err ConsumePartition: ", err)
	}
	testMsg := <-test.Messages()
	s.Equal(ReceiveWeatherData, testMsg.Topic)
	s.Equal(seedUtPartition, testMsg.Partition)
	s.Equal(seedUtMsg, string(testMsg.Value))

	log.WithFields(log.Fields{
		"topic":     testMsg.Topic,
		"partition": testMsg.Partition,
		"offset":    testMsg.Offset,
		"value":     string(testMsg.Value),
	}).Info("consuming")

	ProcessWeatherData(testMsg.Value)
	count, _ := deremsmodels.WeatherForecasts().Count(models.GetDB())
	s.Equal(1, int(count))
}
