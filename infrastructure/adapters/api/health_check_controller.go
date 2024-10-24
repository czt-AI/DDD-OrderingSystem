package api

import (
	"net/http"
)

// HealthCheckController 健康检查控制器
type HealthCheckController struct{}

// NewHealthCheckController 创建健康检查控制器实例
func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// HealthCheck 健康检查端点
func (hcc *HealthCheckController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// 假设这里有一些逻辑来检查服务的健康状态
	// 例如，检查数据库连接、服务之间的通信等

	// 对于健康检查，我们通常返回200 OK状态码
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
