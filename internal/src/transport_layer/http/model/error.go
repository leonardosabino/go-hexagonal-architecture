package model

import (
	"net/http"
	"time"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

const layout = "2006-01-02T15:04:05.999999Z07:00"

var (
	namespace = errorx.NewNamespace("common")
	NotFound  = errorx.NewType(namespace, "not_found", errorx.NotFound())
)

func ErrorHandlerResponse(errx *errorx.Error) *Error {
	var errorResponse *Error

	switch true {
	case errorx.HasTrait(errx, errorx.NotFound()):
		errorResponse = NewError(http.StatusNotFound, errx.Error())
	case errorx.IsOfType(errx, errorx.IllegalArgument):
		errorResponse = NewError(http.StatusBadRequest, errx.Error())
	default:
		errorResponse = NewError(http.StatusInternalServerError, errx.Error())
	}

	return errorResponse
}

type Error struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
}

func NewError(statusCode int, message string) *Error {
	return &Error{
		Status:     http.StatusText(statusCode),
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  time.Now().Format(layout),
	}
}

type logFunc func(*errorx.Error)

func ResponseError(context echo.Context, errx error) error {
	// logFunc(errx)
	// errorResponse := ErrorHandlerResponse(errx)
	return errx
}
