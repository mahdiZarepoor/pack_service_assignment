package packs

import (
	"reflect"
	"testing"
)

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
			result := Calculate(tt.packSizes, tt.order)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Calculate(%d) = %v, want %v", tt.order, result, tt.expected)
			}
		})
	}
}

// func TestFindMinimalTotal(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		order     int
// 		packSizes []int
// 		expected  int
// 	}{
// 		{
// 			name:      "exact pack size",
// 			order:     2000,
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			expected:  2000,
// 		},
// 		{
// 			name:      "smaller than smallest pack",
// 			order:     100,
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			expected:  250,
// 		},
// 		{
// 			name:      "between pack sizes",
// 			order:     1500,
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			expected:  2000,
// 		},
// 		{
// 			name:      "large order",
// 			order:     7000,
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			expected:  7000,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := findMinimalTotal(tt.packSizes, tt.order)
// 			if result != tt.expected {
// 				t.Errorf("findMinimalTotal(%d) = %d, want %d", tt.order, result, tt.expected)
// 			}
// 		})
// 	}
// }

// func TestCanReachTotal(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		packSizes []int
// 		total     int
// 		expected  bool
// 	}{
// 		{
// 			name:      "exact pack size",
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			total:     2000,
// 			expected:  true,
// 		},
// 		{
// 			name:      "sum of multiple packs",
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			total:     3000,
// 			expected:  true,
// 		},
// 		{
// 			name:      "impossible total",
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			total:     150,
// 			expected:  false,
// 		},
// 		{
// 			name:      "zero total",
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			total:     0,
// 			expected:  true,
// 		},
// 		{
// 			name:      "large possible total",
// 			packSizes: []int{2000, 250, 1000, 5000, 500},
// 			total:     7500,
// 			expected:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := canReachTotal(tt.packSizes, tt.total)
// 			if result != tt.expected {
// 				t.Errorf("canReachTotal(%d) = %v, want %v", tt.total, result, tt.expected)
// 			}
// 		})
// 	}
// }
