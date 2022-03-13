package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"sai/common"
	"sai/pkg/handler/pending"
	"sai/pkg/logger"
	"sai/pkg/util"
	"strings"
	"sync"
)

type ConsumerGroupHandler struct {
	GroupId string
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var taskInfoList = []common.TaskInfo{}

	for msg := range claim.Messages() {
		json.Unmarshal(msg.Value, &taskInfoList)
		groupId:=util.GetGroupIdByTaskInfo(taskInfoList[0])
		if consumer.GroupId== groupId{
			for _,taskInfo:=range taskInfoList{
				pending.GetPool(groupId).Schedule(pending.HandlerMessage(taskInfo))
			}

		}
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
			err := cg.Consume(ctx, []string{topic}, ConsumerGroupHandler{GroupId: groupId})
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
