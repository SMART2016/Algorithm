package main

import "sort"

//Reduces the input array as below:
// arr = {10,40,20} --> {0,2,1}
//10 is the smalest so its assigned 0
//40 is the largest so assigned 2 and
//20 is the middle element so assigned 1
func reduceArray(arr []int) []int {
	var tmp []int
	m := make(map[int]int)
	for _, n := range arr {
		tmp = append(tmp, n)
		m[n] = 0
	}
	sort.Ints(arr)
	for i, num := range arr {
		m[num] = i
	}

	for i1, n1 := range tmp {
		arr[i1] = m[n1]
	}

	return arr
}
