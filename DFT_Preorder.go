package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"

	"golang.org/x/tour/tree"
)

func PreorderDFT(t *tree.Tree) {
	ordersSet := make([]int, 0)
	store := stack.New()
	current := t
	store.Push(current)
	for store.Len() > 0 {
		current = store.Pop().(*tree.Tree)
		ordersSet = append(ordersSet, current.Value)
		if current.Right != nil {
			store.Push(current.Right)
		}
		if current.Left != nil {
			store.Push(current.Left)
		}
	}

	fmt.Println("Preorder set is: ", ordersSet)
}

func populateStackWithLeftChild(store *stack.Stack, treeNode *tree.Tree) *stack.Stack {
	current := treeNode
	for current != nil {
		store.Push(current)
		current = current.Left
	}
	return store
}
