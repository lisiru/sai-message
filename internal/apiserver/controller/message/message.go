package message

import (
	"sai/cache"
	"sai/service"
	"sai/store"
)

type MessageController struct {
	service service.Service
}

func NewMessageController(store store.Factory,cache cache.Factory) *MessageController  {
	return &MessageController{
		service: service.NewService(store,cache),
	}
}


