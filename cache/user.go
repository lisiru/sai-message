package cache

import (
	"context"
	"github.com/go-redis/redis"
)

type UserCache interface {
	GetUserCache(ctx context.Context,uid int64) (interface{},error)
	GetSendPhoneCodeFromCache(ctx context.Context,phone string) (interface{},error)
	SetSendPhoneCodeCache(ctx context.Context,phone string,code string) error
	TestHash(ctx context.Context,key string,fields map[string]interface{}) (string,error)
	TestHset(ctx context.Context,key string,field string,values interface{}) (bool,error)
	TestGetHash(ctx context.Context,key string,field string) (string,error)
	TestGetHAll(ctx context.Context,key string) (map[string]string,error)
	TestZadd(key string,members ...redis.Z) (int64,error)
	TestZRangeByScore(key string ,opt redis.ZRangeBy) ([]string,error)
	Eval(script string,keys []string,args ...interface{}) (interface{},error)
}
