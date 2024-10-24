package service

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
	"DDD-OrderingSystem/Domain/Query"
)

// UserService 用户服务接口
type UserService interface {
	RegisterUser(ctx context.Context, username, email, password string, role Model.UserRole) error
	LoginUser(ctx context.Context, username, password string) (string, error)
	GetUserDetails(ctx context.Context, userID uint64) (*Model.User, error)
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	commandHandler Command.Handler
	queryHandler  Query.Handler
	repository    Repository.UserRepository
}

// NewUserServiceImpl 创建用户服务实例
func NewUserServiceImpl(commandHandler Command.Handler, queryHandler Query.Handler, repository Repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		commandHandler: commandHandler,
		queryHandler:  queryHandler,
		repository:    repository,
	}
}

// RegisterUser 注册用户
func (s *UserServiceImpl) RegisterUser(ctx context.Context, username, email, password string, role Model.UserRole) error {
	// 创建命令
	createUserCommand := Command.NewUserRegisterCommand(username, email, password, role)

	// 处理命令
	if err := s.commandHandler.Handle(ctx, createUserCommand); err != nil {
		return err
	}

	// 获取用户ID
	userID := createUserCommand.GetUserID()

	// 查询用户详情
	user, err := s.queryHandler.Handle(ctx, Query.NewUserDetailsQuery(userID))
	if err != nil {
		return err
	}

	return nil
}

// LoginUser 用户登录
func (s *UserServiceImpl) LoginUser(ctx context.Context, username, password string) (string, error) {
	// 创建命令
	loginCommand := Command.NewUserLoginCommand(username, password)

	// 处理命令
	if err := s.commandHandler.Handle(ctx, loginCommand); err != nil {
		return "", err
	}

	// 获取JWT令牌
	token := loginCommand.GetToken()

	return token, nil
}

// GetUserDetails 获取用户详情
func (s *UserServiceImpl) GetUserDetails(ctx context.Context, userID uint64) (*Model.User, error) {
	return s.queryHandler.Handle(ctx, Query.NewUserDetailsQuery(userID))
}
