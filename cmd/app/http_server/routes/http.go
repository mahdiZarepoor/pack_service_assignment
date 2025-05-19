package routes

import (
	"context"
<<<<<<< Updated upstream:cmd/http_server/routes/http.go
	"fmt"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/cache"
=======
	"github.com/mahdiZarepoor/pack_service_assignment/cmd/app/configs"
>>>>>>> Stashed changes:cmd/app/http_server/routes/http.go
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
)

type Master struct {
	cfg         configs.Config
	restHandler HTTP
	logging     logging.Logger
}

func NewInstance(cfg configs.Config, logging logging.Logger) *Master {
	return &Master{
		cfg:     cfg,
		logging: logging,
	}
}

func (m *Master) Bootstrap(ctx context.Context) (err error) {

	if m.cache, err = provideCache(ctx, m.cfg); err != nil {
		m.logging.Error(logging.Redis, logging.RedisInit, fmt.Sprintf("Failed to initialize Cache %s", err.Error()), nil)
		return
	}

	m.restHandler = NewHttpServer(
		m.logging,
		m.cfg,
	)

	return nil
}

// Start starts the application master
func (m *Master) Start() {
	m.restHandler.StartBlocking()
}
