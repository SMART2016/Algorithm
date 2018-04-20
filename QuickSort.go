package main

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition(a, lo, hi)
	if lo < p-1 {
		quickSort(a, lo, p-1)
	}
	quickSort(a, p+1, hi)
}
func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}
