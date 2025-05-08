package packs

import "sort"

// var packSizes = []int{2000, 250, 1000, 5000, 500}

func calculatePacks(packSizes []int, order int) map[int]int {
	// Ensure pack sizes are sorted in descending order for greedy approach
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// Step 1: Find the minimal total >= order that can be formed by pack sizes
	minTotal := findMinimalTotal(packSizes, order)
	if minTotal == -1 {
		return nil // no solution (shouldn't happen with pack size 1, but our smallest is 250)
	}

	// Step 2: Find the minimal packs to sum to minTotal
	packCount := make(map[int]int)
	remaining := minTotal
	for _, pack := range packSizes {
		if remaining <= 0 {
			break
		}
		count := remaining / pack
		if count > 0 {
			packCount[pack] = count
			remaining -= count * pack
		}
	}

	return packCount
}

func findMinimalTotal(packSizes []int, order int) int {
	// We'll check from order upwards until we find a feasible total
	maxPack := packSizes[len(packSizes)-1]
	for total := order; total <= order+maxPack; total++ {
		if canReachTotal(packSizes, total) {
			return total
		}
	}
	return -1
}

func canReachTotal(packSizes []int, total int) bool {
	// Check if total can be formed by any combination of pack sizes
	// Using a greedy approach may not work, so use DP
	dp := make([]bool, total+1)
	dp[0] = true
	for i := 1; i <= total; i++ {
		for _, pack := range packSizes {
			if i >= pack && dp[i-pack] {
				dp[i] = true
				break
			}
		}
	}
	return dp[total]
}
