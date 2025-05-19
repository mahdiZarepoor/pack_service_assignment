package service

import (
	"context"
	"github.com/mahdiZarepoor/pack_service_assignment/cmd/app/configs"
	"reflect"
	"testing"

	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
		packSizes []int
		mockSetup func(*MockLogger)
		wantErr   bool
	}{
		{
			name:      "successful update",
			packSizes: []int{250, 500, 1000, 2000, 5000},
			mockSetup: func(ml *MockLogger) {
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := new(MockLogger)
			cfg := configs.Config{}

			service := NewPackService(cfg, mockLogger)
			err := service.Update(context.Background(), tt.packSizes)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			mockLogger.AssertExpectations(t)
		})
	}
}

func TestPackService_List(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*MockLogger)
		give      []int
		want      []int
		wantErr   bool
	}{
		{
			name:    "successful list",
			give:    []int{250, 500, 1000, 2000, 5000},
			want:    []int{250, 500, 1000, 2000, 5000},
			wantErr: false,
		},
		{
			name:    "empty list",
			give:    []int{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := new(MockLogger)
			cfg := configs.Config{}
			ctx := context.Background()

			service := NewPackService(cfg, mockLogger)
			service.Update(ctx, tt.give)
			got, err := service.List(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockLogger.AssertExpectations(t)
		})
	}
}

func TestPackService_Calculate(t *testing.T) {
	tests := []struct {
		name    string
		total   int
		want    map[int]int
		wantErr bool
	}{
		{
			name:    "successful calculation",
			total:   1200,
			want:    map[int]int{1000: 1, 250: 1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := new(MockLogger)
			cfg := configs.Config{}

			service := NewPackService(cfg, mockLogger)
			service.Update(context.Background(), []int{250, 500, 1000, 2000, 5000})
			got, err := service.Calculate(context.Background(), tt.total)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockLogger.AssertExpectations(t)
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		order     int
		packSizes []int
		expected  map[int]int
	}{
		{
			name:      "exact pack size match",
			order:     2000,
			packSizes: []int{2000, 250, 1000, 5000, 500},
			expected: map[int]int{
				2000: 1,
			},
		},
		{
			name:      "smaller than smallest pack",
			order:     100,
			packSizes: []int{2000, 250, 1000, 5000, 500},
			expected: map[int]int{
				250: 1,
			},
		},
		{
			name:      "multiple packs needed",
			order:     3000,
			packSizes: []int{2000, 250, 1000, 5000, 500},
			expected: map[int]int{
				2000: 1,
				1000: 1,
			},
		},
		{
			name:      "large order requiring multiple pack sizes",
			order:     7500,
			packSizes: []int{2000, 250, 1000, 5000, 500},
			expected: map[int]int{
				5000: 1,
				2000: 1,
				500:  1,
			},
		},
		{
			name:      "edge case",
			order:     500000,
			packSizes: []int{23, 31, 53},
			expected: map[int]int{
				23: 2,
				31: 7,
				53: 9429,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePacks(tt.packSizes, tt.order)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Calculate(%d) = %v, want %v", tt.order, result, tt.expected)
			}
		})
	}
}
