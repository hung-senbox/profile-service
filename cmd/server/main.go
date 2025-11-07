package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// "os"

	"profile-service/pkg/config"
	"profile-service/pkg/consul"
	"profile-service/pkg/db"
	"profile-service/pkg/router"

	"profile-service/pkg/zap"

	consulapi "github.com/hashicorp/consul/api"
	redis "github.com/hung-senbox/senbox-cache-service/pkg/redis"
)

func main() {
	filePath := os.Args[1]
	if filePath == "" {
		filePath = "configs/config.yaml"
	}

	config.LoadConfig(filePath)

	cfg := config.AppConfig

	// logger.WriteLogData("info", map[string]any{"id": 123, "name": "Hung"})

	//logger
	logger, err := zap.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	//consul
	consulConn := consul.NewConsulConn(logger, cfg)
	consulClient := consulConn.Connect()
	defer consulConn.Deregister()

	if err := waitPassing(consulClient, "go-main-service", 60*time.Second); err != nil {
		logger.Fatalf("Dependency not ready: %v", err)
	}

	//db
	db.ConnectMongoDB()

	// redis cache
	cacheClientRedis, err := redis.InitRedisCache(config.AppConfig.Database.RedisCache.Host, config.AppConfig.Database.RedisCache.Port, config.AppConfig.Database.RedisCache.Password, config.AppConfig.Database.RedisCache.DB)
	if err != nil {
		logger.Fatalf("Failed to initialize Redis cache: %v", err)
	}

	r := router.SetupRouter(consulClient, cacheClientRedis, db.OwnerCodeCollection, db.OrganizationProfileCollection)
	port := cfg.Server.Port
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

func waitPassing(cli *consulapi.Client, name string, timeout time.Duration) error {
	dl := time.Now().Add(timeout)
	for time.Now().Before(dl) {
		entries, _, err := cli.Health().Service(name, "", true, nil)
		if err == nil && len(entries) > 0 {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("%s not ready in consul", name)
}
