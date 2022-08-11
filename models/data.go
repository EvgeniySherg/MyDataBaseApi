package models

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Jenre  string `json:"jenre"`
}

type Author struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
}

var BookData []Book

func init() {
	book1 := Book{
		1,
		"Boook",
		Author{
			"Ivan",
			"Ivanov",
		},
		"DeepDarkFantasy",
	}
	BookData = append(BookData, book1)
}

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
