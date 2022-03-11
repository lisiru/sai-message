package options

import (
	"sai/pkg/logger"
	genericoptions "sai/pkg/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server" mapstructure:"server"`
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	InsecuresServing *genericoptions.InsecureServerOptions `json:"insecure" mapstructure:"insecure"`
	Log *logger.Options `json:"log" mapstructure:"log"`
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`
	SmsOptions *genericoptions.SmsOptions `json:"sms" mapstructure:"sms"`
	KafkaOptions *genericoptions.KafkaOptions `json:"kafka" mapstructure:"kafka"`

}

func NewOptions() *Options  {
	o:=Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
		InsecuresServing: genericoptions.NewInsecureServerOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
		Log: logger.NewOptions(),
		SmsOptions: genericoptions.NewSmsOptions(),
		KafkaOptions: genericoptions.NewKafkaOptions(),

	}
	return &o
}
