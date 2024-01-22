package user

import (
	"context"
	"fmt"
	"time"
)

type UserUC interface {
	Register(ctx context.Context, params *ParamsUserRegister) (*User, error)
	Login(ctx context.Context, params *ParamsUserLoginRequest) (*ParamsUserLoginResponse, error)
	Logout(ctx context.Context, params *ParamsUserLogout) error
	FindById(ctx context.Context, params *ParamsUserFindById) (*User, error)
	FindByEmail(ctx context.Context, params *ParamsUserFindByEmail) (*User, error)
	Update(ctx context.Context, params *ParamsUserUpdate) (*User, error)
	SendConfirm(ctx context.Context, params *ParamsConfirm) error
	SendConfirmOK(ctx context.Context, params *ParamsConfirmOK) error
	ResetPass(ctx context.Context, params *ParamsResetPass) error
	NewPass(ctx context.Context, params *ParamsNewPass) error
}

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ParamsUserRegister struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (p *ParamsUserRegister) Validate() error { return nil }

type ParamsUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *ParamsUserLoginRequest) Validate() error { return nil }

type ParamsUserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ParamsUserLogout struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (p *ParamsUserLogout) Validate() error {
	if p.AccessToken == "" {
		return fmt.Errorf("access_token empty")
	}

	if p.RefreshToken == "" {
		return fmt.Errorf("refresh_token empty")
	}

	return nil
}

type ParamsUserFindById struct {
	UserID string `json:"user_id"`
}

func (p *ParamsUserFindById) Validate() error {
	if p.UserID == "" {
		return fmt.Errorf("user_id empty")
	}
	return nil
}

type ParamsUserFindByEmail struct {
	Email string `json:"email"`
}

func (p *ParamsUserFindByEmail) Validate() error {
	if p.Email == "" {
		return fmt.Errorf("email empty")
	}
	return nil
}

type ParamsUserUpdate struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (p *ParamsUserUpdate) Validate() error {
	if p.UserID == "" {
		return fmt.Errorf("user_id empty")
	}

	return nil
}

type ParamsConfirm struct {
	To           string
	IntervalSend time.Duration
}

type ParamsConfirmOK struct {
	ConfirmCode string
}

func (p *ParamsConfirmOK) Validate() error {
	if p.ConfirmCode == "" {
		return ErrInvalidCode{}
	}

	return nil
}

type ParamsResetPass struct {
	Email string `json:"email"`
}

func (p *ParamsResetPass) Validate() error {
	if p.Email == "" {
		return ErrInvalidEmail{}
	}

	return nil
}

type ParamsNewPass struct {
	NewPassCode string `json:"code"`
	NewPass     string `json:"new_pass"`
	ConfirmPass string `json:"confirm_pass"`
}

func (p *ParamsNewPass) Validate() error {
	if p.NewPass == "" {
		return fmt.Errorf("new_pass empty")
	}

	if p.ConfirmPass == "" {
		return fmt.Errorf("confirm_pass empty")
	}

	if p.ConfirmPass != p.NewPass {
		return fmt.Errorf("passwords not match")
	}

	return nil
}

type ParamsRefreshTokens struct {
	AccessToken  string
	RefreshToken string
}

func (p *ParamsRefreshTokens) Validate() error {
	if p.AccessToken == "" {
		return ErrEmptyToken{}
	}

	if p.RefreshToken == "" {
		return ErrEmptyToken{}
	}

	return nil
}

type RefreshTokens struct {
	AccessToken  string
	RefreshToken string
}
