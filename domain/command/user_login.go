package command

import (
	"DDD-OrderingSystem/Domain/Model"
)

// UserLoginCommand 用户登录命令
type UserLoginCommand struct {
	Username string
	Password string
}

// NewUserLoginCommand 创建用户登录命令
func NewUserLoginCommand(username, password string) *UserLoginCommand {
	return &UserLoginCommand{
		Username: username,
		Password: password,
	}
}

// Execute 执行命令
func (c *UserLoginCommand) Execute(ctx context.Context) error {
	// 这里可以添加用户登录的逻辑，例如验证用户名和密码，生成JWT令牌等

	// 创建用户模型
	user, err := Model.FindUserByUsername(ctx, c.Username)
	if err != nil {
		return err
	}

	// 验证密码
	if user.Password != c.Password {
		return fmt.Errorf("invalid username or password")
	}

	// 生成JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"iat":      time.Now().Unix(),
	})

	signedToken, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return err
	}

	// 返回令牌
	return nil
}

// GetToken 获取生成的JWT令牌
func (c *UserLoginCommand) GetToken() string {
	// 返回生成的JWT令牌
	return "your_generated_token"
}
