package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Pgx   StorageConfig `envPrefix:"PGX_"`
	TgBot TgBotConfig
	Redis RedisConfig
}

type TgBotConfig struct {
	BotApiKey string `env:"TG_BOT_API_KEY"`
}

type RedisConfig struct {
	Password string `env:"REDIS_PASSWORD"`
}

type StorageConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	Database string `env:"DATABASE" envDefault:"app"`
	Username string `env:"USERNAME" envDefault:"app"`
	Password string `env:"PASSWORD" envDefault:"app"`
}

var cfg Config
var once sync.Once

func GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	return cfg
}
