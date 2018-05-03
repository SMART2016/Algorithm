package main

import "fmt"

func (l *Graph) BFT(node int) {
	var empty bool
	queue := make(chan int, 100)
	if _, ok := l.lst[node]; ok {
		l.frontier[node] = true
		queue <- node

		for {
			select {
			case n := <-queue:
				//Explore all child nodes and add them to the queue if they arent already explored, also process the newly formed edge
				l.process_node(queue, n)
				l.explored[n] = true
				delete(l.frontier, n)
				break
			default:
				empty = true
				break
			}
			if empty {
				break
			}
		}
	} else {
		fmt.Println("The node is not a part of the graph")
	}
}

func (l *Graph) process_node(queue chan int, node int) {
	edgeLst := l.lst[node]
	if edgeLst == nil {
		return
	}
	if _, ok := l.explored[node]; !ok {

		curr := edgeLst.Front()
		for i := 0; i < edgeLst.Length(); i++ {
			if _, ok := l.explored[curr.Value.(int)]; !ok {
				l.process_edge(node, curr.Value.(int))

				if _, ok := l.frontier[curr.Value.(int)]; !ok {
					l.frontier[curr.Value.(int)] = true
					l.parent[curr.Value.(int)] = node
					queue <- curr.Value.(int)

				}
			}
			curr = curr.Next()
		}
	}
	l.process_vertex(node)
}

func (l *Graph) PrintParentLst(traversalAlgo string) {
	fmt.Println(traversalAlgo, " Parent List: ", l.parent)
}

//Called ones for each edge, and if the adjacent node to the current
//node is already explored that means the edge was already visited earlier
func (l *Graph) process_edge(x int, y int) {
	fmt.Println("Proces Edge: ", x, "-->", y)
}

func (l *Graph) process_vertex(x int) {
	//fmt.Println("Proces vertex: ", x)
}
