package router

import (
	"net/http"

	"github.com/Alok-Raturi/go/event-application/internal/user"
)

func NewRouter(userHandler *user.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/users", userHandler.GetUsers)
	mux.HandleFunc("POST /api/v1/users", userHandler.CreateUser)
	return mux
}
