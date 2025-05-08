package packs

import "sort"

// Calculate returns a map of pack sizes to their quantities needed to fulfill the order
func Calculate(packSizes []int, order int) map[int]int {
	if order <= 0 {
		return nil
	}

	// Sort pack sizes in ascending order
	sort.Ints(packSizes)

	// For orders smaller than the smallest pack size,
	// return the smallest pack size
	if order <= packSizes[0] {
		return map[int]int{
			packSizes[0]: 1,
		}
	}

	// Initialize dynamic programming arrays
	dp := make([]int, order+packSizes[len(packSizes)-1]+1)
	prev := make([]int, order+packSizes[len(packSizes)-1]+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	// Find all possible sums and track the combinations
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

	// Reconstruct the solution
	result := make(map[int]int)
	current := minValidSum
	for current > 0 {
		pack := prev[current]
		result[pack]++
		current -= pack
	}

	return result
}

func optimizePacks(packSizes []int, current map[int]int, total int) map[int]int {
	// Sort pack sizes in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// Try to replace larger packs with combinations of smaller packs
	for i := 0; i < len(packSizes)-1; i++ {
		largerPack := packSizes[i]
		if current[largerPack] == 0 {
			continue
		}

		// Try to replace with smaller packs
		smallerPacks := packSizes[i+1:]
		replacement := make(map[int]int)
		remaining := largerPack

		for _, pack := range smallerPacks {
			for remaining >= pack {
				replacement[pack]++
				remaining -= pack
			}
		}

		// If we can replace with smaller packs and use fewer total packs
		if remaining == 0 {
			newCount := make(map[int]int)
			for pack, count := range current {
				if pack != largerPack {
					newCount[pack] = count
				}
			}
			for pack, count := range replacement {
				newCount[pack] += count
			}

			// Check if we're using fewer packs
			totalPacks := 0
			for _, count := range newCount {
				totalPacks += count
			}
			oldTotalPacks := 0
			for _, count := range current {
				oldTotalPacks += count
			}

			if totalPacks < oldTotalPacks {
				return newCount
			}
		}
	}

	return nil
}

func findMinimalTotal(packSizes []int, order int) int {
	// Sort pack sizes in ascending order
	sortedSizes := make([]int, len(packSizes))
	copy(sortedSizes, packSizes)
	sort.Ints(sortedSizes)

	// Find the smallest pack size that is >= order
	for _, size := range sortedSizes {
		if size >= order {
			return size
		}
	}

	// If no single pack is large enough, find the smallest combination that works
	maxPack := sortedSizes[len(sortedSizes)-1]
	for total := order; total <= order+maxPack; total++ {
		if canReachTotal(packSizes, total) {
			return total
		}
	}
	return -1
}

func canReachTotal(packSizes []int, total int) bool {
	// Check if total can be formed by any combination of pack sizes
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
