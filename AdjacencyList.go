package main

import (
	"fmt"

	linklist "github.com/dustinspecker/go-singly-linked-list"
)

type Graph struct {
	lst      map[int]*linklist.SinglyLinkedList
	explored map[int]bool
	frontier map[int]bool
	parent   map[int]int
	cycles   []string
}

func NewGraph() *Graph {
	return &Graph{
		lst:      make(map[int]*linklist.SinglyLinkedList),
		explored: make(map[int]bool),
		frontier: make(map[int]bool),
		parent:   make(map[int]int),
		cycles:   make([]string, 0),
	}
}

func (l *Graph) addEdge(key int, edgeNode int) {
	node := &linklist.Node{Value: edgeNode}
	if val, ok := l.lst[key]; ok {
		val.Append(node)
	} else {
		edgeLst := linklist.New()
		edgeLst.Append(node)
		l.lst[key] = edgeLst

	}
}

func (l *Graph) string() {
	for k, v := range l.lst {
		fmt.Print(k, "-->")
		for curr := v.Front(); curr != nil; curr = curr.Next() {
			fmt.Print(curr.Value, ",")
		}
		fmt.Print("\n")
	}
}

func (l *Graph) getEdges(key int) *linklist.SinglyLinkedList {
	return l.lst[key]
}