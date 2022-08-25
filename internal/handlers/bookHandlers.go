package handlers

import (
	"BookApi/internal/models"
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

// TODO: везде прокинуть ошибку чтобы видеть ее + логирование

// сделать с остальными
func (bh *BookHandler) GetBook(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("cannot strconv.Atoi: %v", err)
		return c.String(http.StatusBadRequest, "incorrect id num for get book")
	}
	book, err := bh.repository.GetByID(c.Request().Context(), ID)
	if err != nil {
		log.Printf("repository.GetByID: %v", err)
		return c.String(http.StatusBadRequest, "database error, incorrect id")
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
	err = bh.repository.CreateBook(c.Request().Context(), &book)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	log.Println("new book create")
	return c.JSON(http.StatusCreated, "create")
}

// Update
func (bh *BookHandler) UpdateBook(c echo.Context) error {
	var newBook models.Book
	// take json with new information from request and update old bookData
	err := json.NewDecoder(c.Request().Body).Decode(&newBook) // take json from request
	if err != nil {
		log.Println("User enter incorrect json Data for update")
		return c.String(http.StatusBadRequest, "incorrect book data")
	}
	err = bh.repository.UpdateBookById(c.Request().Context(), &newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "book update")
}

// Delete
func (bh *BookHandler) DeleteBook(c echo.Context) error {
	bookId := c.Param("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		return c.String(http.StatusBadRequest, "incorrect id num for delete book")
	}
	err = bh.repository.DeleteBookById(c.Request().Context(), id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "error delete")
	}
	return c.JSON(http.StatusOK, "book delete")
}
