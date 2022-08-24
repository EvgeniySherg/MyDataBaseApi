package config

import (
	"BookApi/internal/postgres"
	"database/sql"
	"time"
)

type Config struct {
	Port         string        `env:"SERVER_PORT"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"`
	DBPostgres   postgres.PostgresConfig
	DB           *sql.DB // подойдет ???
}

// TODO функция которая спарсит данные env в этот конфиг
func InitConfig() (*Config, error) {
	cfg := &Config{
		Port:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		DBPostgres: postgres.PostgresConfig{
			Host:         "",
			Port:         "5432",
			User:         "",
			Password:     "admin",
			Sslmode:      "disable",
			DatabaseName: "postgres",
		},
	}
	return cfg, nil
}
