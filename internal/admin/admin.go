package admin

import (
	"context"
	"time"
)

type AdminUC interface {
	Create(ctx context.Context, params *ParamsCreateAdmin) (*Admin, error)
	Search(ctx context.Context, params *ParamsSearch) ([]*Admin, error)
}

type ParamsCreateAdmin struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Admin struct {
	UserID    string    `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ParamsSearch struct {
	Query  string
	Role   string
	Page   int
	OffSet int
	Limit  int
}

func (p *ParamsSearch) Validate() error { return nil }

type ErrEmaiCadastred struct {
	Message string
}

func (e ErrEmaiCadastred) Error() string {
	return "email cadastred"
}
