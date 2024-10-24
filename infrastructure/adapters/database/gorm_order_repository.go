package database

import (
	"DDD-OrderingSystem/Domain/Model"
	"gorm.io/gorm"
)

// GormOrderRepository 使用GORM的订单仓库实现
type GormOrderRepository struct {
	db *gorm.DB
}

// NewGormOrderRepository 创建新的GORM订单仓库实例
func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{
		db: db,
	}
}

// Save 保存订单
func (r *GormOrderRepository) Save(ctx context.Context, order *Model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

// FindById 通过ID查找订单
func (r *GormOrderRepository) FindById(ctx context.Context, id uint64) (*Model.Order, error) {
	var order Model.Order
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// FindAll 查找所有订单
func (r *GormOrderRepository) FindAll(ctx context.Context) ([]Model.Order, error) {
	var orders []Model.Order
	if err := r.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
