package main

import "fmt"

// This package will implement the heap data structure from the scratch
// Reference: https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc#ba57
type MinHeap []int

func main() {
	arr := []int{9, 4, 3, 8, 2, 1, 6}
	minHeap := &MinHeap{}
	//*minHeap = arr
	minHeap.BuildHeap(arr)
	fmt.Println("build min heap: ", *minHeap)

	minHeap.insert(0)

	fmt.Println("min heap after insert value 2: ", *minHeap)

	fmt.Println("Deleted item from the heap", minHeap.remove())

	fmt.Println("After deleting an item: ", *minHeap)
}

func (h *MinHeap) BuildHeap(arr []int) {
	lastNonLeafNodeIdx := (len(arr) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		fmt.Println("I am here", currentIdx)
		endIdx := len(arr) - 1
		h.shiftDown(currentIdx, endIdx)
	}
}
func (h *MinHeap) remove() int {
	n := len(*h)
	//swap the first element and last element of heap array
	h.swap(0, n-1)
	valueToRemove := (*h)[n-1]

	//pop the last element of the array
	*h = (*h)[:n-1]

	//call shifdown to update order of the heap
	h.shiftDown(0, n-2)
	return valueToRemove
}

func (h *MinHeap) insert(val int) {
	*h = append(*h, val)
	h.shiftUp()
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its parent node until it is in the correct position
func (h *MinHeap) shiftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && (*h)[currentIdx] < (*h)[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its smaller child node until it is in the correct position
func (h *MinHeap) shiftDown(currentIdx int, endIdx int) {
	leftChildIndex := currentIdx*2 + 1
	for leftChildIndex <= endIdx {
		rightChildIndex := currentIdx*2 + 2
		if rightChildIndex > endIdx {
			rightChildIndex = -1
		}
		// get the smaller child node to swap
		idxToSwap := leftChildIndex
		if rightChildIndex != -1 && (*h)[rightChildIndex] < (*h)[leftChildIndex] {
			idxToSwap = rightChildIndex
		}

		// check if value of swap node is less than the value at currentIdx
		if (*h)[idxToSwap] < (*h)[currentIdx] {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftChildIndex = currentIdx*2 + 1
		} else {
			return
		}

	}
}
