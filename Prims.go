package main

import (
	"fmt"

	queue "github.com/jupp0r/go-priority-queue"
)

const (
	infinity = 9999999999
)

func (wg *WeightedGraph) MinSpanningPrims(root int) {
	mstSet := make([]int, len(wg.vertices))
	if _, ok := wg.vertices[root]; ok {
		mstMap := make(map[int]float64)
		wg.createMstMap(mstMap, root)
		q := wg.createPriorityQueue(mstMap)
		wg.prims(mstMap, mstSet, q)
		fmt.Println("Minimum spanning Tree:", mstSet)
	}
}

func (wg *WeightedGraph) prims(mstMap map[int]float64, mstSet []int, q queue.PriorityQueue) {
	for i := 0; i <= q.Len(); i = i + 1 {
		node, _ := q.Pop()
		mstSet[i] = node.(int)
		delete(mstMap, node.(int))

		for k, v := range wg.vertices[node.(int)] {
			if !wg.contains(mstSet, k) {

				if mstMap[k] > v.weight {
					mstMap[k] = v.weight
					q.UpdatePriority(k, v.weight)
				}
			}
		}
	}

	if q.Len() != 0 {
		n, _ := q.Pop()
		delete(mstMap, n.(int))
		mstSet[len(wg.vertices)-1] = n.(int)

	}
}
func (wg *WeightedGraph) createPriorityQueue(mstMap map[int]float64) queue.PriorityQueue {
	q := queue.New()
	for node, weight := range mstMap {
		q.Insert(node, weight)
	}
	return q
}

func (wg *WeightedGraph) createMstMap(mstMap map[int]float64, root int) {
	for node, _ := range wg.vertices {
		if node == root {
			mstMap[node] = 0
		} else {
			mstMap[node] = infinity
		}
	}
}

func (wg *WeightedGraph) contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//func (wg *WeightedGraph) printMstMap(mstMap map[int]float64) {
//	for k, v := range mstMap {
//		fmt.Println(k, "-->", v.priority)
//	}
//}
