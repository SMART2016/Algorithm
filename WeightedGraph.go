package main

import "fmt"

type WeightedGraph struct {
	vertices map[int]map[int]*edge
}

type edge struct {
	x      int
	y      int
	weight float64
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		vertices: make(map[int]map[int]*edge),
	}
}

func (wg *WeightedGraph) addEdge(x, y int, weight float64) {
	edgeMap := wg.vertices[x]
	side := &edge{x: x,
		y:      y,
		weight: weight,
	}
	if edgeMap != nil {
		edgeMap[y] = side
	} else {
		edgeMap = make(map[int]*edge)
		wg.vertices[x] = edgeMap
		edgeMap[y] = side

	}

}

func (wg *WeightedGraph) String() {
	for k, v := range wg.vertices {

		for k1, v1 := range v {
			fmt.Println(k, "-->", k1, " Weight:", v1.weight)
		}
	}
}
