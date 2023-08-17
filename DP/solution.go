package main

import "fmt"

func main() {
	n := 16
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	//fmt.Println(dp, len(dp))
	fmt.Println(FibonacciUsingMemoization(16, dp))
	fmt.Println(FibonacciUsingTabulation(16))
	fmt.Println(FibonacciNoSC(16))
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
