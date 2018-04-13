package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"

	"golang.org/x/tour/tree"
)

func InorderDFT(t *tree.Tree) {
	ordersSet := make([]int, 0)
	store := stack.New()
	current := t
	store = populateStackWithLeftChild(store, t)
	for store.Len() > 0 {
		current = store.Pop().(*tree.Tree)
		ordersSet = append(ordersSet, current.Value)
		if current.Right != nil {
			current = current.Right
			if current.Left != nil {
				store = populateStackWithLeftChild(store, current)
			} else {
				store.Push(current)
			}
		}
		current = nil
	}

	fmt.Println("Inorder set is: ", ordersSet)
}
