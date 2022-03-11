package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
	"sai/pkg/logger"
)

type ConsumerGroupHandler struct {
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		session.MarkMessage(msg, "")
	}
	return nil

}

func ConsumerGroup(topic, groupId string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cg, err := sarama.NewConsumerGroup(strings.Split(kafkaConfig.Host, ","), groupId, config)
	if err != nil {
		logger.Errorf("NewConsumerGroup err: ", err)
	}
	defer cg.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := cg.Consume(ctx, []string{topic}, ConsumerGroupHandler{})
			if err != nil {
				logger.Errorf("Consume err: ", err)
			}
			if ctx.Err() != nil {
				return
			}

		}
	}()
	wg.Wait()

}
