package main

import "fmt"

func backTrack(a []int, k int, n int) {
	c := make([]int, n)
	var nCandidates int
	if is_solution(a, k, n) {
		process_solution(a, k, n)
	} else {
		k = k + 1
		if k < n {
			nCandidates = constrcuct_candidate(a, k, n, c, nCandidates)
			for i := 0; i < nCandidates; i++ {

				a[k] = c[i]
				backTrack(a, k, n)
			}
		}
	}
}

func constrcuct_candidate(a []int, k int, n int, c []int, nCandidates int) int {
	c[0] = 1
	c[1] = 0
	nCandidates = 2
	return nCandidates
}

func is_solution(a []int, k int, n int) bool {
	return k == n-1
}

func process_solution(a []int, k int, n int) {
	fmt.Print("{")
	for i := 1; i <= k; i++ {
		if a[i] == 1 {
			fmt.Print(i)
		}
	}
	fmt.Print("}\n")
}

func generate_subset(n int) {
	a := make([]int, n)
	backTrack(a, -1, n)
}
