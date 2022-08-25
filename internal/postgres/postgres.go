package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgresConfig struct {
	Host         string `env:"DB_HOST"`
	Port         string `env:"DB_PORT"`
	User         string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	Sslmode      string `env:"DB_SSL_MODE"`
	DatabaseName string `env:"DATABASE_NAME"`
}

func InitDB(cnf *PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.DatabaseName, cnf.Sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Printf("connection to database not created")
		return nil, err
	}
	log.Printf("connection to database create successfully")
	return db, nil
}
