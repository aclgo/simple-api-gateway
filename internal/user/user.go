package user

import (
	"context"

	userService "github.com/aclgo/simple-api-gateway/proto-services/user"
	"google.golang.org/grpc"
)

type UserService struct {
	userServiceClient userService.UserServiceClient
}

func NewUser(userServiceClient userService.UserServiceClient) *UserService {
	return &UserService{
		userServiceClient: userServiceClient,
	}
}

func (u *UserService) Register(ctx context.Context, in *userService.CreateUserRequest, opts ...grpc.CallOption) (error, error) {
	return nil, nil
}
func (u *UserService) Login(ctx context.Context, in *userService.UserLoginRequest, opts ...grpc.CallOption) (error, error) {
	return nil, nil
}
func (u *UserService) Logout(ctx context.Context, in *userService.UserLogoutRequest, opts ...grpc.CallOption) (error, error) {
	return nil, nil
}
func (u *UserService) FindById(ctx context.Context, in *userService.FindByIdRequest, opts ...grpc.CallOption) (error, error) {
	return nil, nil
}
func (u *UserService) FindByEmail(ctx context.Context, in *userService.FindByEmailRequest, opts ...grpc.CallOption) (error, error) {
	return nil, nil
}
func (u *UserService) Update(ctx context.Context, in *userService.UpdateRequest, opts ...grpc.CallOption) (error, error) {
	return nil, nil
}
