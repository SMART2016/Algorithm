package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type traverseAlgorithm func(node *tree.Tree, queue chan *tree.Tree)

func BFT(t *tree.Tree) {
	fmt.Println("Tree length is: ", treeLength(t, 0))
	queue := make(chan *tree.Tree, 100)
	waitqueue := make(chan int, 100)

	if t != nil {
		queue <- t
	}

	go func() {
		traverseTree(t, queue, waitqueue, levelOrderAlgo)
	}()

	for n := 1; n < treeLength(t, 0); n++ {
		<-waitqueue
	}
}

func traverseTree(t *tree.Tree, queue chan *tree.Tree, waitqueue chan int, algo traverseAlgorithm) {
	len := treeLength(t, 0)
	for i := 1; i < len; i++ {
		select {
		case node := <-queue:
			algo(node, queue)
		default:
			fmt.Print("No data")
		}
		waitqueue <- i
	}
}

func levelOrderAlgo(node *tree.Tree, queue chan *tree.Tree) {
	fmt.Println("val is:", node.Value)

	if node.Left != nil {
		queue <- node.Left
	}

	if node.Right != nil {
		queue <- node.Right
	}
}
func treeLength(t *tree.Tree, oldLength int) int {

	if t != nil {
		if t.Left == nil && t.Right == nil {
			oldLength = oldLength + 1
		} else {
			oldLength = treeLength(t.Left, oldLength) + 1
			oldLength = treeLength(t.Right, oldLength) + 1
		}
	}
	return oldLength
}
