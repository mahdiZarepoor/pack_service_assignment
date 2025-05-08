package services

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/ports/packs_port"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCache is a mock implementation of cache.Interface
type MockCache struct {
	mock.Mock
}

func (m *MockCache) Instance() *redis.Client {
	args := m.Called()
	return args.Get(0).(*redis.Client)
}

func (m *MockCache) Get(ctx context.Context, key string) ([]byte, error) {
	args := m.Called(ctx, key)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockCache) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *MockCache) FlushAll(ctx context.Context) {
	m.Called(ctx)
}

func (m *MockCache) Stop() error {
	args := m.Called()
	return args.Error(0)
}

// MockLogger is a mock implementation of logging.Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Init() {
	m.Called()
}

func (m *MockLogger) Debug(category logging.Category, subCategory logging.SubCategory, message string, extra map[logging.ExtraKey]interface{}) {
	m.Called(category, subCategory, message, extra)
}

func (m *MockLogger) DebugF(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *MockLogger) Info(category logging.Category, subCategory logging.SubCategory, message string, extra map[logging.ExtraKey]interface{}) {
	m.Called(category, subCategory, message, extra)
}

func (m *MockLogger) InfoF(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *MockLogger) Warn(category logging.Category, subCategory logging.SubCategory, message string, extra map[logging.ExtraKey]interface{}) {
	m.Called(category, subCategory, message, extra)
}

func (m *MockLogger) WarnF(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *MockLogger) Error(category logging.Category, subCategory logging.SubCategory, message string, extra map[logging.ExtraKey]interface{}) {
	m.Called(category, subCategory, message, extra)
}

func (m *MockLogger) ErrorF(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *MockLogger) Fatal(category logging.Category, subCategory logging.SubCategory, message string, extra map[logging.ExtraKey]interface{}) {
	m.Called(category, subCategory, message, extra)
}

func (m *MockLogger) FatalF(template string, args ...interface{}) {
	m.Called(template, args)
}

func TestPackService_Update(t *testing.T) {
	tests := []struct {
		name      string
		packSizes []uint
		mockSetup func(*MockCache, *MockLogger)
		wantErr   bool
	}{
		{
			name:      "successful update",
			packSizes: []uint{250, 500, 1000, 2000, 5000},
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				mc.On("Set", mock.Anything, "pack_sizes", []uint{250, 500, 1000, 2000, 5000}, time.Duration(0)).Return(nil)
			},
			wantErr: false,
		},
		{
			name:      "cache error",
			packSizes: []uint{250, 500, 1000, 2000, 5000},
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				mc.On("Set", mock.Anything, "pack_sizes", []uint{250, 500, 1000, 2000, 5000}, time.Duration(0)).Return(redis.ErrClosed)
				ml.On("Error", logging.Redis, logging.InternalError, "Failed to store pack sizes in cache", mock.Anything).Return()
			},
			wantErr: true,
		},
		{
			name:      "nil cache",
			packSizes: []uint{250, 500, 1000, 2000, 5000},
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				// No setup needed for nil cache
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCache := new(MockCache)
			mockLogger := new(MockLogger)
			cfg := configs.Config{}

			if tt.name != "nil cache" {
				tt.mockSetup(mockCache, mockLogger)
			}

			var service packs_port.IPackService
			if tt.name == "nil cache" {
				service = NewPackService(cfg, mockLogger, nil)
			} else {
				service = NewPackService(cfg, mockLogger, mockCache)
			}
			err := service.Update(context.Background(), tt.packSizes)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			mockCache.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}

func TestPackService_List(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*MockCache, *MockLogger)
		want      []uint
		wantErr   bool
	}{
		{
			name: "successful list",
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				data, _ := json.Marshal([]uint{250, 500, 1000, 2000, 5000})
				mc.On("Get", mock.Anything, "pack_sizes").Return(data, nil)
			},
			want:    []uint{250, 500, 1000, 2000, 5000},
			wantErr: false,
		},
		{
			name: "empty list",
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				mc.On("Get", mock.Anything, "pack_sizes").Return([]byte{}, redis.Nil)
				ml.On("Error", logging.Redis, logging.InternalError, "Debug error info", mock.MatchedBy(func(extra map[logging.ExtraKey]interface{}) bool {
					return true
				})).Return()
			},
			want:    []uint{},
			wantErr: false,
		},
		{
			name: "cache error",
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				mc.On("Get", mock.Anything, "pack_sizes").Return([]byte{}, redis.ErrClosed)
				ml.On("Error", logging.Redis, logging.InternalError, "Debug error info", mock.MatchedBy(func(extra map[logging.ExtraKey]interface{}) bool {
					return true
				})).Return()
				ml.On("Error", logging.Redis, logging.InternalError, "Failed to retrieve pack sizes from cache", mock.MatchedBy(func(extra map[logging.ExtraKey]interface{}) bool {
					return true
				})).Return()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil cache",
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				// No setup needed for nil cache
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCache := new(MockCache)
			mockLogger := new(MockLogger)
			cfg := configs.Config{}

			if tt.name != "nil cache" {
				tt.mockSetup(mockCache, mockLogger)
			}

			var service packs_port.IPackService
			if tt.name == "nil cache" {
				service = NewPackService(cfg, mockLogger, nil)
			} else {
				service = NewPackService(cfg, mockLogger, mockCache)
			}

			got, err := service.List(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockCache.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}

func TestPackService_Calculate(t *testing.T) {
	tests := []struct {
		name      string
		total     uint
		mockSetup func(*MockCache, *MockLogger)
		want      map[int]int
		wantErr   bool
	}{
		{
			name:  "successful calculation",
			total: 1200,
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				data, _ := json.Marshal([]int{250, 500, 1000, 2000, 5000})
				mc.On("Get", mock.Anything, "pack_sizes").Return(data, nil)
			},
			want:    map[int]int{1000: 1, 250: 1},
			wantErr: false,
		},
		{
			name:  "cache not found",
			total: 1200,
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				mc.On("Get", mock.Anything, "pack_sizes").Return([]byte{}, redis.Nil)
				ml.On("Error", logging.Redis, logging.InternalError, "Debug error info", mock.MatchedBy(func(extra map[logging.ExtraKey]interface{}) bool {
					return true
				})).Return()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "cache error",
			total: 1200,
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				mc.On("Get", mock.Anything, "pack_sizes").Return([]byte{}, redis.ErrClosed)
				ml.On("Error", logging.Redis, logging.InternalError, "Debug error info", mock.MatchedBy(func(extra map[logging.ExtraKey]interface{}) bool {
					return true
				})).Return()
				ml.On("Error", logging.Redis, logging.InternalError, "Failed to retrieve pack sizes from cache", mock.MatchedBy(func(extra map[logging.ExtraKey]interface{}) bool {
					return true
				})).Return()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "nil cache",
			total: 1200,
			mockSetup: func(mc *MockCache, ml *MockLogger) {
				// No setup needed for nil cache
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCache := new(MockCache)
			mockLogger := new(MockLogger)
			cfg := configs.Config{}

			if tt.name != "nil cache" {
				tt.mockSetup(mockCache, mockLogger)
			}

			var service packs_port.IPackService
			if tt.name == "nil cache" {
				service = NewPackService(cfg, mockLogger, nil)
			} else {
				service = NewPackService(cfg, mockLogger, mockCache)
			}

			got, err := service.Calculate(context.Background(), tt.total)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockCache.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
