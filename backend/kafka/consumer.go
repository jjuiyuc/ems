package kafka

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/config"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/services"
)

type WeatherDetail struct {
	ValidDate   string  `json:"validDate" validate:"required"`
	Acpcpsfc    float32 `json:"acpcpsfc"`
	Capesfc     float32 `json:"capesfc"`
	Cpratavesfc string  `json:"cpratavesfc"`
	Cpratsfc    string  `json:"cpratsfc"`
	Crainavesfc string  `json:"crainavesfc"`
	Crainsfc    string  `json:"crainsfc"`
	Dlwrfsfc    string  `json:"dlwrfsfc"`
	Dpt2m       string  `json:"dpt2m"`
	Dswrfsfc    string  `json:"dswrfsfc"`
	Lftxsfc     string  `json:"lftxsfc"`
	Lhtflsfc    string  `json:"lhtflsfc"`
	No4lftxsfc  string  `json:"no4lftxsfc"`
	Prateavesfc string  `json:"prateavesfc"`
	Pratesfc    string  `json:"pratesfc"`
	Pressfc     string  `json:"pressfc"`
	Pwatclm     string  `json:"pwatclm"`
	Rh2m        string  `json:"rh2m"`
	Shtflsfc    string  `json:"shtflsfc"`
	Lcdclcll    string  `json:"lcdclcll"`
	Mcdcmcll    string  `json:"mcdcmcll"`
	Tmpsfc      string  `json:"tmpsfc"`
	Ulwrfsfc    string  `json:"ulwrfsfc"`
	Ulwrftoa    string  `json:"ulwrftoa"`
	Uswrfsfc    string  `json:"uswrfsfc"`
	Uswrftoa    string  `json:"uswrftoa"`
}

type LatestWeather struct {
	Lat  float32         `json:"lat" validate:"required"`
	Lng  float32         `json:"lng" validate:"required"`
	Alt  float32         `json:"alt"`
	List []WeatherDetail `json:"values"`
}

type DBWeatherDataField struct {
	Acpcpsfc    float32 `json:"acpcpsfc"`
	Capesfc     float32 `json:"capesfc"`
	Cpratavesfc string  `json:"cpratavesfc"`
	Cpratsfc    string  `json:"cpratsfc"`
	Crainavesfc string  `json:"crainavesfc"`
	Crainsfc    string  `json:"crainsfc"`
	Dlwrfsfc    string  `json:"dlwrfsfc"`
	Dpt2m       string  `json:"dpt2m"`
	Dswrfsfc    string  `json:"dswrfsfc"`
	Lftxsfc     string  `json:"lftxsfc"`
	Lhtflsfc    string  `json:"lhtflsfc"`
	No4lftxsfc  string  `json:"no4lftxsfc"`
	Prateavesfc string  `json:"prateavesfc"`
	Pratesfc    string  `json:"pratesfc"`
	Pressfc     string  `json:"pressfc"`
	Pwatclm     string  `json:"pwatclm"`
	Rh2m        string  `json:"rh2m"`
	Shtflsfc    string  `json:"shtflsfc"`
	Lcdclcll    string  `json:"lcdclcll"`
	Mcdcmcll    string  `json:"mcdcmcll"`
	Tmpsfc      string  `json:"tmpsfc"`
	Ulwrfsfc    string  `json:"ulwrfsfc"`
	Ulwrftoa    string  `json:"ulwrftoa"`
	Uswrfsfc    string  `json:"uswrfsfc"`
	Uswrftoa    string  `json:"uswrftoa"`
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
	config := config.GetConfig()

	for msg := range claim.Messages() {
		log.WithFields(log.Fields{
			"topic":     msg.Topic,
			"partition": msg.Partition,
			"offset":    msg.Offset,
			"value":     string(msg.Value),
		}).Info("consuming")

		if msg.Topic == config.GetString("kafka.topic.receiveWeatherData") {
			ProcessWeatherData(msg.Value)
		}

		// Mark a message as consumed
		sess.MarkMessage(msg, "")
	}
	return nil
}

func ConsumerWorker(topics []string, group string) {
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
			log.Info("running: ConsumerWorker")
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

	validate := validator.New()
	err = validate.Struct(latestWeather)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Error("err validate: ", err)
			return
		}
	}

	for i, weatherDetail := range latestWeather.List {
		dbWeatherDataField := DBWeatherDataField{
			Acpcpsfc:    weatherDetail.Acpcpsfc,
			Capesfc:     weatherDetail.Capesfc,
			Cpratavesfc: weatherDetail.Cpratavesfc,
			Cpratsfc:    weatherDetail.Cpratsfc,
			Crainavesfc: weatherDetail.Crainavesfc,
			Crainsfc:    weatherDetail.Crainsfc,
			Dlwrfsfc:    weatherDetail.Dlwrfsfc,
			Dpt2m:       weatherDetail.Dpt2m,
			Dswrfsfc:    weatherDetail.Dswrfsfc,
			Lftxsfc:     weatherDetail.Lftxsfc,
			Lhtflsfc:    weatherDetail.Lhtflsfc,
			No4lftxsfc:  weatherDetail.No4lftxsfc,
			Prateavesfc: weatherDetail.Prateavesfc,
			Pratesfc:    weatherDetail.Pratesfc,
			Pressfc:     weatherDetail.Pressfc,
			Pwatclm:     weatherDetail.Pwatclm,
			Rh2m:        weatherDetail.Rh2m,
			Shtflsfc:    weatherDetail.Shtflsfc,
			Lcdclcll:    weatherDetail.Lcdclcll,
			Mcdcmcll:    weatherDetail.Mcdcmcll,
			Tmpsfc:      weatherDetail.Tmpsfc,
			Ulwrfsfc:    weatherDetail.Ulwrfsfc,
			Ulwrftoa:    weatherDetail.Ulwrftoa,
			Uswrfsfc:    weatherDetail.Uswrfsfc,
			Uswrftoa:    weatherDetail.Uswrftoa,
		}
		dbWeatherDataFieldJson, _ := json.Marshal(dbWeatherDataField)

		t, _ := time.Parse(time.RFC3339, weatherDetail.ValidDate)

		weatherForecast := &deremsmodels.WeatherForecast{
			Lat:       latestWeather.Lat,
			LNG:       latestWeather.Lng,
			Alt:       null.NewFloat32(latestWeather.Alt, true),
			ValidDate: t,
			Data:      null.NewJSON(dbWeatherDataFieldJson, true),
		}

		log.WithFields(log.Fields{
			"i":            i,
			"Lat":          weatherForecast.Lat,
			"LNG":          weatherForecast.LNG,
			"Alt":          weatherForecast.Alt,
			"ValidDate":    weatherForecast.ValidDate,
			"string(Data)": string(weatherForecast.Data.JSON),
		}).Debug("upsert weatherForecast data")
		err := services.UpsertWeatherForecast(weatherForecast)
		if err != nil {
			log.Error("err UpsertWeatherForecast: ", err)
		}
	}
}
