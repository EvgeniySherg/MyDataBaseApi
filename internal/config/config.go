package config

import (
	"BookApi/internal/postgres"
	"time"
)

type Config struct {
	Port string `envconfig:"SERVERPORT"`
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	DBPostgres postgres.PostgresConfig
}

TODO функция которая спарсит данные env в этот конфиг