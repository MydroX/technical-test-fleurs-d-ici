package error

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewInternal(serviceName string, message string) *Error {
	log.Fatalf("%s | %s", serviceName, message)
	return New(http.StatusInternalServerError, "internal error")
}

func (e *Error) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(e.Code)
	json.NewEncoder(w).Encode(e)
}
