package utils

import "fmt"

type error interface {
	Error() string
}

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Err)
}
