package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// Config 全局配置结构
type Config struct {
	Database DatabaseConfig
	GRPC     GRPCConfig
	API      APIConfig
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

// GRPCConfig gRPC配置
type GRPCConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	KeepAliveTime time.Duration `yaml:"keep_alive_time"`
	KeepAliveTimeout time.Duration `yaml:"keep_alive_timeout"`
}

// APIConfig API配置
type APIConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
}

// Load 加载配置文件
func Load(filePath string) (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Database.DSN == "" {
		return fmt.Errorf("database DSN is required")
	}
	if c.GRPC.Host == "" || c.GRPC.Port == 0 {
		return fmt.Errorf("gRPC host and port are required")
	}
	if c.API.Host == "" || c.API.Port == 0 {
		return fmt.Errorf("API host and port are required")
	}
	return nil
}
