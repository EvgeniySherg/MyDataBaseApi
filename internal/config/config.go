package config

import (
	"BookApi/internal/postgres"
	"time"
)

type Config struct {
	Port string `env:"SERVER_PORT"`
	ReadTimeout time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"`
	DBPostgres postgres.PostgresConfig
}

// TODO функция которая спарсит данные env в этот конфиг
func InitConfig() (*Config, error) {
	cfg := &Config{

	}
	return cfg, nil
}
