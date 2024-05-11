package main

import "fmt"

func main() {
	fmt.Println("hello DP .... :-)")
	fmt.Println("0/1 Knapsack result", TestKnapsack01())
	fmt.Println("Target Sum result", TestTargetSum())
	fmt.Println("Subset Sum result", TestSubsetSum() > 0)
	fmt.Println("Target Subset Sum result", TestSubsetSum())
}
