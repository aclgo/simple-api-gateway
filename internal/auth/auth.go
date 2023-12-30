package auth

import "net/http"

type Auth interface {
	ValidateToken(next http.HandlerFunc) http.HandlerFunc
	ValidateUpdate(next http.HandlerFunc) http.HandlerFunc
	ValidateCreateAdmin(next http.HandlerFunc) http.HandlerFunc
	ValidateIsAdmin(next http.HandlerFunc) http.HandlerFunc
}

type (
	Level                string
	CtxParamsUpdate      ParamsUpdate
	CtxParamsCreateAdmin ParamsCreateAdmin
	CtxParamsToken       ParamsToken
)

var (
	SUPERADMIN Level = "super-admin"
	ADMIN      Level = "admin"

	KeyCtxParamsUpdate      CtxParamsUpdate
	KeyCtxParamsCreateAdmin CtxParamsCreateAdmin
	KeyCtxParamsToken       CtxParamsToken
)

type ParamsUpdate struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ParamsCreateAdmin struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ParamsToken struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
