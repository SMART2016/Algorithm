package main

import "math"

func minimumPartitionArraySumDifference(items []int) int {
	dp := make([][]int, len(items)+1)
	totalSum := 0
	//1. Find total Sum
	for _, num := range items {
		totalSum = totalSum + num
	}

	//2. Initialize the first column with 1
	for itemIndex := 0; itemIndex <= len(items); itemIndex++ {
		dp[itemIndex] = make([]int, (totalSum/2)+1)
		dp[itemIndex][0] = 1
	}

	//3. Now fill the dp array by including or excluding the current element from the current subarray sum
	for itemIndex := 1; itemIndex <= len(items); itemIndex++ {
		//dp[itemIndex] = make([]int, (totalSum/2)+1)
		for currentTargetSum := 0; currentTargetSum <= totalSum/2; currentTargetSum++ {
			//Exclude the current element from the current subarray to see if the current sum can be achieved
			if dp[itemIndex-1][currentTargetSum] == 1 {
				dp[itemIndex][currentTargetSum] = 1
				//Include the current element to the current subarray to see if the current sum can be achieved
			} else if items[itemIndex-1] <= currentTargetSum {
				dp[itemIndex][currentTargetSum] = dp[itemIndex-1][currentTargetSum-items[itemIndex-1]]
			}
		}
	}

	//4. Find the minimum difference
	S1 := 0
	for j := totalSum / 2; j >= 0; j-- {
		if dp[len(items)][j] == 1 {
			S1 = j
			break
		}
	}

	S2 := totalSum - S1

	return int(math.Abs(float64(S2) - float64(S1)))
}

func TestMinimumPartitionArraySumDifference() int {
	items := []int{3, 7, 4, 12, 2}
	return minimumPartitionArraySumDifference(items[:])
}
