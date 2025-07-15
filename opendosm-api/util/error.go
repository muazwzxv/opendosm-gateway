package util

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
