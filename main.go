package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Alok-Raturi/go/event-application/config"
)

type application struct {
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
	}


	mux := http.NewServeMux()

	mux.Handle("POST /login", auth.login)
	mux.Handle("POST /register", auth.register)
	

	app.logger.Println("Server started on port", config.Port)
	err := http.ListenAndServe(":"+config.Port, nil)

	if err != nil {
		app.logger.Fatal(err)
	}
}
