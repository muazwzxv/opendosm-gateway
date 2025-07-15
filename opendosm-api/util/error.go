package util

import (
	"errors"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/slog"
)

type ErrorResponse struct {
	HttpCode int    `json:"http_code"`
	Message  string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return "" // TODO: figure it out
}

func BuildError(httpCode int, message string) error {
	return &ErrorResponse{
		HttpCode: httpCode,
		Message:  message,
	}
}

func HandleError(resp *goyave.Response, log *slog.Logger, err error) {
	var errData *ErrorResponse
	if !errors.As(err, &errData) {
		log.Warn("invalid error data")
	}
	resp.JSON(errData.HttpCode, errData)
}
