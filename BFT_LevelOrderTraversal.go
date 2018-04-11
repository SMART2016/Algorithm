package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func LevelOrder(t *tree.Tree) {
	fmt.Println("Tree length is: ", treeLength(t, 0))
	queue := make(chan *tree.Tree, 100)
	waitqueue := make(chan int, 100)
	if t != nil {
		queue <- t
	}

	go func() {
		len := treeLength(t, 0)
		for i := 1; i < len; i++ {
			select {
			case val := <-queue:

				fmt.Println("val is:", val.Value)

				if val.Left != nil {
					queue <- val.Left
				}

				if val.Right != nil {
					queue <- val.Right
				}
			default:
				fmt.Print("No data")
			}
			waitqueue <- i
		}

	}()

	for n := 1; n < treeLength(t, 0); n++ {
		<-waitqueue
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
