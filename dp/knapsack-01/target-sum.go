package main

func targetSum(items []int, target int) int {
	dp := make([][]int, len(items)+1)
	totalSum := 0
	for _, num := range items {
		totalSum = totalSum + num
	}

	// Transform to subset sum and below is the identified subset sum target
	targetSubsetSum := (totalSum + target) / 2
	dp[0] = make([]int, targetSubsetSum+1)
	dp[0][0] = 1

	for itemIndex := 1; itemIndex <= len(items); itemIndex++ {
		dp[itemIndex] = make([]int, targetSubsetSum+1)
		for currentTargetSum := 0; currentTargetSum <= targetSubsetSum; currentTargetSum++ {
			if items[itemIndex-1] <= currentTargetSum && currentTargetSum-items[itemIndex-1] >= 0 {
				dp[itemIndex][currentTargetSum] = dp[itemIndex-1][currentTargetSum] + dp[itemIndex-1][currentTargetSum-items[itemIndex-1]]
			} else {
				dp[itemIndex][currentTargetSum] = dp[itemIndex-1][currentTargetSum]
			}

		}
	}
	return dp[len(items)][targetSubsetSum]
}

func TestTargetSum() int {
	items := []int{1, 1, 1, 1}
	target := 2
	return targetSum(items[:], target)
}
