package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/aclgo/simple-api-gateway/internal/user"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
	"github.com/aclgo/simple-api-gateway/proto-service/mail"
	protoUser "github.com/aclgo/simple-api-gateway/proto-service/user"
	"github.com/go-redis/redis/v8"
)

type userUc struct {
	clientUserGRPC protoUser.UserServiceClient
	clientMailGRPC mail.MailServiceClient
	redisClient    *redis.Client
	logger         logger.Logger
}

func NewuserUC(clientUser protoUser.UserServiceClient,
	clientMail mail.MailServiceClient,
	redisClient *redis.Client, logger logger.Logger) *userUc {
	return &userUc{
		clientUserGRPC: clientUser,
		clientMailGRPC: clientMail,
		redisClient:    redisClient,
		logger:         logger,
	}
}

func (u *userUc) Register(ctx context.Context, params *user.ParamsUserRegister) (*user.User, error) {
	created, err := u.clientUserGRPC.Register(ctx, &protoUser.CreateUserRequest{
		Name:     params.Name,
		LastName: params.Lastname,
		Password: params.Password,
		Email:    params.Email,
	})

	if err != nil {
		return nil, err
	}

	return &user.User{
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

func (u *userUc) Login(ctx context.Context, params *user.ParamsUserLoginRequest) (*user.ParamsUserLoginResponse, error) {
	resp, err := u.clientUserGRPC.Login(ctx, &protoUser.UserLoginRequest{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user.ParamsUserLoginResponse{
		AccessToken:  resp.Tokens.RefreshToken,
		RefreshToken: resp.Tokens.RefreshToken,
	}, nil
}
func (u *userUc) Logout(ctx context.Context, params *user.ParamsUserLogout) error {
	_, err := u.clientUserGRPC.Logout(ctx, &protoUser.UserLogoutRequest{
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
	})

	if err != nil {
		return err
	}

	return nil
}
func (u *userUc) FindById(ctx context.Context, params *user.ParamsUserFindById) (*user.User, error) {
	resp, err := u.clientUserGRPC.FindById(ctx, &protoUser.FindByIdRequest{Id: params.UserID})
	if err != nil {
		return nil, err
	}

	return &user.User{
		UserID:    resp.User.Id,
		Name:      resp.User.Name,
		Lastname:  resp.User.LastName,
		Password:  resp.User.Password,
		Email:     resp.User.Email,
		Role:      resp.User.Role,
		CreatedAt: resp.User.CreatedAt.AsTime(),
		UpdatedAt: resp.User.UpdatedAt.AsTime(),
	}, nil
}
func (u *userUc) FindByEmail(ctx context.Context, params *user.ParamsUserFindByEmail) (*user.User, error) {
	resp, err := u.clientUserGRPC.FindByEmail(ctx, &protoUser.FindByEmailRequest{Email: params.Email})
	if err != nil {
		return nil, err
	}

	return &user.User{
		UserID:    resp.User.Id,
		Name:      resp.User.Name,
		Lastname:  resp.User.LastName,
		Password:  resp.User.Password,
		Email:     resp.User.Email,
		Role:      resp.User.Role,
		CreatedAt: resp.User.CreatedAt.AsTime(),
		UpdatedAt: resp.User.UpdatedAt.AsTime(),
	}, nil
}
func (u *userUc) Update(ctx context.Context, params *user.ParamsUserUpdate) (*user.User, error) {
	updated, err := u.clientUserGRPC.Update(ctx,
		&protoUser.UpdateRequest{
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

	return &user.User{
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

const (
	redisUser        = "user-id: %s"
	redisCodeNewPass = "new-pass: %s"
)

func (u *userUc) SendConfirm(ctx context.Context, params *user.ParamsConfirm) error {

	err := u.redisClient.Get(ctx, params.To).Err()
	if err != nil && err != redis.Nil {
		return err
	}

	_, err = u.clientMailGRPC.SendService(ctx, &mail.MailRequest{
		From:     "",
		To:       params.To,
		Subject:  "",
		Body:     "",
		Template: "",
	})
	if err != nil {
		return err
	}

	if err := u.redisClient.Set(ctx, params.To, nil, time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
func (u *userUc) ResetPass(ctx context.Context, params *user.ParamsResetPass) error {

	resp, err := u.clientMailGRPC.SendService(ctx, &mail.MailRequest{
		From:     "",
		To:       params.Email,
		Subject:  "",
		Body:     "",
		Template: "",
	})
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
func (u *userUc) NewPass(ctx context.Context, params *user.ParamsNewPass) error {

	// user, err := u.clientUserGRPC.FindById(ctx, &protoUser.FindByIdRequest{Id: params.UserID})

	updated, err := u.clientUserGRPC.Update(ctx, &protoUser.UpdateRequest{
		// Id:       params.UserID,
		Password: params.NewPass,
	})

	if err != nil {
		return err
	}

	fmt.Println(updated)

	return nil
}
