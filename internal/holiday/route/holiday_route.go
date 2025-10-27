package route

import (
	"profile-service/internal/gateway"
	"profile-service/internal/holiday/handler"
	"profile-service/internal/term/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterHolidayRoutes(r *gin.Engine, h *handler.HolidayHandler, userGw gateway.UserGateway) {
	// Admin routes
	adminGroup := r.Group("/api/v1/admin")
	adminGroup.Use(middleware.Secured(userGw))
	{
		holidaysAdmin := adminGroup.Group("/holidays")
		{
			holidaysAdmin.POST("", h.UploadHolidays)
			holidaysAdmin.GET("", h.GetHolidays4Web)
		}
	}
}
