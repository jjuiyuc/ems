package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/config"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/services"
)

type LatestWeather struct {
	Lat    float32                  `json:"lat"`
	Lng    float32                  `json:"lng"`
	Alt    float32                  `json:"alt"`
	Values []map[string]interface{} `json:"values"`
}

type consumerGroupHandler struct {
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.WithFields(log.Fields{
			"topic":     msg.Topic,
			"partition": msg.Partition,
			"offset":    msg.Offset,
			"value":     string(msg.Value),
		}).Info("consuming")

		if msg.Topic == ReceiveWeatherData {
			ProcessWeatherData(msg.Value)
		}

		// Mark a message as consumed
		sess.MarkMessage(msg, "")
	}
	return nil
}

func WeatherConsumerWorker(topics []string, group string) {
	config := config.GetConfig()
	// Init sarama config
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cg, err := sarama.NewConsumerGroup(
		[]string{config.GetString("kafka.broker")},
		group,
		saramaConfig)
	if err != nil {
		log.Fatal("err NewConsumerGroup: ", err)
	}
	defer cg.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler := consumerGroupHandler{}
		for {
			log.Info("running: WeatherConsumerWorker")
			err = cg.Consume(ctx, topics, handler)
			if err != nil {
				log.Error("err Consume: ", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()
}

func ProcessWeatherData(msg []byte) {
	log.Debug("processWeatherData")
	var latestWeather LatestWeather
	err := json.Unmarshal(msg, &latestWeather)
	if err != nil {
		log.Error("err json.Unmarshal: ", err)
		return
	}

	for i, data := range latestWeather.Values {
		const validDate = "validDate"
		var t time.Time

		for key, value := range data {
			if key == validDate {
				t, _ = time.Parse(time.RFC3339, fmt.Sprintf("%v", value))
				break
			}
		}
		delete(data, validDate)
		dataJson, _ := json.Marshal(data)

		weatherForecast := &deremsmodels.WeatherForecast{
			Lat:       latestWeather.Lat,
			Lng:       latestWeather.Lng,
			Alt:       null.NewFloat32(latestWeather.Alt, true),
			ValidDate: t,
			Data:      null.NewJSON(dataJson, true),
		}

		log.WithFields(log.Fields{
			"i":            i,
			"Lat":          weatherForecast.Lat,
			"Lng":          weatherForecast.Lng,
			"Alt":          weatherForecast.Alt,
			"ValidDate":    weatherForecast.ValidDate,
			"string(Data)": string(weatherForecast.Data.JSON),
		}).Debug("upsert weatherForecast data")
		err = services.UpsertWeatherForecast(weatherForecast)
		if err != nil {
			log.Error("err UpsertWeatherForecast: ", err)
		}
	}
}
