package service

import (
	"context"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/admin"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
)

type adminService struct {
	adminUC admin.AdminUC
	logger  logger.Logger
}

func NewadminService(adminUC admin.AdminUC, logger logger.Logger) *adminService {
	return &adminService{
		adminUC: adminUC,
		logger:  logger,
	}
}

func (s *adminService) Create(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("create"))
	}
}

func (s *adminService) Search(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("search"))
	}
}
