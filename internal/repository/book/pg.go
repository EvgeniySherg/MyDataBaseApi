package book

import (
	"BookApi/internal/models"
	"context"
	"database/sql"
)

// TODO: написать реализацию интерфейса с базой для book

type pg struct {
	db *sql.DB
}

func (p *pg) GetByID(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}
func (p *pg) FindBookById(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}
func (p *pg) DeleteBookById(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}
func (p *pg) UpdateBookById(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}
func (p *pg) CreateBook(ctx context.Context, b *models.Book) (*models.Book, error) {
	return nil, nil
}

func NewRepository(db *sql.DB) models.BookRepository {
	return &pg{
		db: db,
	}
}
