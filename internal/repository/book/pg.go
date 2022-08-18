package book

import (
	"BookApi/internal/models"
	"database/sql"
)

// TODO: написать реализацию интерфейса с базой для book

type pg struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) models.BookRepository {

}