package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aclgo/simple-api-gateway/config"
	adminUC "github.com/aclgo/simple-api-gateway/internal/admin/usecase"
	svcAdmin "github.com/aclgo/simple-api-gateway/internal/delivery/http/service/admin"
	svcUser "github.com/aclgo/simple-api-gateway/internal/delivery/http/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userUC "github.com/aclgo/simple-api-gateway/internal/user/usecase"
	redis "github.com/aclgo/simple-api-gateway/pkg/redis"

	authUC "github.com/aclgo/simple-api-gateway/internal/auth/usecase"
	"github.com/aclgo/simple-api-gateway/pkg/logger"
	protoAdmin "github.com/aclgo/simple-api-gateway/proto-service/admin"
	protoMail "github.com/aclgo/simple-api-gateway/proto-service/mail"
	protoUser "github.com/aclgo/simple-api-gateway/proto-service/user"
)

var (
	AddrServiceAdmin    = ":50051"
	OptionsServiceAdmin = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	AddrServiceUser    = ":50052"
	OptionsServiceUser = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	AddrServiceMail    = ":50053"
	OptionsServiceMail = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
)

func main() {

	cfg := config.Load(".")

	logger, err := logger.NewapiLogger(cfg)
	if err != nil {
		log.Fatalf("logger.NewapiLogger: %s\n", err)
	}

	logger.Info("logger initialized")

	//	CONNECTING IN MICROSERVICES
	connUser, err := grpc.Dial(AddrServiceUser, OptionsServiceUser...)
	if err != nil {
		logger.Errorf("grpc.Dial: connection in user service: %v", err)
	}

	connAdmin, err := grpc.Dial(AddrServiceAdmin, OptionsServiceAdmin...)
	if err != nil {
		logger.Errorf("grpc.Dial: connection in admin service: %v", err)
	}

	connMail, err := grpc.Dial(AddrServiceMail, OptionsServiceMail...)
	if err != nil {
		logger.Errorf("grpc.Dial: connection in mail service: %v", err)
	}

	redisClient := redis.NewRedisClient(cfg)

	////////////////////////////////

	clientUserService := protoUser.NewUserServiceClient(connUser)
	adminUserService := protoAdmin.NewAdminServiceClient(connAdmin)
	mailUserService := protoMail.NewMailServiceClient(connMail)

	user := userUC.NewuserUC(clientUserService, mailUserService, redisClient, logger)
	admin := adminUC.NewadminUC(adminUserService, mailUserService, redisClient, logger)

	userHandler := svcUser.NewuserService(user, logger)
	adminHandler := svcAdmin.NewadminService(admin, logger)

	authUC := authUC.NewAuthUC(clientUserService)

	ctx := context.Background()

	//MICROSERVICE GRPC USER
	http.HandleFunc("/api/login", userHandler.Login(ctx))
	http.HandleFunc("/api/logout", authUC.ValidateToken(userHandler.Logout(ctx)))
	http.HandleFunc("/api/user/register", userHandler.Register(ctx))
	http.HandleFunc("/api/user/find", authUC.ValidateToken(userHandler.Find(ctx)))
	http.HandleFunc("/api/user/update", authUC.ValidateUpdate(userHandler.Update(ctx)))
	//MICROSERVCE GRPC MAIL
	http.HandleFunc("/api/user/confirm", userHandler.UserConfirm(ctx))
	http.HandleFunc("/api/user/resetpass", userHandler.UserResetPass(ctx))
	http.HandleFunc("/api/user/newpass", userHandler.UserNewPass(ctx))

	//MICROSERVICE GRPC ADMIN
	http.HandleFunc("/api/admin/register", authUC.ValidateCreateAdmin(adminHandler.Create(ctx)))
	http.HandleFunc("/api/admin/search", authUC.ValidateIsAdmin(adminHandler.Search(ctx)))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ApiPort),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		ErrorLog:     log.Default(),
	}

	logger.Infof("server running port %d", cfg.ApiPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("http.ListenAndServe:%v", err)
	}
}
