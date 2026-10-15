package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type UserHandler struct {
	userSvc *UserService
	logger  *log.Logger
}

func NewHandler(userSvc *UserService, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
		logger:  logger,
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

	err = WriteJsonResponse(w, userResponses, http.StatusOK)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "INTERNAL_SERVER_ERROR")
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
				WriteErrorResponse(w, http.StatusConflict, "Email already exist", "EMAIL_ALREADY_EXISTS")
				return
			}
		default:
			WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "INTERNAL_SERVER_ERROR")
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	createdUser := &UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	err = WriteJsonResponse(w, createdUser, http.StatusCreated)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "INTERNAL_SERVER_ERROR")
	}
}
