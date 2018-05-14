package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {

	//Rabin-Karp Algo call
	fmt.Println("------ RABIN KARP ---------\n")
	pattern := "abc" //hash(abc) = 97 + 98 * 101 + 99 * 101 * 101
	text := "deabcfabc"
	fmt.Println("Indexes matched for pattern in input text: ", compare(pattern, text))

	//CountPairWithGivenSum.go call
	fmt.Println("------ PAIR with Given Sum---------\n")
	a := []int{1, 3, 5, 7, 8, 6, 0, -1}
	sum := 6
	fmt.Println("Num of Pairs for sum=6 is :", count(a, sum))

	//Reduce Array Problem
	fmt.Println("------ REDUCE ARRAYR ---------\n")
	arr := []int{10, 40, 20}
	fmt.Println("Reduced Array for {10, 40, 20} is", reduceArray(arr))

	//OptimalJobSchedule
	fmt.Println("------ MAX JOB Schedule ---------\n")
	s := []interval{
		{2, 8},
		{3, 5},
		{2, 4},
		{10, 13},
		{14, 17},
		{9, 11},
	}
	fmt.Println("Optimal set of Jobs: ", getOptimalJobSet(s))

	//LevelOrderTraversal
	fmt.Println("------ TREE LEVEL-ORDER ---------\n")
	t := getTree()
	BFT(t)

	//Inorder Traversal
	fmt.Println("------ TREE IN-ORDER ---------\n")
	InorderDFT(t)

	//PreOrder traversal
	fmt.Println("------ TREE PRE-ORDER ---------\n")
	PreorderDFT(t)

	//Heap Sort MAX/Min heap : selection sort using heap datastructure
	fmt.Println("------ HEAP sort ---------\n")
	arr = []int{4, 10, 3, 5, 1}
	HeapSort(arr, 5)

	//Partition
	fmt.Println("------ QUICK sort ---------\n")
	arr = []int{1, 2, 8, 5, 1, 2, 9}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Quick Sort: ", arr)

	//Counting sort
	fmt.Println("------ Counting sort ---------\n")
	countSort(arr, 10)

	//Is Graph Tree
	fmt.Println("------------IS GRAPH TREE-------------\n")
	vertices := 3
	degree := []int{1, 2, 1}
	g := creatGraph()
	g.isTree(vertices, degree)

	//Breath First Traversal Graph
	fmt.Println("------GRAPH BFT ---------\n")
	g = creatGraph()
	g.BFT(1)
	g.PrintParentLst("BFT")
	fmt.Println("BFT shortest path", g.shortestPathFrmRoot(5, make([]int, 0)))

	//Depth First Traversal Graph
	fmt.Println("------GRAPH DFT ---------\n")
	g = creatGraph()
	g.DFT(1, 0)
	g.PrintParentLst("DFT")
	fmt.Println(g.parent[1])
	fmt.Println("DFT cycles", g.getCycles())
	fmt.Println("DFT Bridges:", g.getBridges())
	fmt.Println("DFT Articulation Points:", g.getArticulationPoints())

	//Weighted Graph
	fmt.Println("------Weighted Graph Prims Algo---------\n")
	wg := createWeightedUndirectedGraph()
	wg.MinSpanningPrims(1)

	fmt.Println("------Weighted Graph Djkstras Algo---------\n")
	wg = createWeightedUndirectedGraph()
	wg.ShortestPath(1)
}

func createWeightedUndirectedGraph() *WeightedGraph {
	wg := NewWeightedGraph()
	wg.addEdge(1, 2, 2)
	wg.addEdge(2, 1, 2)
	wg.addEdge(2, 3, 5)
	wg.addEdge(3, 2, 5)
	wg.addEdge(3, 1, 10)
	wg.addEdge(1, 3, 10)
	wg.addEdge(1, 4, 11)
	wg.addEdge(4, 1, 11)
	wg.addEdge(4, 3, 6)
	wg.addEdge(3, 4, 6)
	return wg
}
func creatGraph() *Graph {
	g := NewGraph()
	g.addEdge(1, 2)
	g.addEdge(1, 6)
	g.addEdge(2, 3)
	g.addEdge(2, 1)
	g.addEdge(2, 5)
	g.addEdge(3, 2)
	g.addEdge(3, 4)
	g.addEdge(4, 3)
	g.addEdge(4, 5)
	g.addEdge(5, 6)
	g.addEdge(5, 4)
	g.addEdge(5, 2)
	g.addEdge(6, 5)
	g.addEdge(6, 1)
	return g
}
func getTree() *tree.Tree {
	t := new(tree.Tree)
	t.Left = new(tree.Tree)
	t.Right = new(tree.Tree)
	t.Value = 2
	t.Left.Value = 7
	t.Right.Value = 5
	t.Left.Right = new(tree.Tree)
	t.Left.Left = new(tree.Tree)
	t.Left.Right.Value = 11
	t.Left.Left.Value = 13
	return t
}

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}
