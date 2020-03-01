package render

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error represents a custom rendering error
type Error struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e Error) Error() string {
	return e.Message
}

// BadRequestF returns http status code 400 with formatted message
func BadRequestF(msg string, args ...interface{}) Error {
	return Error{
		Status:  400,
		Message: fmt.Sprintf(msg, args...),
	}
}

// InternalF returns http status code 500 with formatted message
func InternalF(msg string, args ...interface{}) Error {
	return Error{
		Status:  500,
		Message: fmt.Sprintf(msg, args...),
	}
}

// Respond with an error
func Respond(w http.ResponseWriter, err Error) error {
	w.WriteHeader(err.Status)
	return json.NewEncoder(w).Encode(&err)
}
