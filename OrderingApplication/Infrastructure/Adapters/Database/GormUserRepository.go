package database

import (
	"DDD-OrderingSystem/OrderingApplication/Domain/Model"
	"gorm.io/gorm"
)

// GormUserRepository 使用GORM的用户仓库实现
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository 创建新的GORM用户仓库实例
func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

// Save 保存用户
func (r *GormUserRepository) Save(ctx context.Context, user *Model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindById 通过ID查找用户
func (r *GormUserRepository) FindById(ctx context.Context, id uint64) (*Model.User, error) {
	var user Model.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll 查找所有用户
func (r *GormUserRepository) FindAll(ctx context.Context) ([]Model.User, error) {
	var users []Model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
