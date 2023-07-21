package main

import (
	"Algorithm/src/greedy"
	"encoding/json"
	"fmt"
	_ "fmt"
)

func main() {

	//Rabin-Karp Algo call
	// fmt.Println("------ RABIN KARP ---------\n")
	// pattern := "abc" //hash(abc) = 97 + 98 * 101 + 99 * 101 * 101
	// text := "deabcfabc"
	// fmt.Println("Indexes matched for pattern in input text: ", compare(pattern, text))

	// //CountPairWithGivenSum.go call
	// fmt.Println("------ PAIR with Given Sum---------\n")
	// a := []int{1, 3, 5, 7, 8, 6, 0, -1}
	// sum := 6
	// fmt.Println("Num of Pairs for sum=6 is :", count(a, sum))

	// //Reduce Array Problem
	// fmt.Println("------ REDUCE ARRAYR ---------\n")
	// arr := []int{10, 40, 20}
	// fmt.Println("Reduced Array for {10, 40, 20} is", reduceArray(arr))

	// //OptimalJobSchedule
	// fmt.Println("------ MAX JOB Schedule ---------\n")
	// s := []interval{
	// 	{2, 8},
	// 	{3, 5},
	// 	{2, 4},
	// 	{10, 13},
	// 	{14, 17},
	// 	{9, 11},
	// }
	// fmt.Println("Optimal set of Jobs: ", getOptimalJobSet(s))

	// //LevelOrderTraversal
	// fmt.Println("------ TREE LEVEL-ORDER ---------\n")
	// t := getTree()
	// BFT(t)

	// //Inorder Traversal
	// fmt.Println("------ TREE IN-ORDER ---------\n")
	// InorderDFT(t)

	// //PreOrder traversal
	// fmt.Println("------ TREE PRE-ORDER ---------\n")
	// PreorderDFT(t)

	// //Heap Sort MAX/Min heap : selection sort using heap datastructure
	// fmt.Println("------ HEAP sort ---------\n")
	// arr = []int{4, 10, 3, 5, 1}
	// HeapSort(arr, 5)

	// //Partition
	// fmt.Println("------ QUICK sort ---------\n")
	// arr = []int{1, 2, 8, 5, 1, 2, 9}
	// quickSort(arr, 0, len(arr)-1)
	// fmt.Println("Quick Sort: ", arr)

	// //Counting sort
	// fmt.Println("------ Counting sort ---------\n")
	// countSort(arr, 10)

	// //Is Graph Tree
	// fmt.Println("------------IS GRAPH TREE-------------\n")
	// vertices := 3
	// degree := []int{1, 2, 1}
	// g := creatGraph()
	// g.isTree(vertices, degree)

	// //Breath First Traversal Graph
	// fmt.Println("------GRAPH BFT ---------\n")
	// g = creatGraph()
	// g.BFT(1)
	// g.PrintParentLst("BFT")
	// fmt.Println("BFT shortest path", g.shortestPathFrmRoot(5, make([]int, 0)))

	// //Depth First Traversal Graph
	// fmt.Println("------GRAPH DFT ---------\n")
	// g = creatGraph()
	// g.DFT(1, 0)
	// g.PrintParentLst("DFT")
	// fmt.Println(g.parent[1])
	// fmt.Println("DFT cycles", g.getCycles())
	// fmt.Println("DFT Bridges:", g.getBridges())
	// fmt.Println("DFT Articulation Points:", g.getArticulationPoints())

	// //Weighted Graph
	// fmt.Println("------Weighted Graph Prims Algo---------\n")
	// wg := createWeightedUndirectedGraph()
	// wg.MinSpanningPrims(1)

	// fmt.Println("------Weighted Graph Djkstras Algo---------\n")
	// wg = createWeightedUndirectedGraph()
	// wg.ShortestPath(1)

	// fmt.Println("------Floyd Warshall Algo---------\n")
	// AllPairShortestPath()

	// fmt.Println("------Backtracking : All Subsets of a set---------\n")
	// generate_subset(2)

	// fmt.Println("------Backtracking : All Permutation of n---------\n")
	// finAllPerm(4)

	// fmt.Println("------Longest Substring---------\n")
	// a1 := []int{1, 4, 2, 1, 2, 4}
	// fmt.Println("maxlength for: Arr ", a1, " is:", findSubsetwithEqlSumHalf(a1, 6))
	//concat(16, "abc", "xyz")

	runCandyDistributionProblem()
}

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}

func runCandyDistributionProblem() {
	marks := []int{357, 287, 248, 96, 22, 295, -78, -239, -482, 74, 269, 265, -269, -130, 351, 109, 490, 431, -171, 335, 321, -485, 313, -475, -43, 322, -132, 42, -171, 389, 493, -444, -265, -458, -71, -395, 134, -233, 211, 356, -330, -336, 274, -193, -421, -163, 29, 329, -466, -60, 96, 432, 171, -385, -52, 120, -403, -325, -97, 100, 61, -80, -82, 426, 263, -256, 476, -390, 444, -148, 414, 376, -298, 192, -183, -53, 386, 127, -329, 125, -328, 490, -12, 132, 40, 414, 347, 439, 255, -343, -84, 256, 38, -368, 307, 463, 29, -428, -25, -51, -385, 145, 86, 441, 361, 183, -407, -227, -404, -225, -279, -37, 280, -13, -258, -393, 164, 361, 146, -293, 248, -320, -389, -337, 341, -1, -445, -420, 414, -63, 328, 120, -462, 318, 500, -358, 233, -165, 274, -388, -393, -48, 312, -173, 281, 325, -167, 383, 353, 125, -416, -179, -285, -449, 418, 288, 62, -186, 413, -500, 199, 281, -163, -99, 193, -130, 414, 392, 299, 328, 156, -247, -85, -455, -274, 54, -161, 82, -265, 311, -129, -143, 45, 308, 408, 346, 438, -441, -237, 402, -428, -230, 24, 317, -189, -356, -53, 419, 426, -362, 399, 460, 335, 84, 177, 138, 229, 162, 416, -284, -44, -423, -9, -271, 425, -166, -482, -335, -357, 6, -191, 261, 391, -68, 224, 402, -487, -45, 312, 233, -393, 138, -95, -139, -239, 234, 227, -292, 117, -131, -221, 373, 97, -456, -292, 460, 238, 280, 43, 206, -8, -117, 274, -10, 182, -77, -421, -309, 493, -355, -88, 14, 67, 112, 177, 426, 468, 449, 263, -208, 198, 378, -431, 444, -383}
	minCandies := greedy.DistCandy(marks)
	fmt.Println("****** Min Candy Result:= ", minCandies)
}
