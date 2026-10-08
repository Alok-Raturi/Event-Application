package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Alok-Raturi/go/event-application/config"
	"github.com/Alok-Raturi/go/event-application/internal/user"
	"github.com/Alok-Raturi/go/event-application/router"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

type application struct {
	logger *log.Logger
	db     *sql.DB
}

func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	db, err := initDB()
	if err != nil {
		logger.Fatal("Error initializing database:", err)
	}
	defer db.Close()
	app := &application{
		logger: logger,
		db:     db,
	}

	userRepo := user.NewUserRepository(db)
	userSvc := user.NewUserService(userRepo)
	userHandler := user.NewHandler(userSvc)

	mux := router.NewRouter(userHandler)

	app.logger.Println("Server started on port", config.Port)
	err = http.ListenAndServe(":"+config.Port, mux)

	if err != nil {
		app.logger.Fatal(err)
	}
}
