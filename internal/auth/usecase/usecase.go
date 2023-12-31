package authUC

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/auth"

	protoUser "github.com/aclgo/simple-api-gateway/proto-service/user"
)

type authUC struct {
	userSvcClient protoUser.UserServiceClient
}

func NewAuthUC(userSvcClient protoUser.UserServiceClient) *authUC {
	return &authUC{
		userSvcClient: userSvcClient,
	}
}

func (a *authUC) validateToken(ctx context.Context, token string) (*auth.ParamsToken, error) {
	resp, err := a.userSvcClient.ValidateToken(
		ctx,
		&protoUser.ValidateTokenRequest{Token: token},
	)

	if err != nil {
		return nil, err
	}

	return &auth.ParamsToken{
		UserID: resp.UserID,
		Role:   resp.UserRole,
	}, nil

}

func getAccessToken(r *http.Request) string {
	accessToken := r.Header.Get("access-token")
	if len(accessToken) < 7 && accessToken[:7] != "baerer " {
		return ""
	}

	return accessToken[7:]
}

func getRefreshToken(r *http.Request) string {
	refreshToken := r.Header.Get("refresh-token")
	if len(refreshToken) < 7 && refreshToken[:7] != "baerer " {
		return ""
	}

	return refreshToken[7:]
}

func (a *authUC) ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := getAccessToken(r)
		if accessToken == "" {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: auth.ErrInvalidToken{}.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		paramsToken, err := a.validateToken(context.Background(), accessToken)
		if err != nil {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		v := context.WithValue(context.Background(), auth.KeyCtxParamsToken, paramsToken)

		next.ServeHTTP(w, r.WithContext(v))

	}
}
func (a *authUC) ValidateUpdate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := getAccessToken(r)
		if accessToken == "" {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: auth.ErrInvalidToken{}.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		_, err := a.validateToken(context.Background(), accessToken)
		if err != nil {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		params := auth.ParamsUpdate{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusBadRequest),
				Message: auth.ErrInvalidToken{}.Error(),
			}

			auth.Json(w, resp, http.StatusBadRequest)

			return
		}

		v := context.WithValue(context.Background(), auth.KeyCtxParamsUpdate, params)

		next.ServeHTTP(w, r.WithContext(v))
	}
}
func (a *authUC) ValidateCreateAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := getAccessToken(r)
		if accessToken == "" {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: auth.ErrInvalidToken{}.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		paramsToken, err := a.validateToken(context.Background(), accessToken)
		if err != nil {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		if paramsToken.Role != string(auth.SUPERADMIN) {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		params := auth.ParamsCreateAdmin{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusBadRequest),
				Message: auth.ErrInvalidToken{}.Error(),
			}

			auth.Json(w, resp, http.StatusBadRequest)

			return
		}

		v := context.WithValue(context.Background(), auth.KeyCtxParamsCreateAdmin, params)

		next.ServeHTTP(w, r.WithContext(v))

	}
}
func (a *authUC) ValidateIsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := getAccessToken(r)
		if accessToken == "" {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: auth.ErrInvalidToken{}.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		paramsToken, err := a.validateToken(context.Background(), accessToken)
		if err != nil {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		if paramsToken.Role != string(auth.ADMIN) {
			resp := auth.Response{
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			}

			auth.Json(w, resp, http.StatusUnauthorized)

			return
		}

		v := context.WithValue(context.Background(), auth.KeyCtxParamsToken, paramsToken)

		next.ServeHTTP(w, r.WithContext(v))

	}
}
