package service

import (
	"context"
	"sai/cache"
	"sai/common"
	"sai/pkg/util/kafka"
	"sai/store"
)

type MessageService interface {
	SendMessage(ctx context.Context,param common.SendRequestParam) error
	SendTest(ctx context.Context)

}

type messageService struct {
	store store.Factory
	cache cache.Factory
	manager Manager
}

func NewMessageService(s *service) *messageService  {
	return &messageService{
		store: s.store,
		cache: s.cache,
	}
}

func (m *messageService) SendMessage(ctx context.Context,param common.SendRequestParam) error {
	processContext:=&common.ProcessContext{
		Code:          param.Code,
		SendTaskModel: common.SendTaskModel{
			MessageTemplateId: param.MessageTemplateId,
			MessageParamList: []common.MessageParam{param.MessageParam},
		},
	}
	processManager:=NewManager(m.store)
	// 进入责任链模式对发送消息的一系列处理
	err := processManager.Run(ctx,processContext)
	if err != nil {
		return err
	}
	return nil
}

func (m *messageService) SendTest(ctx context.Context)  {
	kafka.NewProducer(ctx).Send("austin","lallalll")
}







