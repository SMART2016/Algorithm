package main

import "fmt"

func main() {
	fmt.Println("hello DP .... :-)")
	fmt.Println("0/1 Knapsack result", TestKnapsack01())
	fmt.Println("Target Sum result", TestTargetSum())
	fmt.Println("Subset Sum result", TestSubsetSum() > 0)
	fmt.Println("Target Subset Sum result", TestSubsetSum())
	fmt.Println("Minimum Partition Array Sum Difference result", TestMinimumPartitionArraySumDifference())
	fmt.Println("Minimum refuel stops result", TestMinimumRefuelStops())
	fmt.Println("equal Sum subarray result", TestEqualSumSubarray() > 0)
	fmt.Println("# of square sub matrix", TestCountSquareSubMatrix())
}
