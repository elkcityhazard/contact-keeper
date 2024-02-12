package errors

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Errors map[string][]string `json:"errors"`
}

// New Errors returns a new error struct with an empty map of errors

func NewErrors() *Error {
	return &Error{
		Errors: make(map[string][]string),
	}
}

// AddError adds an error to the error struct

func (e *Error) AddError(field, message string) {
	e.Errors[field] = append(e.Errors[field], message)
}

// HasErrors returns true if the error struct has any errors

func (e *Error) HasErrors() bool {
	return len(e.Errors) > 0
}

// Get returns the errors for a given field

func (e *Error) Get(field string) []string {
	return e.Errors[field]
}

func (e *Error) SendJSON(errorCode int, w http.ResponseWriter, envelope string) error {

	w.WriteHeader(errorCode)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		return err
	}
	return nil
}
