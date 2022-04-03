package error

import "net/http"

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type FieldValidationError struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
}

func NotFoundError(details interface{}) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
		Details: details,
	}
}

func ServerError(details interface{}) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Details: details,
	}
}

func ValidationError(details interface{}) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Details: details,
	}
}
