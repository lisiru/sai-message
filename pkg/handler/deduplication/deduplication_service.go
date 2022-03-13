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


func (b *baseDeDuplication) DeDuplication(param common.DeduplicationParam) {
	taskInfo := param.TaskInfo
	var inRedisValue map[string]string
	var filterReceiver []string
	var readyPutRedisReceiver []string
	keys := b.deduplicationAllKey(taskInfo)
	inRedisValue, _ = b.redisClient.MessageCaches().Mget(keys)
	//从redis中获取key已存在的值
	for _, receiver := range taskInfo.Receiver {
		key := b.deduplicationSingleKey(taskInfo, receiver)
		redisValue := inRedisValue[key]
		if util.StringToInt(redisValue) >= param.CountNum {
			filterReceiver = append(filterReceiver, receiver)
		} else {
			readyPutRedisReceiver = append(readyPutRedisReceiver, receiver)
		}

	}
	b.putInRedis(readyPutRedisReceiver,inRedisValue,param)
	taskInfo.Receiver=readyPutRedisReceiver
}

func (b *baseDeDuplication) putInRedis(readyPutRedisReceiver []string,inRedisValue map[string]string,param common.DeduplicationParam)  {
	var keyValues = make(map[string]string)
	for _,receiver:=range readyPutRedisReceiver{
		key:=b.deduplicationSingleKey(param.TaskInfo,receiver)
		if inRedisValue[key] !="" {
			keyValues[key] = inRedisValue[key]
		}else {
			keyValues[key] = util.IntToString(1)
		}
	}
	if len(keyValues)>0 {
		_ = b.redisClient.MessageCaches().PutInRedis(keyValues, param.DeduplicationTime)
	}


}

func (b *baseDeDuplication) deduplicationAllKey(info common.TaskInfo) []string {
	result := make([]string, len(info.Receiver))
	for _, val := range info.Receiver {
		key := b.deduplicationSingleKey(info, val)
		result = append(result, key)

	}
	return result

}

func NewBaseDeDuplication(deDuplicationType int) *baseDeDuplication {
	redisClient,_:=redis.NewRedisFactoryOr(nil)
	return &baseDeDuplication{
		redisClient: redisClient,
		DeDuplicationService: SelectDeDuplicationService(deDuplicationType),
	}

}
