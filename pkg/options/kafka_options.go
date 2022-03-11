package options

type KafkaOptions struct {
	ReturnSuccess bool   `json:"return-success,omitempty" mapstructure:"return-success"`
	Host          string `json:"host,omitempty" mapstructure:"host"`
}

func NewKafkaOptions() *KafkaOptions {
	return &KafkaOptions{
		ReturnSuccess: true,
		Host:          "",
	}
}
