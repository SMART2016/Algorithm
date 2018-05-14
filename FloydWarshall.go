package main

import (
	"fmt"
	"math"
)

func AllPairShortestPath() {
	graph := getAdjacencyMatrix()
	for k := 0; k < len(graph); k++ {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				graph[i][j] = math.Min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	fmt.Println("All Pair shortest Path:", graph)
}

func getAdjacencyMatrix() [4][4]float64 {
	a := [4][4]float64{
		{0, 4, infinity, 2},
		{3, 0, 5, infinity},
		{infinity, infinity, 0, infinity},
		{infinity, infinity, 3, 0},
	}
	return a
}
