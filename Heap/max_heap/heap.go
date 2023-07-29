package main

import "fmt"

// This package will implement the max heap
// Reference: https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc#ba57

type MaxHeap []int

func main() {
	arr := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	maxHeap := &MaxHeap{}
	*maxHeap = arr
	maxHeap.BuildHeap(arr)

	fmt.Println("Max heap elements:", maxHeap)
	elem := maxHeap.remove()

	fmt.Println("Removed Element:", elem)

	fmt.Println("Max heap after remove the element:", maxHeap)

	maxHeap.insert(100)

	fmt.Println("max heap after inserting 100", maxHeap)

	for i := 0; i < len(arr); i++ {
		fmt.Println("Print all the elements in descending order", maxHeap.remove())
	}

}

func (h *MaxHeap) BuildHeap(arr []int) {
	lastNonLeafNodeIdx := (len(arr) - 2) / 2
	for i := lastNonLeafNodeIdx; i >= 0; i-- {
		endIdx := len(arr) - 1
		h.shiftDown(i, endIdx)
	}
}

func (h *MaxHeap) shiftDown(currentIdx, endIdx int) {
	leftChildIdx := currentIdx*2 + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx*2 + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		idxToSwap := leftChildIdx
		// find the largest among left and right child
		if rightChildIdx != -1 && (*h)[rightChildIdx] > (*h)[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check the largest among child and current root node
		if (*h)[idxToSwap] > (*h)[currentIdx] {
			// swap the currentParent & child elements
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftChildIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

func (h MaxHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) remove() int {
	n := len(*h)

	h.swap(0, n-1) // swap the first and last element

	valueToRemove := (*h)[n-1] // extract the removed val from the heap

	*h = (*h)[:n-1] //delete the last element from the heap

	h.shiftDown(0, n-2) // reorder the heap with shiftdown function

	return valueToRemove
}

// Time: O(logn) | Space: O(1)
// insert a new value to the end of the tree and update heap ordering
func (h *MaxHeap) insert(val int) {
	*h = append(*h, val) // add the element at n-1 position
	h.shiftUp()
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its parent node until it is in the correct position

func (h *MaxHeap) shiftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && (*h)[currentIdx] > (*h)[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}
