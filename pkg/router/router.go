package router

import (
	"profile-service/internal/cache"
	"profile-service/internal/gateway"
	holiday_handler "profile-service/internal/holiday/handler"
	holiday_repo "profile-service/internal/holiday/repository"
	holiday_route "profile-service/internal/holiday/route"
	holiday_service "profile-service/internal/holiday/service"
	"profile-service/internal/term/handler"
	"profile-service/internal/term/repository"
	"profile-service/internal/term/route"
	"profile-service/internal/term/service"
	"profile-service/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"

	cached_service "profile-service/internal/cache/service"

	goredis "github.com/redis/go-redis/v9"
)

func SetupRouter(consulClient *api.Client, cacheClientRedis *goredis.Client, termCollection, holidayCollection *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Gateway setup
	userGateway := gateway.NewUserGateway("go-main-service", consulClient)
	orgGateway := gateway.NewOrganizationGateway("go-main-service", consulClient)
	messageLanguageGW := gateway.NewMessageLanguageGateway("go-main-service", consulClient)

	// cache setup
	systemCache := cache.NewRedisCache(cacheClientRedis)
	cachedUserGateway := cached_service.NewCachedUserGateway(userGateway, systemCache, config.AppConfig.Database.RedisCache.TTLSeconds)

	// Term
	termRepo := repository.NewTermRepository(termCollection)
	termSvc := service.NewTermService(termRepo, cachedUserGateway, orgGateway, messageLanguageGW)
	termHandler := handler.NewHandler(termSvc)

	// Holiday
	holidayRepo := holiday_repo.NewHolidayRepository(holidayCollection)
	holidaySvc := holiday_service.NewHolidayService(holidayRepo, userGateway, orgGateway, messageLanguageGW)
	holidayHandler := holiday_handler.NewHandler(holidaySvc)

	// Register routes
	route.RegisterTermRoutes(r, termHandler, userGateway)
	holiday_route.RegisterHolidayRoutes(r, holidayHandler, userGateway)

	return r
}
