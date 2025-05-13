package packs_port

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/response"
)

type IPackService interface {
	Update(ctx context.Context, packSizes []int) response.Error
	List(ctx context.Context) ([]int, response.Error)
	Calculate(ctx context.Context, total int) (map[int]int, response.Error)
}
