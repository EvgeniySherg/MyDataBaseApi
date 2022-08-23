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
	UpdateBookById(ctx context.Context, id int) error
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

func FindBookById(id int) (*Book, bool) {
	var found bool
	var book Book
	for i, b := range BookData {
		if b.Id == id {
			found = true
			return &BookData[i], found
		}
	}
	return &book, found
}

func FindBookIdxForDel(id int) (int, bool) {
	var found bool
	var bookidx int
	for idx, book := range BookData {
		if book.Id == id {
			found = true
			bookidx = idx
		}
	}
	return bookidx, found
}

func DeleteFromBookSlice(s []Book, idx int) []Book {
	copy(s[idx:], s[idx+1:])
	s[len(s)-1] = Book{}
	s = s[:len(s)-1]
	return s
}

type Message struct {
	Message string
}
