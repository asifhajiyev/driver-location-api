package error

import "net/http"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) AsResponse() *Error {
	return &Error{
		Message: e.Message,
	}
}

func (e Error) AsMessage() string {
	return e.Message
}

func NotFoundError(msg string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func ServerError(msg string) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func UnAuthorizedError(msg string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func ValidationError(msg string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}
