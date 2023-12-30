package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/aclgo/simple-api-gateway/internal/admin"
	"github.com/aclgo/simple-api-gateway/internal/user"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
	protoAdmin "github.com/aclgo/simple-api-gateway/proto-service/admin"

	protoMail "github.com/aclgo/simple-api-gateway/proto-service/mail"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type adminUC struct {
	clientAdmin protoAdmin.AdminServiceClient
	clientMail  protoMail.MailServiceClient
	redisClient *redis.Client
	logger      logger.Logger
}

func NewadminUC(clientAdmin protoAdmin.AdminServiceClient, clientMail protoMail.MailServiceClient, redisClient *redis.Client, logger logger.Logger) *adminUC {
	return &adminUC{
		clientAdmin: clientAdmin,
		clientMail:  clientMail,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (u *adminUC) Create(ctx context.Context, params *admin.ParamsCreateAdmin) (*admin.Admin, error) {
	in := protoAdmin.ParamsCreateAdmin{
		Name:     params.Name,
		Lastname: params.Lastname,
		Password: params.Password,
		Email:    params.Email,
		Role:     params.Role,
	}

	created, err := u.clientAdmin.Register(ctx, &in)
	if err != nil {
		return nil, err
	}

	switch err := u.redisClient.Get(ctx, created.Email).Err(); {
	case err == nil:
		return nil, user.ErrEmailSentCheckInbox{}
	case err == redis.Nil:
		confirmID := uuid.NewString()

		m := protoMail.MailRequest{
			From:     admin.DefaultFromSendMail,
			To:       params.Email,
			Subject:  admin.DefaulfBodySendConfirm,
			Body:     fmt.Sprintf(admin.DefaulfBodySendConfirm, confirmID),
			Template: admin.DefaulfTemplateSendConfirm,
		}

		_, err := u.clientMail.SendService(ctx, &m)
		if err != nil {
			return nil, err
		}

		if err := u.redisClient.Set(ctx, params.Email, confirmID, time.Hour).Err(); err != nil {
			return nil, err
		}

		if err := u.redisClient.Set(ctx, confirmID, params.Email, time.Hour).Err(); err != nil {
			return nil, err
		}

	default:
		return nil, err
	}

	return &admin.Admin{
		UserID:    created.UserId,
		Name:      created.Name,
		Lastname:  created.Lastname,
		Password:  created.Password,
		Email:     created.Email,
		Role:      created.Role,
		CreatedAt: created.CreatedAt.AsTime(),
		UpdatedAt: created.UpdatedAt.AsTime(),
	}, nil
}

func (u *adminUC) Search(ctx context.Context, params *admin.ParamsSearch) ([]*admin.Admin, error) {

	in := protoAdmin.ParamsSearchRequest{
		Query:  params.Query,
		Role:   params.Role,
		Page:   int32(params.Page),
		Offset: int32(params.OffSet),
		Limit:  int32(params.Limit),
	}

	users, err := u.clientAdmin.Search(ctx, &in)
	if err != nil {
		return nil, err
	}

	items := make([]*admin.Admin, len(users.Users))

	for i := 0; i < int(users.Total); i++ {
		items[i] = &admin.Admin{
			UserID:    users.Users[i].UserId,
			Name:      users.Users[i].Name,
			Lastname:  users.Users[i].Lastname,
			Password:  users.Users[i].Password,
			Email:     users.Users[i].Email,
			Role:      users.Users[i].Role,
			CreatedAt: users.Users[i].CreatedAt.AsTime(),
			UpdatedAt: users.Users[i].UpdatedAt.AsTime(),
		}
	}

	return items, nil
}
