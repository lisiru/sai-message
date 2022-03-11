package service

import (
	"context"
	"sai/cache"
	"sai/common"
	"sai/store"
)

type MessageService interface {
	SendMessage(ctx context.Context,param common.SendRequestParam) error

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
	processContext:=common.ProcessContext{
		Code:          param.Code,
		SendTaskModel: common.SendTaskModel{
			MessageTemplateId: param.MessageTemplateId,
			MessageParamList: []common.MessageParam{param.MessageParam},
		},
	}
	processManager:=NewManager(m.store)
	err := processManager.Run(ctx,processContext)
	if err != nil {
		return err
	}
	return nil
}







