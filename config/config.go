package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	XDif       int  `envconfig:"X_DIFF"`
	YDif       int  `envconfig:"Y_DIFF"`
	SaveScreen bool `envconfig:"SAVE_SCREEN"`
}

var (
	config Config
	once   sync.Once
)

func Get() Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			logrus.Fatal(err)
		}
	})
	return config
}
