package mysql

import (
	"context"
	"gorm.io/gorm"
	"sai/model"
)

type smsRecord struct {
	db *gorm.DB
}

func newSmsRecord(ds *datastore) *smsRecord {
	return &smsRecord{db: ds.db}
}

func (sms *smsRecord) CreateRecord(ctx context.Context, record *model.SmsRecord) error {
	return sms.db.Create(record).Error
}
