package options

import (
	"sai/pkg/server"
	"net"
	"strconv"
)

type InsecureServerOptions struct {
	BindAddress string 	`json:"bind-address" mapstructure:"bind-address"`
	BindPort int `json:"bind-port" mapstructure:"bind-port"`
}

func NewInsecureServerOptions() *InsecureServerOptions  {
	return &InsecureServerOptions{
		BindAddress: "127.0.0.1",
		BindPort: 8080,
	}
}

func (s *InsecureServerOptions) ApplyTo(c *server.Config) error  {
	c.InsecureServing = &server.InsecureServingInfo {
		Address:net.JoinHostPort(s.BindAddress,strconv.Itoa(s.BindPort)),
	}
	return nil
}
