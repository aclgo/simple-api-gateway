package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/aclgo/simple-api-gateway/internal/user"
)

type RestError struct {
	ErrStatus  int       `json:"status,omitempty"`
	ErrError   string    `json:"error,omitempty"`
	ErrMessage any       `json:"message,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func NewRestError(status int, err string, message string) *RestError {
	return &RestError{
		ErrStatus:  status,
		ErrError:   err,
		ErrMessage: message,
		Timestamp:  time.Now(),
	}
}

func ParseError(err error, msg string) *RestError {
	switch {
	case errors.Is(err, user.ErrEmailCadastred{}):
		NewRestError(http.StatusBadRequest, err.Error(), msg)
	}

	return NewRestError(http.StatusInternalServerError, err.Error(), msg)
}
