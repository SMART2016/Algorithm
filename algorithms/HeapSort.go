package main

import "fmt"

//N*log(N) time
func HeapSort(arr []int, n int) {
	//Create a Max or Min heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i, n, "max")
	}
	fmt.Println(arr)
	//Swap the maximum with the last if its a max heap
	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, 0, i-1, "max")
	}

	fmt.Println("HeapSorted Set:", arr)
}

func heapify(a []int, index int, size int, algo string) {
	switch algo {
	case "min":
		minHeapify(a, index, size)
		break
	case "max":
		maxHeapify(a, index, size)
		break
	}
}

//Log(N) time
func maxHeapify(a []int, index int, size int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left <= size && a[index] < a[left] {
		largest = left
	}
	if right <= size && a[largest] < a[right] {
		largest = right
	}

	if largest != index {
		swap(a, largest, index)
		maxHeapify(a, largest, size)
	}
}

func minHeapify(a []int, index int, size int) {
	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left <= size && a[index] > a[left] {
		smallest = left
	}
	if right <= size && a[smallest] > a[right] {
		smallest = right
	}

	if smallest != index {
		swap(a, smallest, index)
		minHeapify(a, smallest, size)
	}
}

func swap(arr []int, index1 int, index2 int) {
	tmp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = tmp
}
