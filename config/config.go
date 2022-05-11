package config

import (
	"encoding/json"
	"fmt"
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

		//todo Не заполняется
		config.SaveScreen = true
		config.XDif = -1280

		config.print()
	})

	return config
}

func (cfg *Config) print() {
	jsonConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(string(jsonConfig))
}
