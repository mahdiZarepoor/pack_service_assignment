package service

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/cmd/app/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/consts"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/response"
	"sort"
)

var packSizes []int

type packService struct {
	cfg    configs.Config
	logger logging.Logger
}

func NewPackService(cfg configs.Config, l logging.Logger) IPackService {
	packSizes = make([]int, 0)

	return &packService{
		cfg:    cfg,
		logger: l,
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
	r := CalculatePacks(packSizes, total)
	if r == nil {
		return nil, response.NewServiceError(consts.ServerError)
	}

	return r, nil
}

func CalculatePacks(packSizes []int, order int) map[int]int {
	if order <= 0 {
		return nil
	}

	sort.Ints(packSizes)

	if order <= packSizes[0] {
		return map[int]int{
			packSizes[0]: 1,
		}
	}

	dp := make([]int, order+packSizes[len(packSizes)-1]+1)
	prev := make([]int, order+packSizes[len(packSizes)-1]+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	minValidSum := -1
	for i := 1; i < len(dp); i++ {
		for _, pack := range packSizes {
			if i >= pack && dp[i-pack] != -1 {
				newCount := dp[i-pack] + 1
				if dp[i] == -1 || newCount < dp[i] {
					dp[i] = newCount
					prev[i] = pack
				}
			}
		}
		if i >= order && dp[i] != -1 {
			minValidSum = i
			break
		}
	}

	if minValidSum == -1 {
		return nil
	}

	result := make(map[int]int)
	current := minValidSum
	for current > 0 {
		pack := prev[current]
		result[pack]++
		current -= pack
	}

	return result
}
