package types

import (
	"errors"
	"net/http"
)

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

const (
	// 5XX
	InternalServiceError ErrorCode = "INTERNAL_SERVICE_ERROR"
	ValidationError      ErrorCode = "VALIDATION_ERROR"
	NotFound             ErrorCode = "NOT_FOUND"
	BadRequest           ErrorCode = "BAD_REQUEST"
	Forbidden            ErrorCode = "FORBIDDEN"
	UnprocessableEntity  ErrorCode = "UNPROCESSABLE_ENTITY"
	RequestTimeout       ErrorCode = "REQUEST_TIMEOUT"
	ClientRequestError   ErrorCode = "CLIENT_REQUEST_ERROR"
)

// ApiError represents an error with an HTTP status code and an application-specific error code.
type Error struct {
	Err        error
	StatusCode int
	ErrorCode  ErrorCode
}

const UninitializedStatusCode = 0

func (e *Error) Error() string {
	return e.Err.Error()
}

// NewError creates a new ApiError with the provided status code, error code, and underlying error.
// If the status code is not provided (0), it defaults to http.StatusInternalServerError(500).
// If the error code is empty, it defaults to INTERNAL_SERVICE_ERROR.
func NewError(statusCode int, errorCode ErrorCode, err error) *Error {
	if statusCode == UninitializedStatusCode {
		statusCode = http.StatusInternalServerError
	}
	if errorCode == "" {
		errorCode = InternalServiceError
	}
	return &Error{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Err:        err,
	}
}

func NewErrorWithMsg(statusCode int, errorCode ErrorCode, msg string) *Error {
	return NewError(statusCode, errorCode, errors.New(msg))
}

func NewInternalServiceError(err error) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  InternalServiceError,
		Err:        err,
	}
}

func NewValidationFailedError(err error) *Error {
	return &Error{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  ValidationError,
		Err:        err,
	}
}

var (
	// ErrInvalidUnbondingTx the transaction spends the unbonding path but is invalid
	ErrInvalidUnbondingTx = errors.New("invalid unbonding tx")

	// ErrInvalidStakingTx the stake transaction is invalid as it does not follow the global parameters
	ErrInvalidStakingTx = errors.New("invalid staking tx")

	// ErrInvalidWithdrawalTx the withdrawal transaction is invalid as it does not unlock the expected time lock path
	ErrInvalidWithdrawalTx = errors.New("invalid withdrawal tx")

	// ErrInvalidSlashingTx the slashing transaction is invalid as it does not unlock the expected slashing path
	ErrInvalidSlashingTx = errors.New("invalid slashing tx")
)
