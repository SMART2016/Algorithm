package main

import (
	"fmt"

	queue "github.com/jupp0r/go-priority-queue"
)

func (wg *WeightedGraph) ShortestPath(root int) {
	mstSet := make([]int, len(wg.vertices))
	if _, ok := wg.vertices[root]; ok {
		parents := make(map[int]int)
		mstMap := make(map[int]float64)
		wg.createMstMap(mstMap, root)
		q := wg.createPriorityQueue(mstMap)
		wg.Djkstra(mstMap, mstSet, q, parents)
		fmt.Println("shortest Path:", mstSet)
		fmt.Println("shortest Path Parents:", parents)
	}
}

func (wg *WeightedGraph) Djkstra(mstMap map[int]float64, mstSet []int, q queue.PriorityQueue, parents map[int]int) {
	for i := 0; i <= q.Len(); i = i + 1 {
		node, _ := q.Pop()
		mstSet[i] = node.(int)
		nodeWeight := mstMap[node.(int)]
		delete(mstMap, node.(int))

		for k, v := range wg.vertices[node.(int)] {
			if !wg.contains(mstSet, k) {
				if mstMap[k] > (v.weight + nodeWeight) {
					parents[k] = node.(int)
					mstMap[k] = (v.weight + nodeWeight)
					q.UpdatePriority(k, (v.weight + nodeWeight))
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
