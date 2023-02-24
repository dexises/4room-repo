package apperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ErrNotFound = NewAppError("not found", "", http.StatusNotFound)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             int    `json:"code,omitempty"`
}

func NewAppError(message, developerMessage string, code int) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bytes
}

func UnauthorizedError(message string) *AppError {
	return NewAppError(message, "", http.StatusUnauthorized)
}

func BadRequestError(message string) *AppError {
	return NewAppError(message, "something wrong with user data", http.StatusBadRequest)
}

func MethodNotAllowedResponse(message string) *AppError {
	return NewAppError(message, "method not allowed", http.StatusMethodNotAllowed)
}

func systemError(developerMessage string) *AppError {
	return NewAppError("systen error", developerMessage, http.StatusInternalServerError)
}

func APIError(message, developerMessage string, code int) *AppError {
	return NewAppError(message, developerMessage, code)
}
