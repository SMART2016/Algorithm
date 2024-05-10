package main

import "math"

func Knapsack01(totalItems int, weights []int, values []int, totalCapacity int) int {
	//result := 0
	dp := make([][]int, totalItems+1)

	for itemIndex := 0; itemIndex <= totalItems; itemIndex++ {
		column := make([]int, totalCapacity+1)
		dp[itemIndex] = column
		for currentCapacity := 0; currentCapacity <= totalCapacity; currentCapacity++ {
			if itemIndex == 0 || currentCapacity == 0 {
				dp[itemIndex][currentCapacity] = 0
			} else if weights[itemIndex-1] <= currentCapacity {
				dp[itemIndex][currentCapacity] = int(math.Max(float64(dp[itemIndex-1][weights[itemIndex-1]]), float64(values[itemIndex-1]+dp[itemIndex-1][currentCapacity-weights[itemIndex-1]])))
			} else {
				dp[itemIndex][currentCapacity] = dp[itemIndex-1][currentCapacity]
			}
		}
	}
	return dp[totalItems][totalCapacity]
}

func TestKnapsack() int {
	weights := []int{1, 2, 3, 5}
	values := []int{10, 5, 4, 8}
	totalCapacity := 5
	return Knapsack01(4, weights[:], values[:], totalCapacity)
}
