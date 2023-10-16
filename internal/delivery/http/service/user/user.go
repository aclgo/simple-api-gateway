package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/delivery/http/service"
	"github.com/aclgo/simple-api-gateway/internal/user"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
)

type userService struct {
	userService user.UserUC
	logger      logger.Logger
}

func NewuserService(userSvc user.UserUC, logger logger.Logger) *userService {
	return &userService{
		userService: userSvc,
		logger:      logger,
	}
}

func (s *userService) Register(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params user.ParamsUserRegister

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			service.JSON(w, service.NewRestError(http.StatusBadRequest, err.Error(), "json.NewDecoder"), http.StatusBadRequest)
			return
		}

		if err := params.Validate(); err != nil {
			service.JSON(w, service.NewRestError(http.StatusBadRequest, err.Error(), "params.Validate"), http.StatusBadRequest)
			return
		}

		created, err := s.userService.Register(ctx, &params)
		if err != nil {
			service.JSON(w, service.ParseError(err, "s.userService.Register"), http.StatusBadRequest)
			return
		}

		service.JSON(w, created, http.StatusCreated)
	}
}

func (s *userService) Login(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserLoginRequest{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			service.JSON(w, service.NewRestError(http.StatusBadRequest, err.Error(), "json.NewDecoder"), http.StatusBadRequest)
			return
		}

		if err := params.Validate(); err != nil {
			service.JSON(w, service.NewRestError(http.StatusBadRequest, err.Error(), "params.Validate"), http.StatusBadRequest)
			return
		}

		logged, err := s.userService.Login(ctx, &params)
		if err != nil {
			service.JSON(w, service.ParseError(err, "s.userService.Login"), http.StatusInternalServerError)
			return
		}

		service.JSON(w, logged, http.StatusOK)
	}
}

func (s *userService) Logout(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserLogout{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		if err := params.Validate(); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		if err := s.userService.Logout(ctx, &params); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		service.JSON(w, nil, http.StatusOK)
	}
}

func (s *userService) Find(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserFindById{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		if err := params.Validate(); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		resp, err := s.userService.FindById(ctx, &params)
		if err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) Update(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserUpdate{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		if err := params.Validate(); err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		resp, err := s.userService.Update(ctx, &params)
		if err != nil {
			service.JSON(w, nil, http.StatusBadRequest)
			return
		}

		service.JSON(w, resp, http.StatusOK)

	}
}

func (s *userService) UserConfirm(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("confirm"))
	}
}

func (s *userService) UserResetPass(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("reset pass"))
	}
}

func (s *userService) UserNewPass(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("new pass"))
	}
}
