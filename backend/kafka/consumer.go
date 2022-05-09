package kafka

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	"der-ems/config"
)

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
