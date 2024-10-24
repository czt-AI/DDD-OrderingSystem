package grpc

import (
	"context"
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
	"DDD-OrderingSystem/Infrastructure/Repository"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"userservice"
)

// UserServiceServer gRPC用户服务服务器
type UserServiceServer struct {
	repository Repository.UserRepository
}

// NewUserServiceServer 创建新的UserServiceServer实例
func NewUserServiceServer(repository Repository.UserRepository) *UserServiceServer {
	return &UserServiceServer{
		repository: repository,
	}
}

// CreateUser 创建用户
func (s *UserServiceServer) CreateUser(ctx context.Context, req *userservice.CreateUserRequest) (*userservice.CreateUserResponse, error) {
	// 将请求转换为领域模型
	user := &Model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		Role:      Model.UserRole(req.Role),
	}
	
	// 创建命令
	createUserCommand := Command.NewUserRegisterCommand(user.Username, user.Email, user.Password, user.Role)

	// 处理命令
	if err := Command.Handle(ctx, createUserCommand); err != nil {
		return nil, err
	}

	// 返回响应
	return &userservice.CreateUserResponse{
		UserId: createUserCommand.GetUserID(),
	}, nil
}

// GetUser 获取用户
func (s *UserServiceServer) GetUser(ctx context.Context, req *userservice.GetUserRequest) (*userservice.GetUserResponse, error) {
	user, err := s.repository.FindById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// 将领域模型转换为gRPC响应
	response := &userservice.GetUserResponse{
		UserId: user.ID,
		Username: user.Username,
		Email: user.Email,
		Role: userservice.UserRole(user.Role),
	}

	return response, nil
}

// UpdateUser 更新用户
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *userservice.UpdateUserRequest) (*userservice.UpdateUserResponse, error) {
	// 将请求转换为领域模型
	user := &Model.User{
		ID:        req.UserId,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		Role:      Model.UserRole(req.Role),
	}

	// 创建命令
	updateUserCommand := Command.NewUserRegisterCommand(user.Username, user.Email, user.Password, user.Role)

	// 处理命令
	if err := Command.Handle(ctx, updateUserCommand); err != nil {
		return nil, err
	}

	// 返回响应
	return &userservice.UpdateUserResponse{
		UserId: updateUserCommand.GetUserID(),
	}, nil
}

// DeleteUser 删除用户
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *userservice.DeleteUserRequest) (*userservice.DeleteUserResponse, error) {
	// 创建命令
	deleteUserCommand := Command.NewUserDeleteCommand(req.UserId)

	// 处理命令
	if err := Command.Handle(ctx, deleteUserCommand); err != nil {
		return nil, err
	}

	// 返回响应
	return &userservice.DeleteUserResponse{
		UserId: deleteUserCommand.GetUserID(),
	}, nil
}
