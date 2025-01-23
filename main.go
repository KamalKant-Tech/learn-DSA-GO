package main

import (
	"fmt"
)

func maxDistance(blocks []int) int {
	maxSize := -1
	maxSizeIndex := -1

	// Find the maximum block size and its index
	for i, size := range blocks {
		if size > maxSize {
			maxSize = size
			maxSizeIndex = i
		}
	}

	// Find the leftmost block with the maximum size
	leftDistance := maxSizeIndex
	for i := maxSizeIndex - 1; i >= 0; i-- {
		if blocks[i] < maxSize {
			break
		}
		leftDistance--
	}

	// Find the rightmost block with the maximum size
	rightDistance := len(blocks) - 1 - maxSizeIndex
	for i := maxSizeIndex + 1; i < len(blocks); i++ {
		if blocks[i] < maxSize {
			break
		}
		rightDistance--
	}

	// Return the maximum distance between the frogs
	return max(leftDistance, rightDistance)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	blocks := []int{2,6,8,5}
	fmt.Println("Maximum distance between frogs:", maxDistance(blocks)) // Output: 4
}


// Array of integers named "blocks" with the size N is given and every integer is the size of the block; 
// you have 2 frogs that start on one block (either the first one or the "optimal starting block") and the frogs want to get as 
// far away from eachother as possible. The frogs can only jump to the adjacent block if it is at least as big as the one they 
// are sitting on and can not jump if the adjacent block is smaller than the one they sit on. The exercise also states that 
// distance between blocks numbered J and K, where J<=K is computed as K - J + 1. The question is: what is the longest distance 
// that thez can possibly create between each other if thezy chose the optimal starting block.