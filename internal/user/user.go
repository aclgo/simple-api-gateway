package user

import (
	"context"
	"time"

	userService "github.com/aclgo/simple-api-gateway/proto-services/user"
)

type UserService struct {
	userServiceClient userService.UserServiceClient
}

func NewUser(userServiceClient userService.UserServiceClient) *UserService {
	return &UserService{
		userServiceClient: userServiceClient,
	}
}

type ParamsRegister struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ParamsRegistredUser struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserService) Register(ctx context.Context, params *ParamsRegister) (*ParamsRegistredUser, error) {
	created, err := u.userServiceClient.Register(
		ctx,
		&userService.CreateUserRequest{
			Name:     params.Name,
			LastName: params.Lastname,
			Password: params.Password,
			Email:    params.Email,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ParamsRegistredUser{
		UserID:    created.User.Id,
		Name:      created.User.Name,
		Lastname:  created.User.LastName,
		Password:  created.User.Password,
		Email:     created.User.Email,
		Role:      created.User.Role,
		CreatedAt: created.User.CreatedAt.AsTime(),
		UpdatedAt: created.User.UpdatedAt.AsTime(),
	}, nil
}

type ParamsLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ParamsTokensLogin struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (u *UserService) Login(ctx context.Context, params *ParamsLogin) (*ParamsTokensLogin, error) {
	tokens, err := u.userServiceClient.Login(
		ctx,
		&userService.UserLoginRequest{
			Email:    params.Email,
			Password: params.Password,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ParamsTokensLogin{
		Access:  tokens.Tokens.AccessToken,
		Refresh: tokens.Tokens.RefreshToken,
	}, nil
}

func (u *UserService) Logout(ctx context.Context, params *ParamsTokensLogin) error {
	_, err := u.userServiceClient.Logout(
		ctx,
		&userService.UserLogoutRequest{
			AccessToken:  params.Access,
			RefreshToken: params.Refresh,
		},
	)

	if err != nil {
		return err
	}

	return nil

}

type ParamsFindById struct {
	UserID string `json:"user_id"`
}

func (u *UserService) FindById(ctx context.Context, params *ParamsFindById) (*ParamsRegistredUser, error) {
	found, err := u.userServiceClient.FindById(ctx, &userService.FindByIdRequest{Id: params.UserID})
	if err != nil {
		return nil, err
	}

	return &ParamsRegistredUser{
		UserID:    found.User.Id,
		Name:      found.User.Name,
		Lastname:  found.User.LastName,
		Password:  found.User.Password,
		Email:     found.User.Email,
		Role:      found.User.Role,
		CreatedAt: found.User.CreatedAt.AsTime(),
		UpdatedAt: found.User.UpdatedAt.AsTime(),
	}, nil
}

type ParamsFindByEmail struct {
	UserEmail string `json:"user_email"`
}

func (u *UserService) FindByEmail(ctx context.Context, params *ParamsFindByEmail) (*ParamsRegistredUser, error) {
	found, err := u.userServiceClient.FindByEmail(ctx, &userService.FindByEmailRequest{Email: params.UserEmail})
	if err != nil {
		return nil, err
	}

	return &ParamsRegistredUser{
		UserID:    found.User.Id,
		Name:      found.User.Name,
		Lastname:  found.User.LastName,
		Password:  found.User.Password,
		Email:     found.User.Email,
		Role:      found.User.Role,
		CreatedAt: found.User.CreatedAt.AsTime(),
		UpdatedAt: found.User.UpdatedAt.AsTime(),
	}, nil
}

type ParamsUpdate struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *UserService) Update(ctx context.Context, params *ParamsUpdate) (*ParamsRegistredUser, error) {
	updated, err := u.userServiceClient.Update(
		ctx,
		&userService.UpdateRequest{
			Id:       params.UserID,
			Name:     params.Name,
			Lastname: params.Lastname,
			Password: params.Password,
			Email:    params.Email,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ParamsRegistredUser{
		UserID:    updated.User.Id,
		Name:      updated.User.Name,
		Lastname:  updated.User.LastName,
		Password:  updated.User.Password,
		Email:     updated.User.Email,
		Role:      updated.User.Role,
		CreatedAt: updated.User.CreatedAt.AsTime(),
		UpdatedAt: updated.User.UpdatedAt.AsTime(),
	}, nil
}
