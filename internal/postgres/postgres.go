package postgres

import (
	"database/sql"
	"fmt"
)

// TODO: написать подключение к БД + структура конфиг

type PostgresConfig struct {
	Host         string `envconfig:"HOST"`
	Port         string `envconfig:"PORT"`
	User         string `envconfig:"USER"`
	Password     string `envconfig:"PASSWORD"`
	Sslmode      string `envconfig:"SSLMODE"`
	DatabaseName string `envconfig:"DATABASENAME"`
}

type Postgres struct {
	Db     *sql.DB
	Config *PostgresConfig
}

// TODO: перенести пакет в internal/postgres
func InitDB(ps *Postgres) (*Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		ps.Config.Host, ps.Config.Port, ps.Config.User, ps.Config.Password, ps.Config.DatabaseName, ps.Config.Sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	ps.Db = db
	return ps, nil
}
