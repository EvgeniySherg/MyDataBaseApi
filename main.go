/*
First trying to normal API and realization simple app
*/

package main

import (
	"BookApi/models"
	"BookApi/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

var (
	Prefix                  string = "/bookshelter"
	BookResourcePrefix      string = Prefix + "/book"
	ManyBooksResourcePrefix string = Prefix + "/books"
	port                    string
)

func ShowErrorInLog(err error) {
	log.Println("something wrong", err)
}

// connection to port with .env
func init() {
	err := godotenv.Load()
	if err != nil {
		models.ShowErrorInLog(err)
	}
	port = os.Getenv("port")
}

func main() {
	e := echo.New()
	log.Println("server connection")
	utils.BuildBookResource(BookResourcePrefix, e)
	utils.BuildManyBooksResources(ManyBooksResourcePrefix, e)
	e.Logger.Fatal(e.Start(port))
}
