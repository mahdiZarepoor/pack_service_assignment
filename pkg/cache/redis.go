package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheDriver struct {
	client *redis.Client
	config configs.Config
}

func NewRedisCacheDriver(
	ctx context.Context,
	config configs.Config,
) (Interface, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		DialTimeout:  config.Redis.DialTimeout * time.Second,
		ReadTimeout:  config.Redis.ReadTimeout * time.Second,
		WriteTimeout: config.Redis.WriteTimeout * time.Second,
		PoolSize:     config.Redis.PoolSize,
		PoolTimeout:  config.Redis.PoolTimeout * time.Second,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCacheDriver{
		client: client,
		config: config,
	}, nil
}

func (r *RedisCacheDriver) Instance() *redis.Client {
	return r.client
}

func (r *RedisCacheDriver) Stop() error {
	if err := r.client.Close(); err != nil {
		return err
	}
	return nil
}

func (r *RedisCacheDriver) Get(ctx context.Context, key string) (destination []byte, err error) {
	key = fmt.Sprintf("%s:%s", r.config.Redis.Prefix, key)
	v, err := r.client.Get(ctx, key).Result()

	if err != nil {
		return destination, err
	}

	return []byte(v), nil
}

func (r *RedisCacheDriver) Set(ctx context.Context, key string, value any, expiration time.Duration) (err error) {
	key = fmt.Sprintf("%s:%s", r.config.Redis.Prefix, key)

	data, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *RedisCacheDriver) Delete(ctx context.Context, key string) (err error) {
	key = fmt.Sprintf("%s:%s", r.config.Redis.Prefix, key)

	return r.client.Del(ctx, key).Err()
}

func (r *RedisCacheDriver) FlushAll(ctx context.Context) {
	if r.config.App.Debug {
		r.client.FlushAll(ctx)
	}
}
