package mysql

import (
	"context"
	"github.com/marmotedu/errors"
	"gorm.io/gorm"
	"sai/model"
	"sai/pkg/code"
)

type messageTemplate struct {
	db *gorm.DB
}

func (m *messageTemplate) GetMessageTemplate(ctx context.Context, where map[string]interface{}) (*model.MessageTemplate, error) {
	messageTemplateModel:=model.MessageTemplate{}
	err:=m.db.Where(where).First(&messageTemplateModel).Error
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,errors.WithCode(code.ErrMessageTemplateNotFound,err.Error())
		}
		return nil, errors.WithCode(code.ErrDatabase,err.Error())

	}
	return &messageTemplateModel,nil
}

func NewMessageTemplate(ds *datastore) *messageTemplate  {
	return &messageTemplate{db: ds.db}
}


