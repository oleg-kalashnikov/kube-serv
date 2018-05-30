package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/oleg-kalashnikov/kube-serv/pkg/logger"
)

const (
	// ServiceName contains a service name prefix which used in ENV variables
	SERVICENAME = "KUBESERV"
)

// Config contains ENV variables
type Config struct {
	// Local service host
	LocalHost string `envconfig:"HOST"`
	// Local service port
	LocalPort int `envconfig:"PORT"`
	// Logging level in logger.Level notation
	LogLevel logger.Level `split_words:"true"`
}

// Load settles ENV variables into Config structure
func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
