package store

import (
	"context"
	"sai/model"
)

type ActivityStore interface {
	Create(ctx context.Context,activity *model.Activity) error
	GetActivityByWhere(ctx context.Context,where map[string]interface{}) (*model.Activity,error)

}