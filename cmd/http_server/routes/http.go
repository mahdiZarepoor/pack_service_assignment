package routes

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/cache"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
)

type Master struct {
	cfg         configs.Config
	restHandler HTTP
	logging     logging.Logger
	cache       cache.Interface
}

func NewInstance(cfg configs.Config, logging logging.Logger) *Master {
	return &Master{
		cfg:     cfg,
		logging: logging,
	}
}

func (m *Master) Bootstrap(ctx context.Context) (err error) {

	m.restHandler = NewHttpServer(
		m.cache,
		m.logging,
		m.cfg,
	)

	return nil
}

// Start starts the application master
func (m *Master) Start() {
	m.restHandler.StartBlocking()
}

func provideCache(ctx context.Context, cfg configs.Config) (cache.Interface, error) {
	return cache.NewRedisCacheDriver(ctx, cfg)
}
