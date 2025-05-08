package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/driver/http"
)

func (s *HttpServer) setHealthRoutes(router *gin.RouterGroup) {
	Handler := http.NewHealthHandler()
	router.GET("/health", Handler.HealthCheck)
}
