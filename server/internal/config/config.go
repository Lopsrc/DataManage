package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	Env     string `yaml:"env" env-default:"local"`
	GRPC  struct {
		Port   string `yaml:"port" env-default:"8080"`
		Timeout time.Duration `yaml:"timeout" env-default:"10h"`
	} `yaml:"grpc"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var instance *Config
var once sync.Once

// GetConfig returns the global instance of Config.
// It initializes the instance the first time it is called,
// using the provided path to the configuration file.
// If the initialization fails, it panics with the error.
func GetConfig(pathConfig string) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(pathConfig, instance); err != nil {
			panic(err)
		}
	})
	return instance
}