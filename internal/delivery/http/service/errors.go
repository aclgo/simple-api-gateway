package service

import (
	"errors"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/user"
	"google.golang.org/grpc/codes"
)

var (
	ErrNoParamsInCtx                 = errors.New("no params in ctx")
	ErrSendEmailAndCancelNewRegister = errors.New("error send email and delete new user registred")
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

func ParseGRPCError(code codes.Code) int {
	switch code {
	case codes.NotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
