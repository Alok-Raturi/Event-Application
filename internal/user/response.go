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

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, code string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	apiError := APIError{
		Code:    code,
		Message: message,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(apiError)

	if err != nil {
		return ErrJsonEncodingError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return ErrResponseClientError
	}
	return nil
}

func WriteJsonResponse(w http.ResponseWriter, data any, statusCode int) error {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return ErrJsonEncodingError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return ErrResponseClientError
	}
	return nil
}
