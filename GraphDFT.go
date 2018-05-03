package main

import (
	"fmt"
)

var cycles []string

func (l *Graph) DFT(node int) {
	if val, ok := l.lst[node]; ok {

		l.frontier[node] = true
		curr := val.Front()
		for i := 0; i < val.Length(); i++ {
			if _, ok := l.frontier[curr.Value.(int)]; !ok {
				l.parent[curr.Value.(int)] = node
				l.DFT(curr.Value.(int))
			} else if _, ok := l.explored[curr.Value.(int)]; !ok {
				l.process_DFTedge(node, curr.Value.(int))
			}
			curr = curr.Next()
		}
		l.process_dftVertex(node)
		l.explored[node] = true
	}
}

//Called ones for each edge, and if the adjacent node to the current
//node is already explored that means the edge was already visited earlier
//Also can be used to detect cycles in a graph
func (l *Graph) process_DFTedge(x int, y int) {
	if l.parent[x] != y {
		l.cycles = append(l.cycles, fmt.Sprintf("%d --> %d", x, y))
	}
}

func (l *Graph) getCycles() []string {
	return l.cycles
}
func (l *Graph) process_dftVertex(x int) {
	//fmt.Println("Proces vertex: ", x)
}
