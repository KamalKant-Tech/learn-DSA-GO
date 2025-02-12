package main

import (
	"fmt"
	"math"
)

func main() {
	// n := 16
	// dp := make([]int, n+1)
	// for i := range dp {
	// 	dp[i] = -1
	// }
	// //fmt.Println(dp, len(dp))
	// fmt.Println(FibonacciUsingMemoization(16, dp))
	// fmt.Println(FibonacciUsingTabulation(16))
	// fmt.Println(FibonacciNoSC(16))

	fmt.Println(FrogJumpWithNoSpace(6, []int{30, 10, 60, 10, 60, 50}))
	//fmt.Println(dp, len(dp))
}

// Using Recursion
// TC: O(n), SC: O(n) - stack space + O(n)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// using Memoization
// TC: O(n), SC: O(n) + O(n) - DP array
func FibonacciUsingMemoization(n int, dp []int) int {

	if n <= 1 {
		return n
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = FibonacciUsingMemoization(n-1, dp) + FibonacciUsingMemoization(n-2, dp)

	return dp[n]
}

// using Tabulation
// TC: O(n), SC: O(n) - DP array
func FibonacciUsingTabulation(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// using Variables removing the space complexity
// TC: O(n), SC: O(1)
func FibonacciNoSC(n int) int {
	prev := 0
	prev2 := 1

	for i := 2; i <= n+1; i++ {
		current := prev + prev2
		prev2 = prev
		prev = current
	}

	return prev
}

/**
* Problem: 70. Climbing Stairs
* You are climbing a staircase. It takes n steps to reach the top.
* Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
* This solution gives TLE for n = 45, lets convert this to DP
 */
func climbStairsResursion(n int) int {
	return climdStairsHelper(n)
}

func climdStairsHelper(index int) int {

	if index == 0 {
		return 1
	}

	if index == 1 {
		return 1
	}

	return climdStairsHelper(index-1) + climdStairsHelper(index-2)
}

// This solution using Memoization
func climbStairMemoization(n int) int {
	fmt.Println("Given Array: ", n)
	// declare a dp array
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	return climdStairsHelperMemoization(n, dp)
}

func climdStairsHelperMemoization(index int, dp []int) int {

	if index == 0 {
		return 1
	}

	if index == 1 {
		return 1
	}

	if dp[index] != -1 {
		return dp[index]
	}

	dp[index] = climdStairsHelperMemoization(index-1, dp) + climdStairsHelperMemoization(index-2, dp)
	fmt.Println(dp)
	return dp[index]
}

// This solution uses Tabulation approach
func climbStair(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

/**
* Problem statement
* There is a frog on the '1st' step of an 'N' stairs long staircase.
* The frog wants to reach the 'Nth' stair. 'HEIGHT[i]' is the height of the '(i+1)th' stair.
* If Frog jumps from 'ith' to 'jth' stair, the energy lost in the jump is given by absolute value of ( HEIGHT[i-1] - HEIGHT[j-1] ).
* If the Frog is on 'ith' staircase, he can jump either to '(i+1)th' stair or to '(i+2)th' stair.
* Your task is to find the minimum total energy used by the frog to reach from '1st' stair to 'Nth' stair.

* For Example
* If the given ‘HEIGHT’ array is [10,20,30,10], the answer 20 as the frog can jump from 1st stair to 2nd stair (|20-10| = 10 energy lost) and
* then a jump from 2nd stair to last stair (|10-20| = 10 energy lost). So, the total energy lost is 20.
 */

func FrogJump(n int, height []int) int {
	fmt.Println("Given Array: ", n, height)
	return FrogJumpRecursion(n-1, height)
}

func FrogJumpRecursion(n int, height []int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return int(math.Abs(float64(height[n] - height[n-1])))
	}

	left := FrogJumpRecursion(n-1, height) + int(math.Abs(float64(height[n]-height[n-1])))
	right := FrogJumpRecursion(n-2, height) + int(math.Abs(float64(height[n]-height[n-2])))
	return int(math.Min(float64(left), float64(right)))
}

// Frog Jump Solution Using Memoization
// TC : O(n), SC: O(n) - stack Space + O(n) - memo array
func FrogJumpMemoization(n int, height []int) int {
	fmt.Println("Given Array: ", n, height)
	memo := make([]int, len(height)+1)
	for i := range memo {
		memo[i] = -1
	}
	fmt.Println(memo)
	return FrogJumpRecursionMemoization(n-1, height, memo)
}

func FrogJumpRecursionMemoization(n int, height []int, memo []int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return int(math.Abs(float64(height[n] - height[n-1])))
	}

	if memo[n] != -1 {
		return memo[n]
	}

	left := FrogJumpRecursionMemoization(n-1, height, memo) + int(math.Abs(float64(height[n]-height[n-1])))
	right := FrogJumpRecursionMemoization(n-2, height, memo) + int(math.Abs(float64(height[n]-height[n-2])))
	memo[n] = int(math.Min(float64(left), float64(right)))
	return memo[n]
}

// Frog Jump Solution Using Tabulation
// TC : O(n), SC: O(n) - DP array
// We have eliminated SC for stack space since its not required now.
func FrogJumpTabulation(n int, height []int) int {
	fmt.Println("Given Array: ", n, height)
	dp := make(map[int]int, n)
	dp[0] = 0
	dp[1] = int(math.Abs(float64(height[1]) - float64(height[0])))
	for i := 2; i < n; i++ {
		left := dp[i-1] + int(math.Abs(float64(height[i]-height[i-1])))
		right := dp[i-2] + int(math.Abs(float64(height[i]-height[i-2])))
		dp[i] = int(math.Min(float64(left), float64(right)))
	}
	return dp[n-1]
}

// Frog Jump Solution Using No Variables which will not take space to store
// TC : O(n), SC: Constant
// We have eliminated SC for stack space since its not required now.
func FrogJumpWithNoSpace(n int, height []int) int {
	fmt.Println("Given Array: ", n, height)
	prev2 := 0
	prev := 0
	for i := 1; i < n; i++ {
		l := prev + int(math.Abs(float64(height[i]-height[i-1])))
		r := math.MaxInt32
		if i > 1 {
			r = prev2 + int(math.Abs(float64(height[i]-height[i-2])))
		}
		current := int(math.Min(float64(l), float64(r)))
		prev2 = prev
		prev = current
	}
	return prev
}

// Frog Jump Solution with K upto k jump
// means i, i+1, i+2....i+k
func FrogJumpWithKJump(n int, height []int) int {
	return 0
}

func FrogJumpWithKJumpRecursive(n int, height []int) int {
	fmt.Println("Given Array: ", n, height)
	prev2 := 0
	prev := 0
	for i := 1; i < n; i++ {
		l := prev + int(math.Abs(float64(height[i]-height[i-1])))
		r := math.MaxInt32
		if i > 1 {
			r = prev2 + int(math.Abs(float64(height[i]-height[i-2])))
		}
		current := int(math.Min(float64(l), float64(r)))
		prev2 = prev
		prev = current
	}
	return prev
}
