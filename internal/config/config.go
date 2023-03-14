package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     int64  `envconfig:"POSTGRES_PORT"`
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPass     string `envconfig:"POSTGRES_PASS"`
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
}

func New() (Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
