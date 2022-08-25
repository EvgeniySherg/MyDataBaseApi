package models

import "context"

type Book struct {
	Id     int    `json:"id" db:"book_id"`
	Title  string `json:"title" db:"title"`
	Genre  string `json:"genre" db:"genre"`
	Author string `json:"author" db:"author"`
}

// TODO: написать модель интерфейса для book репозитория

type BookRepository interface {
	GetByID(ctx context.Context, id int) (Book, error)
	FindBookById(ctx context.Context, id int) (*Book, error)
	DeleteBookById(ctx context.Context, id int) error
	UpdateBookById(ctx context.Context, book *Book) error
	CreateBook(ctx context.Context, book *Book) error
}
