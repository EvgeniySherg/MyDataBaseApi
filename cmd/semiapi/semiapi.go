package semiapi

import (
	"BookApi/internal/config"
	"BookApi/internal/handlers"
	"BookApi/internal/postgres"
	bookRep "BookApi/internal/repository/book"
	"context"
	"database/sql"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// TODO: сделать init bookHandler

func initHandlers(app *echo.Echo, db *sql.DB) { //что мы ждем от инит хэндлера ? создание префикса, создание группы, перечисление основных хэндлеров
	bookShelterPrefix := "/bookshelter/"
	bookShelterGroup := app.Group(bookShelterPrefix)
	bookRep := bookRep.NewRepository(db)
	bookHan := handlers.NewBookHandler(bookRep)
	bookShelterGroup.PUT("/update/:id", bookHan.UpdateBook)
	bookShelterGroup.POST("/create", bookHan.CreateBook)
	bookShelterGroup.DELETE("/delete/:id", bookHan.DeleteBook)
	bookShelterGroup.GET("/book/:id", bookHan.GetBook)
}

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}
	cnfDb, err := postgres.InitDB(cnf)
	if err != nil {
		log.Fatalf("failed db config: %v", err)
	}
	app := echo.New()
	initHandlers(app, cnfDb.DB)

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
