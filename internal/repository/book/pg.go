package book

import (
	"BookApi/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var ErrBookNotFound = errors.New("book not found")
var ErrBookNotCreated = errors.New("book not created")

type database struct {
	DB *sql.DB
}

var table = "book"

func (p *database) GetByID(ctx context.Context, id int) (*models.Book, error) {
	var book models.Book
	query := fmt.Sprintf("SELECT id, title, genre, author FROM %s WHERE id=$1", table)

	err := p.DB.QueryRowContext(ctx, query, id).Scan(&book.Id, &book.Title, &book.Genre, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return &book, nil
}

func (p *database) DeleteBookById(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE book_id=$1", table)

	res, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec err -> %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected while delete err -> %v", err)
	}
	if rows == 0 {
		return ErrBookNotFound
	}
	return nil
}

func (p *database) UpdateBookById(ctx context.Context, book *models.Book) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, genre=$2, author=$3 WHERE id=$4", table)

	res, err := p.DB.ExecContext(ctx, query, book.Title, book.Genre, book.Author, book.Id)
	if err != nil {
		return fmt.Errorf("exec err -> %v", err)
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {

		return errors.New(fmt.Sprintf("book not updated: %v", book.Title))
	}
	return nil
}

func (p *database) CreateBook(ctx context.Context, b *models.Book) error {
	var book *models.Book
	query := fmt.Sprintf("INSERT INTO %s (title, genre, author) VALUES ($1, $2, $3)", table)

	err := p.DB.QueryRowContext(ctx, query, b.Title, b.Genre, b.Author).Scan(&book.Id, &book.Title)
	switch {
	case err == sql.ErrNoRows:
		return ErrBookNotCreated
	case err != nil:
		return fmt.Errorf("create book err -> %v", err)
	default:
		log.Printf("book with id - %v,  created title %s\n", book.Id, book.Title)
	}
	return nil
}

func NewRepository(db *sql.DB) models.BookRepository {
	return &database{
		DB: db,
	}
}
