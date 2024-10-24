package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint64
	Username  string
	Email     string
	Password  string // 密码应进行加密存储
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRole 用户角色枚举
type UserRole string

const (
	// UserRoleAdmin 管理员
	UserRoleAdmin UserRole = "ADMIN"
	// UserRoleCustomer 客户
	UserRoleCustomer UserRole = "CUSTOMER"
)

// NewUser 创建新的用户
func NewUser(username, email, password string, role UserRole) *User {
	return &User{
		ID:        0, // 实际应用中应由数据库生成
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
