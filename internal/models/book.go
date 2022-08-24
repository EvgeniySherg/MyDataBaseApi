package models

import "context"

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Jenre  string `json:"jenre"`
	Author Author `json:"author"`
}

// TODO: написать модель интерфейса для book репозитория

// все еще не понимаю context.Context
type BookRepository interface {
	GetByID(ctx context.Context, id int) (*Book, error)
	FindBookById(ctx context.Context, id int) (*Book, error)
	DeleteBookById(ctx context.Context, id int) error
	UpdateBookById(ctx context.Context, book *Book) error
	CreateBook(ctx context.Context, book *Book) error
}

type Author struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
}

var BookData []Book

// func init() {
// 	book1 := Book{
// 		1,
// 		"Boook",
// 		Author{
// 			"Ivan",
// 			"Ivanov",
// 		},
// 		"DeepDarkFantasy",
// 	}
// 	BookData = append(BookData, book1)
// }
