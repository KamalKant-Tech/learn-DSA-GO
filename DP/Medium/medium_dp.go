package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(maxJump([]int{0, 2, 5, 6, 7}))
	// fmt.Println(robTabulationII([]int{1, 2, 3, 1}))
	// fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	// 1, 2, 5
	// 3, 1, 1
	// 3, 3, 3

	// 18, 11, 19
	// 4, 13, 7
	// 1, 8, 13

	// 10, 40, 70
	// 20, 50, 80
	// 30, 60, 90
	// fmt.Println(NinjaTrainningTabulation(3, [][]int{{1, 2, 5}, {3, 1, 1}, {3, 3, 3}}))
	// fmt.Println(NinjaTrainningTabulation(3, [][]int{{18, 11, 19}, {4, 13, 7}, {1, 8, 13}}))
	// fmt.Println(NinjaTrainningTabulation(3, [][]int{{10, 40, 70}, {20, 50, 80}, {30, 60, 90}}))
	fmt.Println(maxProfitUsingRecursion([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitUsingMemoization([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitUsingTabulation([]int{7, 1, 5, 3, 6, 4}))
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

// Problem: 198. House Robber
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed, the only constraint stopping you from robbing each
// of them is that adjacent houses have security systems connected and it will automatically contact the
// police if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.

// The space complexity for the rob function is O(n) due to the recursion stack space.
// TC: O(n), SC:O(n)
func rob(nums []int) int {
	counter := 0
	return robHelper(len(nums)-1, nums, &counter)
}

func robHelper(i int, nums []int, counter *int) int {
	*counter++
	fmt.Println("Stack calls", *counter)
	if i < 0 {
		return 0
	}
	if i == 0 {
		return nums[0]
	}

	// Recursive relation
	return int(math.Max(float64(robHelper(i-1, nums, counter)), float64(nums[i]+robHelper(i-2, nums, counter))))
}

// Using Memoization
// TC: O(n), SC:O(n) for DP array
func robMemo(nums []int) int {
	dp := make([]int, len(nums))
	return robHelperMemo(len(nums)-1, nums, dp)
}

func robHelperMemo(i int, nums, dp []int) int {
	if i < 0 {
		return 0
	}
	if i == 0 {
		return nums[0]
	}

	if dp[i] != 0 {
		return dp[i]
	}

	// Recursive relation
	dp[i] = int(math.Max(float64(robHelperMemo(i-1, nums, dp)), float64(nums[i]+robHelperMemo(i-2, nums, dp))))
	return dp[i]
}

// TC: O(n), SC:O(n) for DP array
func robTabulation(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[1]), float64(nums[0])))
	n := len(nums)

	for i := 2; i < n; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}
	return dp[n-1]
}

// 213. House Robber II
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
// That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police
// if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.
// Intuition House can be robbed Nums[0] -> Nums[len(nums)-1] OR Nums[1] -> Nums[len(nums)]
// Using Recursion
func robII(nums []int) int {
	fmt.Println("Given Array: ", nums)
	if len(nums) == 1 {
		return nums[0]
	}
	startIndex := len(nums) - 1
	return int(math.Max(float64(robHelperII(startIndex, 1, nums)), float64(robHelperII(startIndex-1, 0, nums))))
}

func robHelperII(i, endIndex int, nums []int) int {
	if i < endIndex {
		return 0
	}

	if i == endIndex {
		return nums[i]
	}

	// Recursive relation
	return int(math.Max(float64(robHelperII(i-1, endIndex, nums)), float64(nums[i]+robHelperII(i-2, endIndex, nums))))
}

// Using Memoization
func robMemoII(nums []int) int {
	fmt.Println("Given Array: ", nums)
	if len(nums) == 1 {
		return nums[0]
	}
	startIndex := len(nums) - 1
	odddp := make([]int, len(nums))
	evendp := make([]int, len(nums))
	return int(math.Max(float64(robHelperMemoII(startIndex, 1, nums, odddp)), float64(robHelperMemoII(startIndex-1, 0, nums, evendp))))
}

func robHelperMemoII(i, endIndex int, nums []int, dp []int) int {
	if i < endIndex {
		return 0
	}

	if dp[i] != 0 {
		return dp[i]
	}

	if i == endIndex {
		return nums[i]
	}

	// Recursive relation
	dp[i] = int(math.Max(float64(robHelperMemoII(i-1, endIndex, nums, dp)), float64(nums[i]+robHelperMemoII(i-2, endIndex, nums, dp))))
	return dp[i]
}

// Using Tabulation
func robTabulationII(nums []int) int {
	dp := make([]int, len(nums))

	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < len(nums)-1; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}

	rdp := make([]int, len(nums))
	rdp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		rdp[i] = int(math.Max(float64(rdp[i-1]), float64(rdp[i-2]+nums[i])))
	}
	return int(math.Max(float64(rdp[len(nums)-1]), float64(dp[len(nums)-2])))
}

// 122. Best Time to Buy and Sell Stock II
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
// However, you can buy it then immediately sell it on the same day.
// Find and return the maximum profit you can achieve.
// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.

// TC: O(n ^ n-1), SC: O(n)
func maxProfitUsingRecursion(prices []int) int {
	fmt.Println("Given Array: ", prices)
	return maxProfitUsingRecursionHelper(prices, len(prices)-1)
}
func maxProfitUsingRecursionHelper(prices []int, index int) int {

	if index == 0 {
		return 0
	}

	maxProfit := 0
	for i := index - 1; i >= 0; i-- {
		if prices[index] > prices[i] {
			maxProfit = int(math.Max(float64(maxProfit), float64(prices[index]-prices[i]+maxProfitUsingRecursionHelper(prices, i))))
		} else {
			maxProfit = int(math.Max(float64(maxProfit), float64(maxProfitUsingRecursionHelper(prices, i))))
		}
	}

	return maxProfit
}

// Using Memoization
// TC : (n ^2), SC: O(n) + O(n) for DP
func maxProfitUsingMemoization(prices []int) int {
	fmt.Println("Given Array: ", prices)
	dp := make([]int, len(prices))
	return maxProfitUsingMemoizationHelper(prices, len(prices)-1, dp)
}
func maxProfitUsingMemoizationHelper(prices []int, index int, dp []int) int {

	if index == 0 {
		return 0
	}

	if dp[index] != 0 {
		return dp[index]
	}

	maxProfit := 0
	for i := index - 1; i >= 0; i-- {
		if prices[index] > prices[i] {
			maxProfit = int(math.Max(float64(maxProfit), float64(prices[index]-prices[i]+maxProfitUsingMemoizationHelper(prices, i, dp))))
		} else {
			maxProfit = int(math.Max(float64(maxProfit), float64(maxProfitUsingMemoizationHelper(prices, i, dp))))
		}

		dp[index] = maxProfit
	}

	return maxProfit
}

// Using Tabulation
// TC : (n), SC: O(n) for DP
func maxProfitUsingTabulation(prices []int) int {
	fmt.Println("Given Array: ", prices)
	dp := make([]int, len(prices))
	dp[0] = 0

	for i := 1; i < len(prices); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-1]+prices[i]-prices[i-1])))
	}
	return dp[len(prices)-1]
}

// Optimal Solution Using Greedy approach
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}

	profit := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// 2D DP questions
// Ninjas Training
// Ninja is planing this ‘N’ days-long training schedule. Each day, he can perform any one of these three activities.
// (Running, Fighting Practice or Learning New Moves). Each activity has some merit points on each day.
// As Ninja has to improve all his skills, he can’t do the same activity in two consecutive days.
// Can you help Ninja find out the maximum merit points Ninja can earn?
// You are given a 2D array of size N*3 ‘POINTS’ with the points corresponding to each day and activity.
// Your task is to calculate the maximum number of merit points that Ninja can earn.
// For Example
// If the given POINTS array is
// [ [1, 2, 5],
//
//	[3 ,1 ,1],
//	[3, 3, 3]
//
// ],
// the answer will be 11 as 5 + 3 + 3.
func NinjaTrainning(n int, points [][]int) int {
	fmt.Println("Given Array: ", points)
	return NinjaTrainningHelper(n, len(points[0])-1, n-1, points)
}

func NinjaTrainningHelper(n, index, lastIndex int, points [][]int) int {
	if index == 0 {
		max := 0
		for i := 0; i < n; i++ {
			if i != lastIndex {
				max = int(math.Max(float64(max), float64(points[index][i])))
			}
		}
		return max
	}

	max := 0
	for i := 0; i < n; i++ {
		if i != lastIndex {
			max = int(math.Max(float64(max), float64(points[index][lastIndex]+NinjaTrainningHelper(n, index-1, i, points))))
		}
	}
	return max
}

// Using Memoization
// TC: 0(n * m) * n, SC: O(n) + O(n * m)
func NinjaTrainningMemo(n int, points [][]int) int {
	fmt.Println("Given Array: ", points)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return NinjaTrainningMemoHelper(n, len(points[0])-1, n-1, points, dp)
}

func NinjaTrainningMemoHelper(n, index, lastIndex int, points, dp [][]int) int {

	if index == 0 {
		max := 0
		for i := 0; i < n; i++ {
			if i != lastIndex {
				max = int(math.Max(float64(max), float64(points[index][i])))
			}
		}

		return max
	}

	if dp[index][lastIndex] != 0 {
		return dp[index][lastIndex]
	}

	max := 0
	for i := 0; i < n; i++ {
		if i != lastIndex {
			max = int(math.Max(float64(max), float64(points[index][lastIndex]+NinjaTrainningMemoHelper(n, index-1, i, points, dp))))
		}
	}

	dp[index][lastIndex] = max
	return dp[index][lastIndex]
}

// Using Tabulation
func NinjaTrainningTabulation(n int, points [][]int) int {
	fmt.Println("Given Array: ", points)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = int(math.Max(float64(points[0][1]), float64(points[0][2])))
	dp[0][1] = int(math.Max(float64(points[0][0]), float64(points[0][2])))
	dp[0][2] = int(math.Max(float64(points[0][0]), float64(points[0][1])))

	for day := 1; day < n; day++ {
		for activity := 0; activity < n; activity++ {
			dp[day][activity] = 0
			for task := 0; task < n; task++ {
				if task != activity {
					dp[day][activity] = int(math.Max(float64(dp[day][activity]), float64(points[day][activity]+dp[day-1][task])))
				}
			}
		}
	}
	return dp[n-1][2]
}
