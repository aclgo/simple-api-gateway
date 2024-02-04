package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aclgo/simple-api-gateway/internal/user"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
	mail "github.com/aclgo/simple-api-gateway/proto-service/mail"
	protoUser "github.com/aclgo/simple-api-gateway/proto-service/user"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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
		Verified:  created.User.Verified,
		CreatedAt: created.User.CreatedAt.AsTime(),
		UpdatedAt: created.User.UpdatedAt.AsTime(),
	}, nil
}

func (u *userUc) Login(ctx context.Context, params *user.ParamsUserLoginRequest) (*user.ParamsUserLoginResponse, error) {
	resp, err := u.clientUserGRPC.Login(ctx, &protoUser.UserLoginRequest{
		Email:    params.Email,
		Password: params.Password,
	})

	if err != nil && err != (user.ErrUserNotVerified{}) {
		return nil, err
	}

	if err == (user.ErrUserNotVerified{}) {
		return nil, user.ErrUserNotVerified{}
	}

	return &user.ParamsUserLoginResponse{
		AccessToken:  resp.Tokens.AccessToken,
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
		Verified:  resp.User.Verified,
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
		Verified:  resp.User.Verified,
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
		Verified:  updated.User.Verified,
		CreatedAt: updated.User.CreatedAt.AsTime(),
		UpdatedAt: updated.User.UpdatedAt.AsTime(),
	}, nil
}

func (u *userUc) Delete(ctx context.Context, params *user.ParamsUserDelete) error {
	_, err := u.clientUserGRPC.Delete(ctx, &protoUser.DeleteRequest{Id: params.UserID})
	return err
}

func (u *userUc) RefreshTokens(ctx context.Context, params *user.ParamsRefreshTokens) (*user.RefreshTokens, error) {
	tokens, err := u.clientUserGRPC.RefreshTokens(ctx, &protoUser.RefreshTokensRequest{
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
	})

	if err != nil {
		return nil, err
	}

	return &user.RefreshTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (u *userUc) SendConfirm(ctx context.Context, params *user.ParamsConfirm) error {

	err := u.redisClient.Get(ctx, params.To).Err()
	if err != nil && err != redis.Nil {
		return err
	}

	if err == redis.Nil {

		confirmID := uuid.NewString()

		req := &mail.MailRequest{
			From:        user.DefaultFromSendMail,
			To:          params.To,
			Subject:     user.DefaultSubjectSendConfirm,
			Body:        fmt.Sprintf(user.DefaulfBodySendConfirm, confirmID),
			Template:    user.DefaulfTemplateSendConfirm,
			Servicename: user.DefaultServiceName,
		}

		fmt.Println(req)

		_, err = u.clientMailGRPC.SendService(ctx, req)
		if err != nil {
			return err
		}

		if err := u.redisClient.Set(ctx, params.To, confirmID, user.DefaultTimeSendEmails).Err(); err != nil {
			return err
		}

		if err := u.redisClient.Set(ctx, confirmID, params.To, user.DefaultTimeSendEmails).Err(); err != nil {
			return err
		}

		return nil
	}

	return user.ErrEmailSentCheckInbox{}
}

func (u *userUc) SendConfirmOK(ctx context.Context, params *user.ParamsConfirmOK) error {

	userEmail, err := u.redisClient.Get(ctx, params.ConfirmCode).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if err == redis.Nil {
		return user.ErrInvalidCode{}
	}

	foundUser, err := u.clientUserGRPC.FindByEmail(ctx, &protoUser.FindByEmailRequest{Email: userEmail})

	if err != nil {
		return err
	}

	in := protoUser.UpdateRequest{
		Id:       foundUser.User.Id,
		Verified: "yes",
	}

	_, err = u.clientUserGRPC.Update(ctx, &in)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUc) ResetPass(ctx context.Context, params *user.ParamsResetPass) error {
	switch _, err := u.clientUserGRPC.FindByEmail(ctx, &protoUser.FindByEmailRequest{Email: params.Email}); {
	case errors.Is(err, redis.Nil):
		resetCode := uuid.NewString()

		in := mail.MailRequest{
			From:     user.DefaultFromSendMail,
			To:       params.Email,
			Subject:  user.DefaultSubjectResetPass,
			Body:     fmt.Sprintf(user.DefaultBodyResetPass, resetCode),
			Template: user.DefaultTemplateResetPass,
		}

		_, err = u.clientMailGRPC.SendService(ctx, &in)
		if err != nil {
			return err
		}

		if err := u.redisClient.Set(ctx, resetCode, params.Email, user.DefaultTimeSendEmails).Err(); err != nil {
			return err
		}

		if err := u.redisClient.Set(ctx, params.Email, resetCode, user.DefaultTimeSendEmails).Err(); err != nil {
			return err
		}

	case err == nil:
		return user.ErrEmailSentCheckInbox{}
	default:
		return err
	}

	return nil
}

func (u *userUc) NewPass(ctx context.Context, params *user.ParamsNewPass) error {

	idUser, err := u.redisClient.Get(ctx, params.NewPassCode).Result()
	if err != nil {
		return err
	}

	_, err = u.clientUserGRPC.FindById(ctx, &protoUser.FindByIdRequest{Id: idUser})
	if err != nil {
		return err
	}

	updated, err := u.clientUserGRPC.Update(ctx, &protoUser.UpdateRequest{
		Id:       idUser,
		Password: params.NewPass,
	})

	if err != nil {
		return err
	}

	fmt.Println(updated)

	return nil
}
