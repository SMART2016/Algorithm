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
	t := getTree()
	BFT(t)

	//Inorder Traversal
	InorderDFT(t)

	//PreOrder traversal
	PreorderDFT(t)

	//Heap Sort MAX/Min heap : selection sort using heap datastructure
	arr = []int{4, 10, 3, 5, 1}
	HeapSort(arr, 5)

	//Partition
	arr = []int{1, 2, 8, 5, 1, 2, 9}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Quick Sort: ", arr)

	//Counting sort
	countSort(arr, 10)

}

func getTree() *tree.Tree {
	t := new(tree.Tree)
	t.Left = new(tree.Tree)
	t.Right = new(tree.Tree)
	t.Value = 2
	t.Left.Value = 7
	t.Right.Value = 5
	t.Left.Right = new(tree.Tree)
	t.Left.Left = new(tree.Tree)
	t.Left.Right.Value = 11
	t.Left.Left.Value = 13
	return t
}
