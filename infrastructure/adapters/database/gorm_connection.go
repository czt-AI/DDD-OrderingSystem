package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormConnection GORM数据库连接配置
type GormConnection struct {
	DSN string
}

// NewGormConnection 创建新的GORM数据库连接实例
func NewGormConnection(dsn string) *GormConnection {
	return &GormConnection{
		DSN: dsn,
	}
}

// Connect 连接到数据库
func (c *GormConnection) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 设置数据库连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10) // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100) // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	return db, nil
}
