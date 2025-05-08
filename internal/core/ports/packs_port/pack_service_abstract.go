package packs_port

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/response"
)

type IPackService interface {
	Update(ctx context.Context, packSizes []uint) response.Error
	List(ctx context.Context) ([]uint, response.Error)
	Calculate(ctx context.Context, total uint) (map[int]int, response.Error)
}
