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
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "INTERNAL_SERVER_ERROR")
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
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid Inputs provided", "INVALID_INPUT")
		return
	}

	user, err := h.userSvc.CreateUser(&userReq)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidInput):
			{
				WriteErrorResponse(w, http.StatusBadRequest, "Invalid Inputs provided", "INVALID_INPUT")
				return
			}
		case errors.Is(err, ErrEmailAlreadyExists):
			{
				WriteErrorResponse(w, http.StatusConflict, "Email already exist", "EMAIL_ALREADY_EXIST")
				return
			}
		default:
			WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "INTERNAL_SERVER_ERROR")
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
	if err != nil {
		return
	}
}
