package services

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/consts"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/ports/packs_port"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/cache"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/packs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/response"
)

var packSizes []int

type packService struct {
	cfg    configs.Config
	logger logging.Logger
	cache  cache.Interface
}

func NewPackService(cfg configs.Config, l logging.Logger, cache cache.Interface) packs_port.IPackService {
	packSizes = make([]int, 0)

	return &packService{
		cfg:    cfg,
		logger: l,
		cache:  cache,
	}
}

func (s *packService) Update(ctx context.Context, p []int) response.Error {
	packSizes = p
	return nil
}

func (s *packService) List(ctx context.Context) ([]int, response.Error) {
	if len(packSizes) == 0 {
		return nil, response.NewServiceError(consts.RecordNotFound)
	}

	return packSizes, nil
}

func (s *packService) Calculate(ctx context.Context, total int) (map[int]int, response.Error) {
	if len(packSizes) == 0 {
		return nil, response.NewServiceError(consts.RecordNotFound)
	}

	// Calculate optimal pack distribution
	r := packs.Calculate(packSizes, total)
	if r == nil {
		return nil, response.NewServiceError(consts.ServerError)
	}

	return r, nil
}
