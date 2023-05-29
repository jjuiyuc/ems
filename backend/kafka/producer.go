package kafka

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/infra"
)

// SendDataToGateways godoc
func SendDataToGateways(cfg *viper.Viper, topic string, data []byte, gatewayUUIDs []string) {
	for _, uuid := range gatewayUUIDs {
		replacedTopic := strings.Replace(topic, "{gw-id}", uuid, 1)
		log.Debug("topic: ", replacedTopic)
		Produce(cfg, replacedTopic, string(data))
	}
}

// SendDataToAIServer godoc
func SendDataToAIServer(cfg *viper.Viper, topic string, data []byte) {
	log.Debug("topic: ", topic)
	Produce(cfg, topic, string(data))
}

// Produce godoc
func Produce(cfg *viper.Viper, topic, message string) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer([]string{cfg.GetString("kafka.broker")}, saramaConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "sarama.NewAsyncProducer",
			"err":       err,
		}).Fatal()
	}

	var (
		wg                        sync.WaitGroup
		enqueued, timeout, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for e := range producer.Errors() {
			log.WithFields(log.Fields{
				"caused-by": "producer.Errors",
				"err":       e.Err,
				"msg":       e.Msg,
			}).Error()
			errors++
		}
	}()

	msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}
	ctx, cancel := context.WithTimeout(infra.GetGracefulShutdownCtx(), time.Millisecond*10)
	select {
	case producer.Input() <- msg:
		enqueued++
	case <-ctx.Done():
		timeout++
	}
	cancel()

	// Done
	producer.AsyncClose()
	wg.Wait()
	log.WithFields(log.Fields{
		"enqueued": enqueued,
		"timeout":  timeout,
		"errors":   errors,
	}).Debug("producer done")
}
