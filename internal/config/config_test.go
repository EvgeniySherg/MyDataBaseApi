package config

import (
	"BookApi/internal/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInitConfig(t *testing.T) {
	exp := &Config{
		Port:         ":8081",
		ReadTimeout:  time.Second*10,
		WriteTimeout: time.Second*10,
		DBPostgres:   postgres.PostgresConfig{
			Host:         "db host",
			Port:         "port",
			User:         "user",
			Password:     "pass",
			Sslmode:      "disable",
			DatabaseName: "testdb",
		},
	}
	res, err := InitConfig()
	assert.Nil(t, err)
	assert.Equal(t, exp, res)

}
