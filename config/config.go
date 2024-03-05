package config

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type Config struct {
}

var c = &Config{}
var once sync.Once

func Get() *Config {

	once.Do(func() {

		envconfig.MustProcess("", c)
	})
	return c
}
