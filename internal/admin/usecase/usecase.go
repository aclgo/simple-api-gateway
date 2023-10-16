package admin

import (
	"context"

	"github.com/aclgo/simple-api-gateway/internal/admin"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
	protoAdmin "github.com/aclgo/simple-api-gateway/proto-service/admin"
	"github.com/go-redis/redis/v8"
)

type adminUC struct {
	clientAdmin protoAdmin.AdminServiceClient
	redisClient *redis.Client
	logger      logger.Logger
}

func NewadminUC(clientAdmin protoAdmin.AdminServiceClient, redisClient *redis.Client, logger logger.Logger) *adminUC {
	return &adminUC{
		clientAdmin: clientAdmin,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (u *adminUC) Create(ctx context.Context, params *admin.ParamsCreateAdmin) (*admin.Admin, error) {
	return nil, nil
}
func (u *adminUC) Search(ctx context.Context, params *admin.ParamsSearch) ([]*admin.Admin, error) {
	return nil, nil
}
