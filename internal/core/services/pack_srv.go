package services

import (
	"context"
	"encoding/json"

	"errors"

	"github.com/go-redis/redis"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/consts"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/ports/packs_port"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/cache"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/packs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/response"
)

type packService struct {
	cfg    configs.Config
	logger logging.Logger
	cache  cache.Interface
}

func NewPackService(cfg configs.Config, l logging.Logger, cache cache.Interface) packs_port.IPackService {
	return &packService{
		cfg:    cfg,
		logger: l,
		cache:  cache,
	}
}

func (s *packService) Update(ctx context.Context, packSizes []uint) response.Error {
	if s.cache == nil {
		return response.NewServiceError(consts.CacheNotInitialized)
	}

	err := s.cache.Set(ctx, "pack_sizes", packSizes, 0)
	if err != nil {
		s.logger.Error(logging.Redis, logging.InternalError, "Failed to store pack sizes in cache", map[logging.ExtraKey]interface{}{
			logging.ErrorMessage: err.Error(),
		})
		return response.NewServiceError(consts.ServerError)
	}

	return nil
}

func (s *packService) List(ctx context.Context) ([]uint, response.Error) {
	if s.cache == nil {
		return nil, response.NewServiceError(consts.CacheNotInitialized)
	}

	packSizes, err := s.cache.Get(ctx, "pack_sizes")
	if err != nil {
		s.logger.Error(logging.Redis, logging.InternalError, "Debug error info", map[logging.ExtraKey]interface{}{
			"error_type":   err.Error(),
			"error_string": err.Error(),
			"is_redis_nil": errors.Is(err, redis.Nil),
		})
		if errors.Is(err, redis.Nil) || err.Error() == "redis: nil" {
			return []uint{}, nil
		}
		s.logger.Error(logging.Redis, logging.InternalError, "Failed to retrieve pack sizes from cache", map[logging.ExtraKey]interface{}{
			logging.ErrorMessage: err.Error(),
		})
		return nil, response.NewServiceError(consts.ServerError)
	}

	var result []uint
	if err := json.Unmarshal(packSizes, &result); err != nil {
		return nil, response.NewServiceError(consts.ServerError)
	}

	return result, nil
}

func (s *packService) Calculate(ctx context.Context, total uint) (map[int]int, response.Error) {
	if s.cache == nil {
		return nil, response.NewServiceError(consts.CacheNotInitialized)
	}

	// Get pack sizes from cache
	packSizes, err := s.cache.Get(ctx, "pack_sizes")
	if err != nil {
		s.logger.Error(logging.Redis, logging.InternalError, "Debug error info", map[logging.ExtraKey]interface{}{
			"error_type":   err.Error(),
			"error_string": err.Error(),
			"is_redis_nil": errors.Is(err, redis.Nil),
		})
		if errors.Is(err, redis.Nil) || err.Error() == "redis: nil" {
			return nil, response.NewServiceError(consts.RecordNotFound)
		}
		s.logger.Error(logging.Redis, logging.InternalError, "Failed to retrieve pack sizes from cache", map[logging.ExtraKey]interface{}{
			logging.ErrorMessage: err.Error(),
		})
		return nil, response.NewServiceError(consts.ServerError)
	}

	var result []int
	if err := json.Unmarshal(packSizes, &result); err != nil {
		return nil, response.NewServiceError(consts.ServerError)
	}

	// Calculate optimal pack distribution
	r := packs.Calculate(result, int(total))
	if result == nil {
		return nil, response.NewServiceError(consts.ServerError)
	}

	return r, nil
}
