package handlers

import (
	"BookApi/internal/models"
)

type BookHandler struct {
	repository models.BookRepository
}

func NewBookHandler(repository models.BookRepository) BookHandler {
	return BookHandler{repository: repository}
}