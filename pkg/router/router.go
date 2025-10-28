package router

import (
	"profile-service/internal/cache"
	"profile-service/internal/gateway"
	"profile-service/pkg/config"

	"profile-service/internal/profile/handler"
	"profile-service/internal/profile/repository"
	"profile-service/internal/profile/route"
	"profile-service/internal/profile/usecase"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"

	"profile-service/internal/cache/caching"
	profile_service "profile-service/internal/profile/service"

	goredis "github.com/redis/go-redis/v9"
)

func SetupRouter(consulClient *api.Client, cacheClientRedis *goredis.Client, ownerCodeCol *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Gateway setup
	userGateway := gateway.NewUserGateway("go-main-service", consulClient)

	// cache setup
	systemCache := cache.NewRedisCache(cacheClientRedis)
	cachingProfileService := caching.NewCachingService(systemCache, config.AppConfig.Database.RedisCache.TTLSeconds)

	// Repository
	ownerCodeRepository := repository.NewOwnerCodeRepository(ownerCodeCol)

	// usecase
	generateOwnerCodeUseCase := usecase.NewGenerateOwnerCodeUseCase(ownerCodeRepository, cachingProfileService)
	getOwnerCodeUseCase := usecase.NewGetOwnerCodeUseCase(ownerCodeRepository)

	// service
	onwerCodeService := profile_service.NewOwnerCodeService(ownerCodeRepository, generateOwnerCodeUseCase, getOwnerCodeUseCase)

	// handler
	ownerCodeHandler := handler.NewOwnerCodeHandler(onwerCodeService)

	// Register routes
	route.RegisterProfileRoutes(r, ownerCodeHandler, userGateway)

	return r
}
