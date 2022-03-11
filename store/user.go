package store

import (
	"sai/model"
	"context"
)

type UserStore interface {
	Create(ctx context.Context, user *model.User ) error
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context,limit int64,offset int64) (*model.UserList, error)
	GetUserByWhere(ctx context.Context,where map[string]interface{}) (*model.User,error)
}
