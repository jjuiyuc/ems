package kafka

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SimpleConsumer godoc
type SimpleConsumer struct {
	ctx      context.Context
	name     string
	consumer sarama.ConsumerGroup
	hander   sarama.ConsumerGroupHandler
	topics   []string
}

// NewSimpleConsumer godoc
func NewSimpleConsumer(
	ctx context.Context,
	cfg *viper.Viper,
	name string,
	hander sarama.ConsumerGroupHandler,
	topics []string,
) (c *SimpleConsumer, err error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup(
		[]string{cfg.GetString("kafka.broker")},
		name,
		saramaConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "sarama.NewConsumerGroup",
			"err":       err,
		}).Fatal()
	}

	c = &SimpleConsumer{
		ctx:      ctx,
		name:     name,
		consumer: consumer,
		hander:   hander,
		topics:   topics,
	}
	return
}

// MainLoop godoc
func (w *SimpleConsumer) MainLoop() {
	defer w.consumer.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler := w.hander
		for {
			log.Info("running: consumer")

			// Check for shutdown signal
			select {
			case <-w.ctx.Done():
				return
			default:
			}

			err := w.consumer.Consume(w.ctx, w.topics, handler)
			if err != nil {
				log.WithFields(log.Fields{
					"caused-by": "consumer.Consume",
					"err":       err,
				}).Error()
			}

			if w.ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()
}
