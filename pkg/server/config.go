package server

import (
	"github.com/gin-gonic/gin"
)


type Config struct {
	InsecureServing *InsecureServingInfo
	Mode            string
	Middlewares     []string
	Healthz         bool

}

type InsecureServingInfo struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},

	}
}
type CompletedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// New returns a new instance of GenericAPIServer from the given config.
func (c CompletedConfig) New() (*GenericAPIServer, error) {
	s := &GenericAPIServer{
		InsecureServingInfo: c.InsecureServing,
		mode:                c.Mode,
		healthz:             c.Healthz,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}
