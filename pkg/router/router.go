package router

import (
	"profile-service/internal/gateway"
	"profile-service/pkg/config"

	"profile-service/internal/profile/handler"
	"profile-service/internal/profile/repository"
	"profile-service/internal/profile/route"
	"profile-service/internal/profile/usecase"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"

	profile_service "profile-service/internal/profile/service"

	cache_service "github.com/hung-senbox/senbox-cache-service/pkg/cache"
	profile_caching_service "github.com/hung-senbox/senbox-cache-service/pkg/cache/caching"
)

func SetupRouter(consulClient *api.Client, cacheClientRedis *cache_service.RedisCache, ownerCodeCol, organizationProfileCollection, studentInformationCollection *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Gateway setup
	userGateway := gateway.NewUserGateway("go-main-service", consulClient)

	// cache setup
	cachingProfileService := profile_caching_service.NewCachingProfileService(cacheClientRedis, config.AppConfig.Database.RedisCache.TTLSeconds)

	// Repository
	ownerCodeRepository := repository.NewOwnerCodeRepository(ownerCodeCol)
	organizationProfileRepository := repository.NewOrganizationProfileRepository(organizationProfileCollection)
	studentInformationRepository := repository.NewStudentInformationRepository(studentInformationCollection)

	// usecase
	generateOwnerCodeUseCase := usecase.NewGenerateOwnerCodeUseCase(ownerCodeRepository, cachingProfileService)
	getOwnerCodeUseCase := usecase.NewGetOwnerCodeUseCase(ownerCodeRepository)
	organizationProfileUsecase := usecase.NewOrganizationProfileUsecase(organizationProfileRepository)
	studentInformationUsecase := usecase.NewStudentInformationUsecase(studentInformationRepository)
	// service
	ownerCodeService := profile_service.NewOwnerCodeService(ownerCodeRepository, generateOwnerCodeUseCase, getOwnerCodeUseCase)
	organizationProfileService := profile_service.NewOrganizationProfileService(organizationProfileUsecase)
	studentProfileService := profile_service.NewStudentProfileService(studentInformationUsecase)

	// handler
	ownerCodeHandler := handler.NewOwnerCodeHandler(ownerCodeService)
	organizationProfileHandler := handler.NewOrganizationProfileHandler(organizationProfileService)
	studentProfileHandler := handler.NewStudentProfileHandler(studentProfileService)
	// Register routes
	route.RegisterProfileRoutes(r, ownerCodeHandler, organizationProfileHandler, studentProfileHandler, userGateway)

	return r
}
