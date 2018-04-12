package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"

	"golang.org/x/tour/tree"
)

func InorderDFT(t *tree.Tree) {
	store := stack.New()
	current := t
	store = populateStackWithLeftChild(store, t)
	for store.Len() > 0 {
		current = store.Pop().(*tree.Tree)
		fmt.Println("Element is:", current.Value)
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
}

func populateStackWithLeftChild(store *stack.Stack, treeNode *tree.Tree) *stack.Stack {
	current := treeNode
	for current != nil {
		store.Push(current)
		current = current.Left
	}
	return store
}
