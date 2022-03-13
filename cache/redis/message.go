package redis

import "time"

type message struct {
	client *redisStore
}

func (m *message) Mget(key []string) (map[string]string,error)  {
	res, err :=m.client.GetMany(key)
	return res,err

}

func (m *message) PutInRedis(keyValues map[string]string,expire time.Duration) error {
	err:=m.client.SetMany(keyValues,expire)
	return err

}
func NewMessage(ch *redisStore) *message {
	return &message{
		ch,
	}
}


