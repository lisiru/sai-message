package store

import (
	"context"
	"sai/model"
)

type MessageTemplate interface {
	GetMessageTemplate(ctx context.Context,where map[string]interface{}) (*model.MessageTemplate,error)
}
