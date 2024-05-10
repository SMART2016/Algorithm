package main

import "fmt"

//Examples:
//Input  :  arr[] = {1, 5, 7, -1},
//          sum = 6
//Output :  2
//Pairs with sum 6 are (1, 5) and (7, -1)

//Input  :  arr[] = {1, 5, 7, -1, 5},
//          sum = 6
//Output :  3
func count(a []int, sum int) int {
	numPairs := 0
	if len(a) == 0 {
		fmt.Println("Array is empty")
	}

	m := make(map[int]int)

	for _, num := range a {
		m[num] = 1
	}

	for _, n := range a {
		if m[sum-n] == 1 {
			numPairs++
		}
	}
	return numPairs / 2
}
