package cache

import "time"

type MessageCache interface {
	Mget(keys []string) (map[string]string,error)
	PutInRedis(keyValues map[string]string,expire time.Duration) error
}
