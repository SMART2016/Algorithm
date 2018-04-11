package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {

	//Rabin-Karp Algo call
	pattern := "abc" //hash(abc) = 97 + 98 * 101 + 99 * 101 * 101
	text := "deabcfabc"
	fmt.Println("Indexes matched for pattern in input text: ", compare(pattern, text))

	//CountPairWithGivenSum.go call
	a := []int{1, 3, 5, 7, 8, 6, 0, -1}
	sum := 6
	fmt.Println("Num of Pairs for sum=6 is :", count(a, sum))

	//Reduce Array Problem
	arr := []int{10, 40, 20}
	fmt.Println("Reduced Array for {10, 40, 20} is", reduceArray(arr))

	//OptimalJobSchedule
	s := []interval{
		{2, 8},
		{3, 5},
		{2, 4},
		{10, 13},
		{14, 17},
		{9, 11},
	}

	fmt.Println("Optimal set of Jobs: ", getOptimalJobSet(s))

	//LevelOrderTraversal
	t := new(tree.Tree)
	t.Left = new(tree.Tree)
	t.Right = new(tree.Tree)
	t.Value = 2
	t.Left.Value = 7
	t.Right.Value = 5
	//	t.Right.Right = tree.New(9)
	//	t.Right.Right.Right = tree.New(4)
	//	t.Left.Right = tree.New(6)
	//	t.Left.Right.Left = tree.New(1)
	//	t.Left.Right.Right = tree.New(11)
	LevelOrder(t)
}
