package user

import (
	"bytes"
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

func WriteJsonResponse(w http.ResponseWriter, data any, statusCode int) error {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(buf.Bytes())
	return err
}
