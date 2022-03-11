package redis

import (
	"github.com/go-redis/redis"
	"sai/cache"
	"context"
	"fmt"
	"time"
)

type users struct {
	client *redisStore
}

func genSendPhoneCodeKey(phone string) string  {
	return fmt.Sprintf(cache.UserCacheSendPhoneCodeKey,phone)
}

func (u *users) SetSendPhoneCodeCache(ctx context.Context, phone string, code string) error {
	key:=genSendPhoneCodeKey(phone)
	return u.client.Set(ctx,key,code,2000*time.Second)
}

func (u *users) GetSendPhoneCodeFromCache(ctx context.Context, phone string) (interface{}, error) {
	key:=genSendPhoneCodeKey(phone)
	return u.client.Get(ctx,key)
}

func (u *users) GetUserCache(ctx context.Context,uid int64) (interface{},error)  {
	key:=fmt.Sprintf(cache.UserCacheInfoKey,uid)
	return u.client.Get(ctx,key)
}

func (u *users) TestHash(ctx context.Context,key string,fields map[string]interface{}) (string,error)  {
	return u.client.HMset(ctx ,key,fields)
}
func (u *users) TestHset(ctx context.Context,key string,field string,values interface{}) (bool,error) {
	return u.client.HSet(ctx,key,field,values)
}

func (u *users) TestGetHash(ctx context.Context,key string,field string) (string,error)  {
	return u.client.HGet(ctx,key,field)
}
func (u *users) TestGetHAll(ctx context.Context,key string) (map[string]string,error)  {
	return u.client.HGetAll(key)
}

func (u *users) TestZadd(key string,members ...redis.Z) (int64,error) {
	return u.client.ZAdd(key,members...)
}

func (u *users) TestZRangeByScore(key string ,opt redis.ZRangeBy) ([]string,error) {
	return u.client.ZrangeByScore(key,opt)
}

func (u *users) Eval(script string,keys []string,args ...interface{}) (interface{},error)  {
	return u.client.Eval(script,keys,args...)
}

func NewUsers(ch *redisStore) *users {
	return &users{
		ch,
	}
}
