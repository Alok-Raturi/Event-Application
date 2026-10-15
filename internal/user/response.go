package user

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	apiError := APIError{
		Code:    code,
		Message: message,
	}
	_ = json.NewEncoder(w).Encode(apiError)
}
