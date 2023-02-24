package rest_error

import "net/http"

type StandardError struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *StandardError) Error() string {
	return e.Message
}

func NewStandardError(message, err string, code int, causes []Cause) *StandardError {
	return &StandardError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *StandardError {
	return &StandardError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewValidationError(message string, causes []Cause) *StandardError {
	return &StandardError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *StandardError {
	return &StandardError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *StandardError {
	return &StandardError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *StandardError {
	return &StandardError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
