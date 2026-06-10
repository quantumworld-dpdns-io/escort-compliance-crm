package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrNotFound           ErrorCode = "NOT_FOUND"
	ErrUnauthorized       ErrorCode = "UNAUTHORIZED"
	ErrForbidden          ErrorCode = "FORBIDDEN"
	ErrBadRequest         ErrorCode = "BAD_REQUEST"
	ErrConflict           ErrorCode = "CONFLICT"
	ErrInternal           ErrorCode = "INTERNAL_ERROR"
	ErrValidation         ErrorCode = "VALIDATION_ERROR"
	ErrRateLimited        ErrorCode = "RATE_LIMITED"
	ErrServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	ErrQuantumUnavailable ErrorCode = "QUANTUM_UNAVAILABLE"
)

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
	HTTP    int       `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		HTTP:    codeToHTTP(code),
	}
}

func Newf(code ErrorCode, format string, args ...interface{}) *AppError {
	return &AppError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		HTTP:    codeToHTTP(code),
	}
}

func Wrap(err error, code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: err.Error(),
		HTTP:    codeToHTTP(code),
	}
}

func codeToHTTP(code ErrorCode) int {
	switch code {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrForbidden:
		return http.StatusForbidden
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrConflict:
		return http.StatusConflict
	case ErrValidation:
		return http.StatusUnprocessableEntity
	case ErrRateLimited:
		return http.StatusTooManyRequests
	case ErrServiceUnavailable:
		return http.StatusServiceUnavailable
	case ErrQuantumUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
