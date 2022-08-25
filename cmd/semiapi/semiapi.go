package main

import (
	"BookApi/internal/config"
	"BookApi/internal/handlers"
	"BookApi/internal/postgres"
	bookRep "BookApi/internal/repository/book"
	"context"
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// TODO: сделать init bookHandler

func initHandlers(app *echo.Echo, db *sql.DB) {
	bookShelterPrefix := "/bookshelter/"
	bookShelterGroup := app.Group(bookShelterPrefix)
	bookRep := bookRep.NewRepository(db)
	bookHan := handlers.NewBookHandler(bookRep)
	bookShelterGroup.PUT("/update/:id", bookHan.UpdateBook)
	bookShelterGroup.POST("/create", bookHan.CreateBook)
	bookShelterGroup.DELETE("/delete/:id", bookHan.DeleteBook)
	bookShelterGroup.GET("/book/:id", bookHan.GetBook)
	app.PUT("/update/:id", bookHan.UpdateBook)    // не работает
	app.POST("/", bookHan.CreateBook)             // работет
	app.DELETE("/delete/:id", bookHan.DeleteBook) // работает
	app.GET("/book/:id", bookHan.GetBook)         // работает

}

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}
	db, err := postgres.InitDB(&cnf.DBPostgres)
	if err != nil {
		log.Fatalf("failed db config: %v", err)
	}
	app := echo.New()
	initHandlers(app, db)

	httpServer := &http.Server{
		Addr:         cnf.Port,
		Handler:      app,
		ReadTimeout:  cnf.ReadTimeout,
		WriteTimeout: cnf.WriteTimeout,
	}

	go func() {
		if err := app.StartServer(httpServer); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	shutdownCtx, forceShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer forceShutdown()

	if err = httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown http server err: %v", err)
	}
}
