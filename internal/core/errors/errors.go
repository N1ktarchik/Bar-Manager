package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorApp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorApp) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *ErrorApp) StatusCode() *int {
	return &e.Code
}

func (e *ErrorApp) Msg() *string {
	return &e.Message
}

func IsErrorApp(err error) (*ErrorApp, bool) {
	var appErr *ErrorApp
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

func INVALID_JWT_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusBadRequest,
		Message: "Invalid Json-Web Token",
	}
}

func INTERNAL_SERVER_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusInternalServerError,
		Message: "Server error",
	}
}

func BAD_REQUEST_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
}

func SHORT_NAME_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusBadRequest,
		Message: "The name of the cocktail is too short",
	}
}

func INVALID_PRICE_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusBadRequest,
		Message: "The price of a cocktail cannot be less than or equal to zero",
	}
}

func INVALID_ID_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusBadRequest,
		Message: "The ID of a cocktail cannot be less than or equal to zero",
	}
}

func ID_NOT_FAUND_ERR() *ErrorApp {
	return &ErrorApp{
		Code:    http.StatusBadRequest,
		Message: "ID not faund",
	}
}
