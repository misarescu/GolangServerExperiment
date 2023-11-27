package models

import "fmt"

type RequestError struct{
	Message	string `json:"message"`
}

func (e *RequestError) Error() string{
	return fmt.Sprintf("Request error: %s", e.Message)
}

func NewRequestError(msg string) *RequestError{
	return &RequestError{
		Message: msg,
	}
}