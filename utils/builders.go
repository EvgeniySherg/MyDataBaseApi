package utils

import (
	"BookApi/handlers"
	"github.com/labstack/echo/v4"
)

func BuildBookResource(prefix string, e *echo.Echo) {
	e.PUT(prefix+"/update/:id", handlers.UpdateBook)
	e.POST(prefix+"/create", handlers.CreateBook)
	e.GET(prefix+"/get/:id", handlers.GetBook)
	e.DELETE(prefix+"/delete/:id", handlers.DeleteBook)
}

func BuildManyBooksResources(prefix string, e *echo.Echo) {
	e.GET(prefix, handlers.GetAllBooks)
}
