package service

import (
	"DDD-OrderingSystem/Infrastructure/Adapters/HealthCheck"
)

// HealthCheckService 健康检查服务接口
type HealthCheckService interface {
	CheckHealth() bool
}

// HealthCheckServiceImpl 健康检查服务实现
type HealthCheckServiceImpl struct {
	healthChecker HealthCheck.HealthChecker
}

// NewHealthCheckServiceImpl 创建健康检查服务实例
func NewHealthCheckServiceImpl(healthChecker HealthCheck.HealthChecker) *HealthCheckServiceImpl {
	return &HealthCheckServiceImpl{
		healthChecker: healthChecker,
	}
}

// CheckHealth 检查服务健康状态
func (s *HealthCheckServiceImpl) CheckHealth() bool {
	return s.healthChecker.IsHealthy()
}
