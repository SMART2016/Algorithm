package main

import (
	"fmt"
)

var cycles []string

func (l *Graph) DFT(node int, id int) int {
	child := 0
	//fmt.Println("Node is:", node)
	if val, ok := l.lst[node]; ok {
		l.id[node] = id
		l.low[id] = id

		l.frontier[node] = true
		curr := val.Front()
		for i := 0; i < val.Length(); i++ {
			//fmt.Printf("Adjacent node for %v is %v\n", node, curr.Value)
			_, ok := l.frontier[curr.Value.(int)]
			_, okExplored := l.explored[curr.Value.(int)]

			if !ok && !okExplored {
				l.parent[curr.Value.(int)] = node
				child = child + 1
				id = l.DFT(curr.Value.(int), id+1)
				l.low[l.id[node]] = l.min(l.low[l.id[node]], l.low[l.id[curr.Value.(int)]])
				val, _ := l.parent[node]
				if val == 0 && child > 1 {
					fmt.Printf("AP is: %v with number of children; %v\n", node, child)
					l.AP = append(l.AP, node)
				}

				if val != 0 && l.id[node] <= l.low[l.id[curr.Value.(int)]] {
					fmt.Printf("AP is: %v with number of children; %v\n", node, child)
					l.AP = append(l.AP, node)
				}
				l.isBridge(node, curr.Value.(int))
			} else if _, ok := l.explored[curr.Value.(int)]; !ok {
				if l.parent[node] == curr.Value.(int) {
					curr = curr.Next()
					continue
				}
				l.process_DFTedge(node, curr.Value.(int))
			}
			curr = curr.Next()
		}
		l.process_dftVertex(node)
		delete(l.frontier, node)
		l.explored[node] = true
	}
	return id
}

//Called ones for each edge, and if the adjacent node to the current
//node is already explored that means the edge was already visited earlier
//- Also can be used to detect cycles in a graph
//- can be used to detect all Back Edges
func (l *Graph) process_DFTedge(x int, y int) {
	l.low[l.id[x]] = l.min(l.low[l.id[x]], l.id[y])
	if l.parent[x] != y {
		l.cycles = append(l.cycles, fmt.Sprintf("%d --> %d", x, y))
	}
}

func (l *Graph) getArticulationPoints() []int {
	return l.AP
}
func (l *Graph) getBridges() []string {
	return l.bridges
}

func (l *Graph) min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
func (l *Graph) isBridge(x int, y int) {
	if l.id[x] < l.low[l.id[y]] {
		l.bridges = append(l.bridges, fmt.Sprintf("%d --> %d", x, y))
	}
}
func (l *Graph) getCycles() []string {
	return l.cycles
}
func (l *Graph) process_dftVertex(x int) {
	//fmt.Println("Proces vertex: ", x)
}
