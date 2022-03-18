package deduplication

import (
	"sai/cache"
	"sai/cache/redis"
	"sai/common"
	"sai/pkg/util"
)

type DeDuplicationService interface {
	deduplicationSingleKey(taskInfo common.TaskInfo, receiver string) string
}

type baseDeDuplication struct {
	DeDuplicationService
	redisClient cache.Factory
}

// 公共去重入口
func (b *baseDeDuplication) DeDuplication(param common.DeduplicationParam) {
	taskInfo := param.TaskInfo
	var inRedisValue map[string]string
	var filterReceiver []string
	var readyPutRedisReceiver []string
	keys := b.deduplicationAllKey(taskInfo) // 各个去重方式需要实现的方法，从而达到实现不同方式的去重
	inRedisValue, _ = b.redisClient.MessageCaches().Mget(keys)
	//从redis中获取key已存在的值
	for _, receiver := range taskInfo.Receiver {
		key := b.deduplicationSingleKey(taskInfo, receiver)
		redisValue := inRedisValue[key]
		if len(redisValue) > 0 && util.StringToInt(redisValue) >= param.CountNum {
			filterReceiver = append(filterReceiver, receiver)
		} else {
			readyPutRedisReceiver = append(readyPutRedisReceiver, receiver)
		}

	}
	b.putInRedis(readyPutRedisReceiver, inRedisValue, param)
	taskInfo.Receiver = readyPutRedisReceiver
}

// 批量保存到redis
func (b *baseDeDuplication) putInRedis(readyPutRedisReceiver []string, inRedisValue map[string]string, param common.DeduplicationParam) {
	var keyValues = make(map[string]string)
	for _, receiver := range readyPutRedisReceiver {
		key := b.deduplicationSingleKey(param.TaskInfo, receiver)
		if inRedisValue[key] != "" {
			keyValues[key] = util.IntToString(util.StringToInt(inRedisValue[key]) + 1)
		} else {
			keyValues[key] = util.IntToString(1)
		}
	}
	if len(keyValues) > 0 {
		_ = b.redisClient.MessageCaches().PutInRedis(keyValues, param.DeduplicationTime)
	}

}

// 获取当前发送内容的全部key
func (b *baseDeDuplication) deduplicationAllKey(info common.TaskInfo) []string {
	var result []string
	for _, val := range info.Receiver {
		key := b.deduplicationSingleKey(info, val)
		result = append(result, key)

	}
	return result

}

func NewBaseDeDuplication(deDuplicationType int) *baseDeDuplication {
	redisClient, _ := redis.NewRedisFactoryOr(nil)
	return &baseDeDuplication{
		redisClient:          redisClient,
		DeDuplicationService: SelectDeDuplicationService(deDuplicationType),
	}

}
