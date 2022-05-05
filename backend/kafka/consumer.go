package kafka

import (
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"

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
		log.Printf("[Consumer] topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		log.Printf("value: %s\n", string(msg.Value))
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
		log.Fatal("NewConsumerGroup err: ", err)
	}
	defer cg.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler := consumerGroupHandler{}
		for {
			log.Println("Running: ConsumerWorker")
			err = cg.Consume(ctx, topics, handler)
			if err != nil {
				log.Println("Consume err: ", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()
}
