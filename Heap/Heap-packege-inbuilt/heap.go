package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func main() {
	arr := []int{23, 56, 43, 4, 5, 6, 10, 95, 8}
	h := IntHeap(arr)
	heap.Init(&h)
	heap.Push(&h, 96) // push an element in a heap
	fmt.Println("Heap: ", &h)

	for i := 0; i < len(arr); i++ {
		fmt.Println("Min Heap Elements: ", heap.Pop(&h))
	}
}

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	oldheap := *h
	n := len(oldheap)
	valToPop := oldheap[n-1]
	*h = oldheap[:len(*h)-1]
	return valToPop
}
