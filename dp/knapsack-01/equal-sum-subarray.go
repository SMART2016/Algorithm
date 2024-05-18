package main

func equalSumSubarray(items []int) int {
	dp := make([][]int, len(items)+1)
	totalSum := 0

	//1. Find total Sum
	for _, num := range items {
		totalSum = totalSum + num
	}
	if totalSum%2 != 0 {
		return 0
	}

	for itemIndex := 0; itemIndex <= len(items); itemIndex++ {
		dp[itemIndex] = make([]int, totalSum/2+1)
		dp[itemIndex][0] = 1
	}

	for itemIndex := 1; itemIndex <= len(items); itemIndex++ {
		for currentSum := 1; currentSum <= totalSum/2; currentSum++ {
			if items[itemIndex-1] > currentSum {
				dp[itemIndex][currentSum] = dp[itemIndex-1][currentSum]
			} else {
				if dp[itemIndex-1][currentSum] == 1 {
					dp[itemIndex][currentSum] = 1
				} else {
					dp[itemIndex][currentSum] = dp[itemIndex-1][currentSum-items[itemIndex-1]]
				}

			}
		}
	}
	return dp[len(items)][totalSum/2]
}

func TestEqualSumSubarray() int {
	items := []int{2, 1, 5, 3, 1}
	return equalSumSubarray(items)
}
