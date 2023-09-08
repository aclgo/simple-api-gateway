package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aclgo/simple-api-gateway/internal/handlers"
	"github.com/aclgo/simple-api-gateway/internal/user"
	userService "github.com/aclgo/simple-api-gateway/proto-services/user"
	"google.golang.org/grpc"
)

var (
	AddrServiceUser    = ":50051"
	OptionsServiceUser = []grpc.DialOption{}

	AddrServiceAdmin    = ":50052"
	OptionsServiceAdmin = []grpc.DialOption{}
)

func main() {
	fmt.Println("simple-api-gateway")

	conn, err := grpc.Dial(AddrServiceUser, OptionsServiceUser...)
	if err != nil {
		log.Fatalf("grpc.Dial:%v", err)
	}

	clientUserService := userService.NewUserServiceClient(conn)

	user := user.NewUser(clientUserService)
	userHandler := handlers.NewUserHandler(user)

	ctx := context.Background()

	http.HandleFunc("/api/user/login", userHandler.Login(ctx))
	http.HandleFunc("/api/user/logout", userHandler.Logout(ctx))
	http.HandleFunc("/api/user/register", userHandler.Register(ctx))
	http.HandleFunc("/api/user/find", userHandler.Find(ctx))
	http.HandleFunc("/api/user/update", userHandler.Update(ctx))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("http.ListenAndServe:%v", err)
	}
}
