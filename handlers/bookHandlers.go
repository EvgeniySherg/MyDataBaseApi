package handlers

import (
	"BookApi/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// realization of CRUD operation v.1.0

// Read
func GetBook(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	bookId := c.Param("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		models.ShowErrorInLog(err)
		return c.String(http.StatusBadRequest, "incorrect id num fot get book")
	}
	book, ok := models.FindBookById(id)
	if ok {
		json.NewEncoder(c.Response()).Encode(book)
		return c.String(200, "this is your book")
	} else {
		return c.String(404, "book not found, try another book id")
	}
}

// Create
func CreateBook(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	log.Println("user create new book")
	var book models.Book
	// take json with newBook Data from request
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		log.Println("User enter incorrect data for create")
		models.ShowErrorInLog(err)
		return c.String(http.StatusBadRequest, "incorrect book data")
	}
	newBookID := len(models.BookData) + 1
	book.Id = newBookID
	models.BookData = append(models.BookData, book)
	log.Println("new book create")
	return c.JSON(http.StatusCreated, json.NewEncoder(c.Response()).Encode(book))
}

// Update
func UpdateBook(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
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

	return c.JSON(200, json.NewEncoder(c.Response()).Encode(oldBook))
}

// Delete
func DeleteBook(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
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
