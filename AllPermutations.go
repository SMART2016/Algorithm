package main

import "fmt"

func finAllPerm(n int) {
	a := make([]int, n)
	backTrackAllPerm(a, -1, n)
}

func backTrackAllPerm(a []int, k int, n int) {
	c := make([]int, n)

	if is_solutionAllPerm(a, k, n) {
		process_solutionAllPerm(a, k, n)
	} else {
		k = k + 1
		if k < n {
			nCandidates := constructCandidate(a, k, n, c)
			//issue with this loop
			for i := 0; i < nCandidates; i++ {

				a[k] = c[i]
				backTrackAllPerm(a, k, n)

			}
		}
	}
}

func constructCandidate(a []int, k int, n int, c []int) int {
	var no int = 0
	for x := 1; x <= n; x++ {
		if !contains(a, k, x) {
			c[no] = x
			no = no + 1
		}
	}
	return no
}

func contains(a []int, k int, num int) bool {
	for i := 0; i < k; i++ {
		if a[i] == num {
			return true
		}
	}
	return false
}
func is_solutionAllPerm(a []int, k int, n int) bool {
	return k == n-1
}

func process_solutionAllPerm(a []int, k int, n int) {
	fmt.Print(a, "\n")
}
