package main

func subsetSum(items []int, target int) int {
	dp := make([][]int, len(items)+1)
	totalSum := 0
	for _, num := range items {
		totalSum = totalSum + num
	}

	dp[0] = make([]int, target+1)
	dp[0][0] = 1

	for itemIndex := 1; itemIndex <= len(items); itemIndex++ {
		dp[itemIndex] = make([]int, target+1)
		for currentTargetSum := 0; currentTargetSum <= target; currentTargetSum++ {
			if items[itemIndex-1] <= currentTargetSum && currentTargetSum-items[itemIndex-1] >= 0 {
				dp[itemIndex][currentTargetSum] = dp[itemIndex-1][currentTargetSum] + dp[itemIndex-1][currentTargetSum-items[itemIndex-1]]
			} else {
				dp[itemIndex][currentTargetSum] = dp[itemIndex-1][currentTargetSum]
			}

		}
	}
	return dp[len(items)][target]
}

func TestSubsetSum() int {
	items := []int{1, 2, 3, 7}
	target := 14
	return subsetSum(items[:], target)
}
