package service

import (
	"errors"

	"github.com/aclgo/simple-api-gateway/internal/user"
)

var (
	ErrNoParamsInCtx = errors.New("no params in ctx")
)

type RestError struct {
	ErrError   string `json:"error,omitempty"`
	ErrMessage any    `json:"message,omitempty"`
}

func NewRestError(err string, message string) *RestError {
	return &RestError{
		ErrError:   err,
		ErrMessage: message,
	}
}

func ParseError(err error, msg string) *RestError {
	switch {
	case errors.Is(err, user.ErrEmailCadastred{}):
		NewRestError(err.Error(), msg)
	}

	return NewRestError(err.Error(), msg)
}
