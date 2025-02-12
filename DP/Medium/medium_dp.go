package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxJump([]int{0, 2, 5, 6, 7}))
}

// Problem: 2498. Frog Jump II
// You are given a 0-indexed integer array stones sorted in strictly increasing order representing the positions of stones in a river.

// A frog, initially on the first stone, wants to travel to the last stone and then return to the first stone. However, it can jump to any stone at most once.

// The length of a jump is the absolute difference between the position of the stone the frog is currently on and the position of the stone to which the frog jumps.

// More formally, if the frog is at stones[i] and is jumping to stones[j], the length of the jump is |stones[i] - stones[j]|.
// The cost of a path is the maximum length of a jump among all jumps in the path.

// Return the minimum cost of a path for the frog.
func maxJumpOptimal(stones []int) int {
	n := len(stones)
	if n == 2 {
		return stones[1] - stones[0]
	}

	maxJump := 0
	for i := 2; i < n; i++ {
		maxJump = int(math.Max(float64(maxJump), float64(stones[i]-stones[i-2])))
	}
	return maxJump
}

func maxJumpRecursive(stones []int, i int, memo map[int]int) int {
	if i == len(stones)-1 {
		return 0
	}
	if val, found := memo[i]; found {
		return val
	}

	minCost := math.MaxInt32
	for j := i + 1; j < len(stones); j++ {
		cost := stones[j] - stones[i]
		maxCost := int(math.Max(float64(cost), float64(maxJumpRecursive(stones, j, memo))))
		minCost = int(math.Min(float64(minCost), float64(maxCost)))
	}

	memo[i] = minCost
	return minCost
}

func maxJump(stones []int) int {
	memo := make(map[int]int)
	return maxJumpRecursive(stones, 0, memo)
}
