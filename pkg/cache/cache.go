package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Interface interface {
	Instance() *redis.Client
	Get(ctx context.Context, key string) (destination []byte, err error)
	Set(ctx context.Context, key string, value any, expiration time.Duration) (err error)
	Delete(ctx context.Context, key string) (err error)
	FlushAll(ctx context.Context)
	Stop() error
}
