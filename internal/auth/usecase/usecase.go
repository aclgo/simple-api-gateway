package authUC

import (
	"context"
	"fmt"
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
	if len(accessToken) < 7 && accessToken != "baerer " {
		return ""
	}

	return accessToken
}

func getRefreshToken(r *http.Request) string {
	refreshToken := r.Header.Get("refresh-token")
	if len(refreshToken) < 7 && refreshToken != "baerer " {
		return ""
	}

	return refreshToken
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

		fmt.Println(paramsToken)

		next.ServeHTTP(w, r)

	}
}
func (a *authUC) ValidateUpdate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		next.ServeHTTP(w, r)
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

		next.ServeHTTP(w, r)

	}
}
