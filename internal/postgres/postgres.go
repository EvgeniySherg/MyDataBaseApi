package postgres

import (
	"BookApi/internal/config"
	"database/sql"
	"fmt"
)

// TODO: написать подключение к БД + структура конфиг

type PostgresConfig struct {
	Host         string `env:"DB_HOST"`
	Port         string `env:"DB_PORT"`
	User         string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	Sslmode      string `env:"DB_SSL_MODE"`
	DatabaseName string `env:"DATABASE_NAME"`
}

// TODO: перенести пакет в internal/postgres
func InitDB(cnf *config.Config) (*config.Config, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBPostgres.Host, cnf.DBPostgres.Port, cnf.DBPostgres.User, cnf.DBPostgres.Password, cnf.DBPostgres.DatabaseName, cnf.DBPostgres.Sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	cnf.DB = db
	return cnf, nil
}
