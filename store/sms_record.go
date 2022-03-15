package store

import (
	"context"
	"sai/model"
)

type SmsRecord interface {
	CreateRecord(ctx context.Context,smsRecord *model.SmsRecord) error
}
