package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/controller/handlers"
)

func (s *HttpServer) setHealthRoutes(router *gin.RouterGroup) {
	Handler := handlers.NewHealthHandler()
	router.GET("/health", Handler.HealthCheck)
}
