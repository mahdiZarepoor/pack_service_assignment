package routes

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/cmd/http_server/middlewares"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/cache"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"net/http"
)

const (
	ReleaseMode = "release"
	DebugMode   = "debug"
)

type HTTP interface {
	StartBlocking()
}

type HttpServer struct {
	Cache      cache.Interface
	Logging    logging.Logger
	Config     configs.Config
	httpServer *http.Server
}

func NewHttpServer(
	cache cache.Interface,
	logging logging.Logger,
	config configs.Config,
) *HttpServer {

	return &HttpServer{
		Cache:   cache,
		Logging: logging,
		Config:  config,
	}
}

func (s *HttpServer) registerRoutes() http.Handler {
	mode := ReleaseMode
	if s.Config.App.Debug {
		mode = DebugMode
	}

	gin.SetMode(mode)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler))
	engine.Use(middlewares.DefaultStructuredLogger(s.Config))

	s.setHealthRoutes(engine.Group(""))
	SetSwaggerRoutes(engine.Group(""), s.Config)

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			s.SetPackRoutes(v1.Group("packs"), s.Cache)
		}

	}

	return engine
}

func (s *HttpServer) StartBlocking() {
	address := fmt.Sprintf(":%s", s.Config.App.Port)

	// Initialize the http.Server
	s.httpServer = &http.Server{
		Addr:    address,
		Handler: s.registerRoutes(),
	}

	s.Logging.Info(logging.App, logging.Bootstrapping, fmt.Sprintf("HTTP server started: %s", s.httpServer.Addr), nil)

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.Logging.Fatal(logging.App, logging.Bootstrapping, fmt.Sprintf("GinHTTP REST Server failed to start with error: %v", err), nil)
	}
}
