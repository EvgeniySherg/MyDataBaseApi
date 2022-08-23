package config

import (
	"BookApi/internal/postgres"
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

// InitConfig TODO: переписать на envconfig либу, перенести в internal/config
func InitConfig() (*postgres.Postgres, error) {
	var pc *postgres.Postgres               // создавать указатель на структуру Postgres и заполнять ее поля
	err := envconfig.Process("", pc.Config) // или указатель на структуру PostgresConfig ?
	if err != nil {
		return nil, fmt.Errorf("cannot load config: %v", err)
	}
	return pc, nil
}
