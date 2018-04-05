package main

import "fmt"

func main() {

	//Rabin-Karp Algo call
	pattern := "abc" //hash(abc) = 97 + 98 * 101 + 99 * 101 * 101
	text := "deabcfabc"
	compare(pattern, text)

	//CountPairWithGivenSum.go call
	a := []int{1, 3, 5, 7, 8, 6, 0, -1}
	sum := 6
	fmt.Println("Num of Pairs:", count(a, sum))
}
