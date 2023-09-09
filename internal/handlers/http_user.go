package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/user"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUserHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) Register(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params user.ParamsRegister

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		created, err := u.userService.Register(ctx, &params)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(created)
		if err != nil {
			log.Println(err)
		}
	}
}

func (u *UserHandler) Login(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params user.ParamsLogin

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		tokensLogin, err := u.userService.Login(ctx, &params)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(tokensLogin)
		if err != nil {
			log.Println(err)
		}
	}
}

func (u *UserHandler) Logout(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params user.ParamsTokensLogin

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if err := u.userService.Logout(ctx, &params); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

var (
	Empty = ""
	Id    = ""
	Email = ""
)

func (u *UserHandler) Find(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get(Id)
		email := r.URL.Query().Get(Email)

		if id != Empty {
			return
		}

		if email != Empty {
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (u *UserHandler) Update(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params user.ParamsUpdate

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		updated, err := u.userService.Update(ctx, &params)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(updated)
		if err != nil {
			log.Println(err)
		}
	}
}
