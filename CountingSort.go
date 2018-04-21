package main

import "fmt"

func countSort(a []int, max int) {
	b := make([]int, max)
	c := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		b[a[i]] = b[a[i]] + 1
	}

	for i := 1; i < len(b); i++ {
		b[i] = b[i-1] + b[i]
	}

	for i := 0; i < len(a); i++ {
		c[b[a[i]]-1] = a[i]
		b[a[i]] = b[a[i]] - 1
	}
	fmt.Println(c)
}
