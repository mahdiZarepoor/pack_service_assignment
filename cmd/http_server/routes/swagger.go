package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetSwaggerRoutes(router *gin.RouterGroup, config configs.Config) {
	if !config.Swagger.Enable {
		return
	}

	docs.SwaggerInfo.Title = config.Swagger.Info.Title
	docs.SwaggerInfo.Description = config.Swagger.Info.Description
	docs.SwaggerInfo.Version = config.Swagger.Info.Version
	docs.SwaggerInfo.Schemes = []string{config.Swagger.Schemes}
	docs.SwaggerInfo.Host = config.Swagger.Host

	authorize := router.Group("/", gin.BasicAuth(gin.Accounts{
		config.Swagger.Username: config.Swagger.Password,
	}))
	authorize.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
