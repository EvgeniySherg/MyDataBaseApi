package semiapi

import (
	"BookApi/internal/handlers"
	"BookApi/internal/repository/book"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	db, err := initDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("failed db config: %v", err)
	}

	// TODO: сделать init bookHandler

	app := echo.New()
	initHandlers(app)

	httpServer := &http.Server{
		Addr:         cfg.Port,
		ReadTimeout:  cfg.HttpReadTimeout,
		WriteTimeout: cfg.HttpWriteTimeout,
		Handler:      app,
	}

	go func() {
		if err = app.StartServer(httpServer); err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down the http Server: %v", err)
		}
	}()

	// gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	shutdownCtx, forceShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer forceShutdown()

	if err = httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown http server err: %v", err)
	}
}

type Config struct {
	Port string `env:"PORT"`
	HttpReadTimeout time.Duration
	HttpWriteTimeout time.Duration
	DBConfig DBConfig
}

type DBConfig struct {
	Host string
	User string
	Pass string
	Port int
	Name string
}

// TODO: перенести пакет в internal/postgres
func initDB(dbCfg *DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Pass, dbCfg.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// TODO: переписать на envconfig либу, перенести в internal/config
func initConfig() (*Config, error)  {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("cannot load config: %v", err)
	}
	port = os.Getenv("semiport")

	cfg := &Config{}
	return cfg, nil
}

func initHandlers(app *echo.Echo, db *sql.DB) {
	bookShelterPrefix := "/bookshelter/"
	bookShelterGroup := app.Group(bookShelterPrefix)

	bookRepository := book.NewRepository(db)
	bookHandler := handlers.NewBookHandler(bookRepository)

	bookShelterGroup.PUT("/update/:id", handlers.UpdateBook)
	bookShelterGroup.POST("/create", handlers.CreateBook)
	bookShelterGroup.DELETE("/delete/:id", handlers.DeleteBook)
	bookShelterGroup.GET("/book/:id", bookHandler.GetBook)
	bookShelterGroup.GET("/books", handlers.GetAllBooks)
}