package handlers

import (
	"BookApi/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetAllBooks(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	log.Println("get information about all book")
	return c.JSON(http.StatusOK, json.NewEncoder(c.Response()).Encode(models.BookData))
}
