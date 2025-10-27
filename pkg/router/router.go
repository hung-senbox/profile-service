package router

import (
	"profile-service/internal/cache"
	"profile-service/internal/gateway"
	"profile-service/pkg/config"

	"profile-service/internal/profile/route"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"

	owner_code_handler "profile-service/internal/profile/handler"
	owner_code_repo "profile-service/internal/profile/repository"
	owner_code_service "profile-service/internal/profile/service"

	cached_service "profile-service/internal/cache/service"

	goredis "github.com/redis/go-redis/v9"
)

func SetupRouter(consulClient *api.Client, cacheClientRedis *goredis.Client, ownerCodeCol *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Gateway setup
	userGateway := gateway.NewUserGateway("go-main-service", consulClient)

	// cache setup
	systemCache := cache.NewRedisCache(cacheClientRedis)
	cachedUserGateway := cached_service.NewCachedUserGateway(userGateway, systemCache, config.AppConfig.Database.RedisCache.TTLSeconds)

	// OwnerCode
	ownerCodeRepository := owner_code_repo.NewOwnerCodeRepository(ownerCodeCol)
	onwerCodeService := owner_code_service.NewOwnerCodeService(ownerCodeRepository, cachedUserGateway)
	ownerCodeHandler := owner_code_handler.NewOwnerCodeHandler(onwerCodeService)

	// Register routes
	//profile_route.RegisterProfileRoutes(r, ownerCodeHandler, cachedUserGateway)
	route.RegisterProfileRoutes(r, ownerCodeHandler, userGateway)

	return r
}
