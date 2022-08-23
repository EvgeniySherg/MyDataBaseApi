package handlers

import (
	"BookApi/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Read

type HttpServer struct {
	router *echo.Echo
}

// сделать с остальными
func (bh *BookHandler) GetBook(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "incorrect id num for get book")
	}
	book, err := bh.repository.GetByID(c.Request().Context(), ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "incorrect id num fot get book")
	}
	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(book))
}

// Create
func (bh *BookHandler) CreateBook(c echo.Context) error {
	log.Println("user create new book")
	var book models.Book
	// take json with newBook Data from request
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		log.Println("User enter incorrect data for create")
		return c.String(http.StatusBadRequest, "incorrect book data")
	}
	newBookID := len(models.BookData) + 1
	book.Id = newBookID
	bk, err := bh.repository.CreateBook(c.Request().Context(), book)
	log.Println("new book create")
	return c.JSON(http.StatusCreated, json.NewEncoder(c.Response()).Encode(bk))
}

// Update
func (bh *BookHandler) UpdateBook(c echo.Context) error {
	var newBook models.Book
	bookId := c.Param("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		models.ShowErrorInLog(err)
		return c.String(http.StatusBadRequest, "incorrect id")
	}
	oldBook, ok := models.FindBookById(id)
	if !ok {
		return c.String(http.StatusNotFound, "book with this Id does not exist")
	}
	// take json with new information from request and update old bookData
	err = json.NewDecoder(c.Request().Body).Decode(&newBook) // take json from request
	if err != nil {
		log.Println("User enter incorrect json Data for update")
		models.ShowErrorInLog(err)
		return c.String(http.StatusBadRequest, "incorrect book data")
	}
	oldBook.Author = newBook.Author
	oldBook.Jenre = newBook.Jenre
	oldBook.Title = newBook.Title

	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(oldBook))
}

// Delete
func (bh *BookHandler) DeleteBook(c echo.Context) error {
	bookId := c.Param("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		models.ShowErrorInLog(err)
		return c.String(http.StatusBadRequest, "incorrect id num fot delete book")
	}
	idx, ok := models.FindBookIdxForDel(id)
	if !ok {
		return c.String(http.StatusNotFound, "book with this id not found")
	}
	if idx == 0 {
		books := &models.BookData
		*books = models.BookData[1:]
	} else {
		models.BookData = models.DeleteFromBookSlice(models.BookData, id)

	}
	return c.String(http.StatusOK, "book delete")
}

func (bh *BookHandler) GetAllBooks(c echo.Context) error {
	log.Println("get information about all book")
	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(models.BookData))
}
