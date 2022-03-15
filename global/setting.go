package global

import (
	"sai/pkg/options"
	"sai/store"
)

var  (
	TencenSmsSetting *options.SmsOptions

	KafkaConfig *options.KafkaOptions
	Store store.Factory
)
