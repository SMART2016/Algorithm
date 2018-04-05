package main

import "fmt"

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
