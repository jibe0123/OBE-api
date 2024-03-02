package v1

import (
	"api/internal/app/controller"
	service "api/internal/app/services"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for API v1.
func RegisterRoutes(r *gin.RouterGroup) {
	csvService := service.NewCsvService()
	collectController := controller.NewCollectController(csvService)

	r.POST("/collect", collectController.CollectData)
}
