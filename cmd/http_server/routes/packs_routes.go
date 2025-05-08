package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/services"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/driver/http"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/cache"
)

func (s *HttpServer) SetPackRoutes(router *gin.RouterGroup, cache cache.Interface) {
	packSrv := services.NewPackService(s.Config, s.Logging, cache)
	packHdl := http.NewPackHandler(s.Logging, s.Config, packSrv)

	router.GET("", packHdl.List)
	router.PUT("", packHdl.Update)
	router.GET("calculate", packHdl.Calculate)
}
