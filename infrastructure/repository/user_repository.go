package repository

import (
	"DDD-OrderingSystem/Domain/Model"
	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	Save(ctx context.Context, user *Model.User) error
	FindById(ctx context.Context, id uint64) (*Model.User, error)
	FindAll(ctx context.Context) ([]Model.User, error)
}

// UserRepositoryImpl 用户仓库实现
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Save 保存用户
func (r *UserRepositoryImpl) Save(ctx context.Context, user *Model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindById 通过ID查找用户
func (r *UserRepositoryImpl) FindById(ctx context.Context, id uint64) (*Model.User, error) {
	var user Model.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll 查找所有用户
func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]Model.User, error) {
	var users []Model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
