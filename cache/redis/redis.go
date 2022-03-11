package redis

import (
	"sai/cache"
	"sai/pkg/logger"
	genericoptions "sai/pkg/options"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

type redisStore struct {
	client redis.UniversalClient
}

func (ch *redisStore) UserCaches() cache.UserCache  {
	return NewUsers(ch)
}

var (
	cacheFactory cache.Factory
	once sync.Once
)

func NewRedisFactoryOr(opts *genericoptions.RedisOptions)  (cache.Factory,error) {
	logger.Debug("creating new Redis connection pool")
	if opts ==nil && cacheFactory ==nil{
		return nil,fmt.Errorf("failed to new redis cache factory")
	}
	var redisClient redis.UniversalClient
	once.Do(func() {
		poolSize:=500
		if opts.MaxActive>0 {
			poolSize=opts.MaxActive
		}

		timeout:=5*time.Second
		if opts.Timeout>0 {
			timeout=time.Duration(opts.Timeout)*time.Second
		}
		var tlsConfig *tls.Config
		if opts.UseSSL{
			tlsConfig=&tls.Config{
				InsecureSkipVerify: opts.SSLInsecureSkipVerify,
			}
		}
		options:=&RedisOption{
			Addrs: getRedisAddrs(opts),
			MasterName: opts.MasterName,
			Password: opts.Password,
			DB: opts.Database,
			DialTimeout: timeout,
			ReadTimeout: timeout,
			WriteTimeout: timeout,
			IdleTimeout: 240*timeout,
			PoolSize: poolSize,
			TLSConfig: tlsConfig,
		}
		logger.Info("--> [REDIS] create single-node")
		redisClient = redis.NewClient(options.simple())
		cacheFactory = &redisStore{redisClient}
	})
	return cacheFactory,nil

}

type RedisOption redis.UniversalOptions

func getRedisAddrs(opts *genericoptions.RedisOptions) (addrs []string)  {
	if len(opts.Addrs) !=0 {
		addrs = opts.Addrs
	}
	if len(addrs) ==0 &&opts.Port!=0 {
		addr:=opts.Host+":" + strconv.Itoa(opts.Port)
		addrs=append(addrs,addr)
	}
	return addrs
}

func (o *RedisOption) simple() *redis.Options  {

	addr:="127.0.0.1:6379"
	if len(o.Addrs)>0{
		addr = o.Addrs[0]
	}
	return &redis.Options{
		Addr: addr,
		OnConnect: o.OnConnect,
		DB: o.DB,
		Password: o.Password,
		MaxRetries: o.MaxRetries,
		MinRetryBackoff: o.MinRetryBackoff,
		MaxRetryBackoff: o.MaxRetryBackoff,
		DialTimeout: o.DialTimeout,
		ReadTimeout: o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		PoolSize: o.PoolSize,
		MinIdleConns: o.MinIdleConns,
		MaxConnAge: o.MaxConnAge,
		PoolTimeout: o.PoolTimeout,
		IdleTimeout: o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
		TLSConfig: o.TLSConfig,
	}
}
// GetObj 获取某个key对应的对象, 对象必须实现 https://pkg.go.dev/encoding#BinaryUnMarshaler
func (r *redisStore) GetObj(ctx context.Context, key string, model interface{}) error {
	cmd := r.client.Get(key)
	if errors.Is(cmd.Err(), redis.Nil) {
		return cache.ErrKeyNotFound
	}

	err := cmd.Scan(model)
	if err != nil {
		return err
	}
	return nil
}

// GetMany 获取某些key对应的值
func (r *redisStore) GetMany(ctx context.Context, keys []string) (map[string]string, error) {
	pipeline := r.client.Pipeline()
	vals := make(map[string]string)
	cmds := make([]*redis.StringCmd, 0, len(keys))

	for _, key := range keys {
		cmds = append(cmds, pipeline.Get(key))
	}

	_, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	errs := make([]string, 0, len(keys))
	for _, cmd := range cmds {
		val, err := cmd.Result()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		key := cmd.Args()[1].(string)
		vals[key] = val
	}
	return vals, nil
}

// Set 设置某个key和值到缓存，带超时时间
func (r *redisStore) Set(ctx context.Context, key string, val string, timeout time.Duration) error {
	return r.client.Set( key, val, timeout).Err()
}

// SetObj 设置某个key和对象到缓存, 对象必须实现 https://pkg.go.dev/encoding#BinaryMarshaler
func (r *redisStore) SetObj(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	return r.client.Set(key, val, timeout).Err()
}

// SetMany 设置多个key和值到缓存
func (r *redisStore) SetMany(ctx context.Context, data map[string]string, timeout time.Duration) error {
	pipline := r.client.Pipeline()
	cmds := make([]*redis.StatusCmd, 0, len(data))
	for k, v := range data {
		cmds = append(cmds, pipline.Set( k, v, timeout))
	}
	_, err := pipline.Exec()
	return err
}

// SetForever 设置某个key和值到缓存，不带超时时间
func (r *redisStore) SetForever(ctx context.Context, key string, val string) error {
	return r.client.Set(key, val, cache.NoneDuration).Err()
}

// SetForeverObj 设置某个key和对象到缓存，不带超时时间，对象必须实现 https://pkg.go.dev/encoding#BinaryMarshaler
func (r *redisStore) SetForeverObj(ctx context.Context, key string, val interface{}) error {
	return r.client.Set(key, val, cache.NoneDuration).Err()
}

// SetTTL 设置某个key的超时时间
func (r *redisStore) SetTTL(ctx context.Context, key string, timeout time.Duration) error {
	return r.client.Expire(key, timeout).Err()
}

// GetTTL 获取某个key的超时时间
func (r *redisStore) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(key).Result()
}

type rememberFunc func() (interface{},error)

func (r *redisStore) Remember(ctx context.Context, key string, timeout time.Duration, rememberFunc rememberFunc, obj interface{}) error {
	err := r.GetObj(ctx, key, obj)
	if err == nil {
		return nil
	}

	if !errors.Is(err, cache.ErrKeyNotFound) {
		return err
	}

	// key not found
	objNew, err := rememberFunc()
	if err != nil {
		return err
	}

	if err := r.SetObj(ctx, key, objNew, timeout); err != nil {
		return err
	}
	if err := r.GetObj(ctx, key, obj); err != nil {
		return err
	}
	return nil
}

func (r *redisStore) Calc(ctx context.Context, key string, step int64) (int64, error) {
	return r.client.IncrBy(key, step).Result()
}

func (r *redisStore) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.IncrBy(key, 1).Result()
}

func (r *redisStore) Decrement(ctx context.Context, key string) (int64, error) {
	return r.client.IncrBy(key, -1).Result()
}

func (r *redisStore) Del(ctx context.Context, key string) error {
	return r.client.Del(key).Err()
}

func (r *redisStore) DelMany(ctx context.Context, keys []string) error {
	pipline := r.client.Pipeline()
	cmds := make([]*redis.IntCmd, 0, len(keys))
	for _, key := range keys {
		cmds = append(cmds, pipline.Del( key))
	}
	_, err := pipline.Exec()
	return err
}

func (r *redisStore) Get(ctx context.Context,key string) (interface{},error) {
	return r.client.Get(key).Result()

}



func (r *redisStore) ZrangeByScore(key string,opt redis.ZRangeBy) ([]string,error) {
	return r.client.ZRangeByScore(key,opt).Result()
}

func (r *redisStore) HExists(ctx context.Context,key string,field string) (bool,error) {
	return r.client.HExists(key,field).Result()
}

func (r *redisStore) HGet(ctx context.Context,key string,field string) (string,error)  {
	return r.client.HGet(key,field).Result()
}

func (r *redisStore) HGetAll(key string) (map[string]string,error)  {
	return r.client.HGetAll(key).Result()
}

func (r *redisStore) HDel(ctx context.Context,key string,field ...string) (int64,error)   {
	return r.client.HDel(key,field...).Result()
}

func (r *redisStore) HSet(ctx context.Context,key string,field string,values interface{}) (bool,error) {
	return r.client.HSet(key,field,values).Result()
}

func (r *redisStore) HMset(ctx context.Context, key string, fields map[string]interface{}) (string, error) {
	return r.client.HMSet(key,fields).Result()
}

func (r *redisStore) Pipeline() redis.Pipeliner  {
	return r.client.Pipeline()
}

func (r *redisStore) ZRem(ctx context.Context,key string,members ...interface{}) (int64,error) {
	return r.client.ZRem(key,members).Result()
}

func (r *redisStore) ZAdd(key string,members ...redis.Z) (int64,error) {
	return r.client.ZAdd(key,members...).Result()
}

func (r *redisStore) Eval(script string,keys []string,args ...interface{}) (interface{},error)  {
	return r.client.Eval(script,keys,args).Result()
}






