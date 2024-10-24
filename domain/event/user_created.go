package event

import (
	"DDD-OrderingSystem/Domain/Model"
	"time"
)

// UserCreatedEvent 用户创建事件
type UserCreatedEvent struct {
	UserID    uint64
	Username  string
	Email     string
	Role      Model.UserRole
	CreatedAt time.Time
}

// NewUserCreatedEvent 创建用户创建事件
func NewUserCreatedEvent(userID uint64, username, email string, role Model.UserRole) *UserCreatedEvent {
	return &UserCreatedEvent{
		UserID:    userID,
		Username:  username,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
}
