package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"strings"
	"time"
	"sai/global"
	"sai/pkg/logger"
)
var TopicName = "saiBusiness"
type Producer struct {
	AsyncProducer sarama.AsyncProducer
}
var kafkaConfig = global.KafkaConfig
func NewProducer(ctx context.Context) *Producer  {
	conf:=sarama.NewConfig()
	conf.Producer.RequiredAcks=sarama.WaitForAll
	conf.Producer.Return.Errors=true
	conf.Producer.Compression=sarama.CompressionZSTD
	conf.Producer.Flush.Messages=10
	conf.Producer.Flush.Frequency=500 *time.Millisecond
	producer,err:=sarama.NewAsyncProducer(strings.Split(kafkaConfig.Host,","),conf)
	if err != nil {
		logger.Errorf("failed to create produver:",err)
	}
	pd:=&Producer{AsyncProducer: producer}
	go func() {
		for  {
			select {
			case err:=<-producer.Errors():
				logger.Errorf("kafka producer send error %s",err.Err.Error())
			case <-ctx.Done(): // 接收主进程结束信号，持久化缓存数据，防止丢失，关闭producer链接
				producer.Close()
				logger.Info("quit:kafka producer")
				return


			}

		}

	}()
	return pd

}

func (p *Producer) Send(topic string,message string)  {
	p.AsyncProducer.Input() <-&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}


}
