package integrationevents

import (
	"DDD-OrderingSystem/Domain/Model"
	"time"
)

// UserCreatedIntegrationEvent 用户创建集成事件
type UserCreatedIntegrationEvent struct {
	UserID    uint64
	Username  string
	Email     string
	Role      Model.UserRole
	CreatedAt time.Time
}

// NewUserCreatedIntegrationEvent 创建用户创建集成事件
func NewUserCreatedIntegrationEvent(userID uint64, username, email string, role Model.UserRole) *UserCreatedIntegrationEvent {
	return &UserCreatedIntegrationEvent{
		UserID:    userID,
		Username:  username,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
}
