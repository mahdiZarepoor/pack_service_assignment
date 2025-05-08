package main

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/cmd/http_server/routes"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"sync"
)

var (
	loggerOnce     sync.Once
	loggerInstance logging.Logger
)

func main() {
	config := configs.GetConfig()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	l := ProvideLogger(config)

	appInstance := routes.NewInstance(config, l)

	// Bootstrap process
	if err := appInstance.Bootstrap(ctx); err != nil {
		cancel()
		loggerInstance.FatalF("failed to bootstrap App instance with error: %v", err)
	}

	// Start Master App (blocking) !
	appInstance.Start()
}

func ProvideLogger(cfg configs.Config) logging.Logger {
	loggerOnce.Do(func() {
		loggerInstance = logging.NewLogger(cfg)
		loggerInstance.Init()

	})
	return loggerInstance
}
