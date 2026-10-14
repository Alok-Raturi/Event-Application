package user

import (
	"encoding/json"
	"errors"
	"net/http"
)

type UserHandler struct {
	userSvc *UserService
}

func NewHandler(userSvc *UserService) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userSvc.GetUsers()
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = UserResponse{
			ID:        user.ID.String(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}

	err = json.NewEncoder(w).Encode(userResponses)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		http.Error(w, ErrInvalidInput.Error(), http.StatusBadRequest)
		return
	}

	err = h.userSvc.CreateUser(&userReq)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidInput):
			{
				http.Error(w, ErrInvalidInput.Error(), http.StatusBadRequest)
				return
			}
		case errors.Is(err, ErrEmailAlreadyExists):
			{
				http.Error(w, ErrEmailAlreadyExists.Error(), http.StatusConflict)
				return
			}
		default:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}
