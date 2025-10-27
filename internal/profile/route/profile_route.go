package route

import (
	"profile-service/internal/gateway"
	"profile-service/internal/middleware"
	"profile-service/internal/profile/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(r *gin.Engine, ho *handler.OwnerCodeHandler, userGw gateway.UserGateway) {
	// Admin routes
	adminGroup := r.Group("/api/v1/admin")
	adminGroup.Use(middleware.Secured(userGw))
	{
		profilesAdmin := adminGroup.Group("/profiles")
		{
			// owner code routes
			ownerCode := profilesAdmin.Group("/owner-code")
			{
				ownerCode.POST("/student/generate", ho.GenerateStudentCode)
			}
		}
	}
}
