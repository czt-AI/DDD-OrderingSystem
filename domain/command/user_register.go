package command

import (
	"DDD-OrderingSystem/Domain/Model"
)

// UserRegisterCommand 用户注册命令
type UserRegisterCommand struct {
	Username  string
	Email     string
	Password  string
	Role      Model.UserRole
}

// NewUserRegisterCommand 创建用户注册命令
func NewUserRegisterCommand(username, email, password string, role Model.UserRole) *UserRegisterCommand {
	return &UserRegisterCommand{
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      role,
	}
}

// Execute 执行命令
func (c *UserRegisterCommand) Execute(ctx context.Context) error {
	// 这里可以添加用户注册的逻辑，例如验证用户信息，创建用户账户等

	// 创建用户模型
	user := Model.User{
		Username:  c.Username,
		Email:     c.Email,
		Password:  c.Password,
		Role:      c.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 应用事件
	// 这里应该将事件保存到事件存储中，或者通过其他方式发布

	return nil
}

// GetUserID 获取用户ID
func (c *UserRegisterCommand) GetUserID() uint64 {
	// 这里应该返回用户ID，实际应用中这个值应该由执行命令的逻辑来设置
	return 0
}
