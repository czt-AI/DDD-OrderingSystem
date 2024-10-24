package service

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(ctx context.Context, username, email, password string, role Model.UserRole) error
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	repository Repository.UserRepository
}

// NewUserServiceImpl 创建用户服务实例
func NewUserServiceImpl(repository Repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

// CreateUser 创建用户
func (s *UserServiceImpl) CreateUser(ctx context.Context, username, email, password string, role Model.UserRole) error {
	// 创建用户模型
	user := Model.User{
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 创建命令
	createUserCommand := Command.NewCreateUserCommand(username, email, password, role)

	// 执行命令
	if err := createUserCommand.Execute(ctx); err != nil {
		return err
	}

	// 保存用户到数据库
	if err := s.repository.Save(ctx, &user); err != nil {
		return err
	}

	return nil
}
