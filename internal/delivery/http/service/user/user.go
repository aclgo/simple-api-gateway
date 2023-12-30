package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/auth"
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
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		created, err := s.userService.Register(ctx, &params)
		if err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]any{
			"message": "user created",
			"user":    created,
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) Login(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserLoginRequest{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		logged, err := s.userService.Login(ctx, &params)
		if err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]any{
			"message": "user logged",
			"tokens":  logged,
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) Logout(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserLogout{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := s.userService.Logout(ctx, &params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]any{
			"message": "user logout",
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) Find(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsUserFindById{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		found, err := s.userService.FindById(ctx, &params)
		if err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]any{
			"message": "user found",
			"user":    found,
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) Update(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxData, ok := r.Context().Value(auth.KeyCtxParamsUpdate).(auth.ParamsUpdate)
		if !ok {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), service.ErrNoParamsInCtx.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		params := user.ParamsUserUpdate{
			UserID:   ctxData.UserID,
			Name:     ctxData.Name,
			Lastname: ctxData.Lastname,
			Password: ctxData.Password,
			Email:    ctxData.Email,
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		updated, err := s.userService.Update(ctx, &params)
		if err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]any{
			"message": "user updated",
			"user":    updated,
		}

		service.JSON(w, resp, http.StatusOK)

	}
}

func (s *userService) UserConfirm(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsConfirmOK{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := s.userService.SendConfirmOK(ctx, &params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]string{
			"message": "user confirmed signup",
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) UserResetPass(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsResetPass{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)
			return

		}

		if err := s.userService.ResetPass(ctx, &params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]string{
			"message": "code to reset pass send to email",
		}

		service.JSON(w, resp, http.StatusOK)
	}
}

func (s *userService) UserNewPass(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := user.ParamsNewPass{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := params.Validate(); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusBadRequest), err.Error())

			service.JSON(w, response, http.StatusBadRequest)

			return
		}

		if err := s.userService.NewPass(ctx, &params); err != nil {
			response := service.NewRestError(http.StatusText(http.StatusInternalServerError), err.Error())

			service.JSON(w, response, http.StatusInternalServerError)

			return
		}

		resp := map[string]string{
			"message": "user updated pass",
		}

		service.JSON(w, resp, http.StatusOK)

	}
}
