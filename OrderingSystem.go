package main

import (
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/GRPC"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/HealthCheck"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/MessageQueue"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/Database"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/API"
	"log"
)

func main() {
	// 配置数据库连接
	dbConfig := Database.NewGormConnection("your-dsn")
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 配置RabbitMQ连接
	rabbitMQConfig := MessageQueue.NewRabbitMQConnectionConfig("your-rabbitmq-connection-config")
	rabbitMQConnection, err := MessageQueue.NewRabbitMQConnection(rabbitMQConfig)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// 配置gRPC服务器
	grpcServerConfig := GRPC.NewGRPCServerConfiguration("localhost", 50051, 5*time.Second, 10*time.Second)
	grpcServer, err := GRPC.NewServer(db, rabbitMQConnection)
	if err != nil {
		log.Fatalf("Failed to create gRPC server: %v", err)
	}

	// 配置API服务器
	apiServerConfig := API.NewAPIServerConfiguration("localhost", 8080)
	apiServer, err := API.NewAPIServer(db, rabbitMQConnection, grpcServer)
	if err != nil {
		log.Fatalf("Failed to create API server: %v", err)
	}

	// 配置健康检查
	healthCheckConfig := HealthCheck.NewHealthCheckConfiguration()
	healthCheckServer, err := HealthCheck.NewHealthCheckServer(healthCheckConfig)
	if err != nil {
		log.Fatalf("Failed to create health check server: %v", err)
	}

	// 启动服务
	go func() {
		log.Println("Starting gRPC server...")
		if err := grpcServer.Start(); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	go func() {
		log.Println("Starting API server...")
		if err := apiServer.Start(); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	go func() {
		log.Println("Starting health check server...")
		if err := healthCheckServer.Start(); err != nil {
			log.Fatalf("Failed to start health check server: %v", err)
		}
	}()

	// 等待程序结束
	select {}
}