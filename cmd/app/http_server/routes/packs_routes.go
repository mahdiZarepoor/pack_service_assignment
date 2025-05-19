package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/controller/handlers"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/service"
)

func (s *HttpServer) SetPackRoutes(router *gin.RouterGroup) {
	packSrv := service.NewPackService(s.Config, s.Logging)
	packHdl := handlers.NewPackHandler(s.Logging, s.Config, packSrv)

	router.GET("", packHdl.List)
	router.PUT("", packHdl.Update)
	router.GET("calculate", packHdl.Calculate)
}
