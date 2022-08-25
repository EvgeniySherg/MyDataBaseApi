package book

import (
	"BookApi/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
)

// TODO: написать реализацию интерфейса с базой для book

type Database struct {
	DB *sql.DB
}

var table = "book"

func (p *Database) GetByID(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	query := fmt.Sprintf("SELECT * FROM %s WHERE book_id=$1", table)
	rows, err := p.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Printf("database error %v \n", err)
		return book, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Genre, &book.Author)
		if err != nil {
			log.Println("database error")
			return book, err
		}
	}
	return book, nil
}

func (p *Database) DeleteBookById(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE book_id=$1", table)
	_, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Database) UpdateBookById(ctx context.Context, book *models.Book) error {
	query := fmt.Sprintf("UPDATE %s  SET title=$1, genre=$2, author=$3 WHERE book_id=$4", table)
	_, err := p.DB.ExecContext(ctx, query, book.Title, book.Genre, book.Author, book.Id)
	if err != nil {
		log.Println("database error")
		return err
	}
	return nil
}
func (p *Database) CreateBook(ctx context.Context, b *models.Book) error {
	query := fmt.Sprintf("INSERT INTO %s (title, genre, author) VALUES ($1, $2, $3)", table)
	_, err := p.DB.QueryContext(ctx, query, b.Title, b.Genre, b.Author)
	if err != nil {
		log.Printf("database error %v \n", err)
		return err
	}
	return nil
}

func (p *Database) FindBookById(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}

func NewRepository(db *sql.DB) models.BookRepository {
	return &Database{
		DB: db,
	}
}
