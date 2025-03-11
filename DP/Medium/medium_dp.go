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
	// fmt.Println(maxProfitUsingRecursion([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(maxProfitUsingMemoization([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(maxProfitWithCooldown([]int{1, 2, 3, 0, 2}))
	// fmt.Println(numSquaresRecursion(18))
	// fmt.Println(numSquaresMemo(18))
	// fmt.Println(numSquaresTabulation(18))
	fmt.Println(totalUniquePathsMemo(2, 6))
	fmt.Println(totalUniquePathsTabulation(2, 6))
	fmt.Println(totalUniquePathsSpaceOptimization(2, 6))
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

// 309. Best Time to Buy and Sell Stock with Cooldown
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// Find the maximum profit you can achieve. You may complete as many transactions as you
// like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
// After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
// Input: prices = [1,2,3,0,2]
// Output: 3
// Explanation: transactions = [buy, sell, cooldown, buy, sell]

// Need to revisit this below Logic
func maxProfitWithCooldown(prices []int) int {
	fmt.Println("Given array: ", prices)
	memo := make(map[[2]int]int)
	return maxProfitWithCooldownUsingRecusion(prices, 0, false, memo)
}

func maxProfitWithCooldownUsingRecusion(prices []int, day int, holding bool, memo map[[2]int]int) int {

	// Base case: If we are out of days, profit is 0
	if day >= len(prices) {
		return 0
	}

	// Create a key for memoization
	key := [2]int{day, boolToInt(holding)}
	if val, exists := memo[key]; exists {
		return val
	}

	// Recursive cases
	var result int
	if holding {
		// Option 1: Sell the stock today
		sell := prices[day] + maxProfitWithCooldownUsingRecusion(prices, day+2, false, memo) // Cooldown for next day
		// Option 2: Do nothing
		hold := maxProfitWithCooldownUsingRecusion(prices, day+1, true, memo)
		result = max(sell, hold)
	} else {
		// Option 1: Buy the stock today
		buy := -prices[day] + maxProfitWithCooldownUsingRecusion(prices, day+1, true, memo)
		// Option 2: Do nothing
		skip := maxProfitWithCooldownUsingRecusion(prices, day+1, false, memo)
		result = max(buy, skip)
	}

	// Store the result in the memoization map
	memo[key] = result
	return result
}

// Helper function to convert boolean to int
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Helper function to find the max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 279. Perfect Squares
// Given an integer n, return the least number of perfect square numbers that sum to n.
// A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself.
// For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

// Greedy approach
// func numSquaresHelper(n int, nCount, count int, counter int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	maxCount := count
// 	for i := counter; i <= nCount; i++ {
// 		iSquare := i * i

// 		// if iSquare <= n {
// 		// 	maxCount++
// 		// 	fmt.Println(iSquare, n, nCount, i, maxCount)
// 		// 	return int(math.Max(float64(maxCount), float64(numSquaresHelper(n-iSquare, nCount, maxCount, counter))))
// 		// }

// 		maxCount++
// 		fmt.Println(iSquare, n, nCount, i, maxCount)
// 		return int(math.Max(float64(maxCount), float64(numSquaresHelper(n-iSquare, nCount, maxCount, counter))))
// 	}
// 	return maxCount
// }

// Using Recursion
func numSquaresRecursion(n int) int {
	memo := make([]int, n+1)
	return numSquaresHelperRecursion(n, memo)
}

// Using Recursion
func numSquaresHelperRecursion(n int, memo []int) int {
	if n < 4 {
		return n
	}

	ans := n

	for i := 1; i*i <= n; i++ {
		square := i * i
		ans = int(math.Min(float64(ans), float64(1+numSquaresHelperMemo(n-square, memo))))
	}

	return ans
}

// Using Memo
func numSquaresMemo(n int) int {
	memo := make([]int, n+1)
	return numSquaresHelperMemo(n, memo)
}

// Using Memo
func numSquaresHelperMemo(n int, memo []int) int {
	if n < 4 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	ans := n

	for i := 1; i*i <= n; i++ {
		square := i * i
		ans = int(math.Min(float64(ans), float64(1+numSquaresHelperMemo(n-square, memo))))
	}

	memo[n] = ans
	return ans
}

func numSquaresTabulation(n int) int {
	// Create a dp array to store the minimum number of perfect squares
	dp := make([]int, n+1)

	// Base case: 0 can be formed with 0 squares
	dp[0] = 0

	// Fill the dp array
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32 // Initialize with a large value
		for j := 1; j*j <= i; j++ {
			square := j * j
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-square])))
		}
	}
	return dp[n]
}

// Problem: Total Unique Paths
// Link: https://www.youtube.com/watch?v=sdE0A2Oxofw&list=PLgUwDviBIf0qUlt5H_kiKYaNSqJ81PMMY&index=11
// You are present at point ‘A’ which is the top-left cell of an M X N matrix, your destination is point ‘B’,
// which is the bottom-right cell of the same matrix. Your task is to find the total number of unique paths
// from point ‘A’ to point ‘B’.In other words, you will be given the dimensions of the matrix as integers ‘M’ and ‘N’,
// your task is to find the total number of unique paths from the cell MATRIX[0][0] to MATRIX['M' - 1]['N' - 1].

// To traverse in the matrix, you can either move Right or Down at each step.
// For example in a given point MATRIX[i] [j], you can move to either MATRIX[i + 1][j] or MATRIX[i][j + 1].
// TC: 2^m*n, SC: O((m-1) + (n-1))
func totalUniquePaths(m, n int) int {

	return totalUniquePathsHelper(m-1, n-1)
}

func totalUniquePathsHelper(row, col int) int {
	// base case if we reach to target
	if row == 0 && col == 0 {
		return 1
	}
	// if row and col exceeding the boundary means if less than 0 then return 0
	if row < 0 || col < 0 {
		return 0
	}

	// Since we can move to right and bottom, In this recursion will have to go just reverse i.e up and left cause we started with m and n
	up := totalUniquePathsHelper(row-1, col)   // row will reduce cause going up
	left := totalUniquePathsHelper(row, col-1) // col will reduce cause going left

	return up + left
}

// Using Momoization
// TC: O(m*n), SC: O((m-1) + (n-1)) + O(n * m)
func totalUniquePathsMemo(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	fmt.Println("Dp Array: ", dp)
	return totalUniquePathsMemoHelper(m-1, n-1, dp)
}

func totalUniquePathsMemoHelper(row, col int, dp [][]int) int {
	// base case if we reach to target
	if row == 0 && col == 0 {
		return 1
	}
	// if row and col exceeding the boundary means if less than 0 then return 0
	if row < 0 || col < 0 {
		return 0
	}

	if dp[row][col] != -1 {
		return dp[row][col]
	}
	// Since we can move to right and bottom, In this recursion will have to go just reverse i.e up and left cause we started with m and n
	up := totalUniquePathsMemoHelper(row-1, col, dp)   // row will reduce cause going up
	left := totalUniquePathsMemoHelper(row, col-1, dp) // col will reduce cause going left

	dp[row][col] = up + left
	return dp[row][col]
}

// Using Tabulation
// TC: O(m*n), SC: O(n * m)
func totalUniquePathsTabulation(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			up := 0
			left := 0
			if row == 0 && col == 0 {
				continue
			} else {
				if row > 0 {
					up = dp[row-1][col]
				}

				if col > 0 {
					left = dp[row][col-1]
				}
				dp[row][col] = up + left
			}
		}
	}
	return dp[m-1][n-1]
}

// Using Space Optimization
// TC: O(m*n), SC: O(n)
func totalUniquePathsSpaceOptimization(m, n int) int {
	prevRow := make([]int, n) // take one dummy row for to refer the the up values
	for row := 0; row < m; row++ {
		currentRow := make([]int, n) // Take one more dummy array to keep track of current row calculation and store it.
		for col := 0; col < n; col++ {
			up := 0
			left := 0
			if row == 0 && col == 0 {
				currentRow[0] = 1
				continue
			} else {
				if row > 0 {
					up = prevRow[col]
				}

				if col > 0 {
					left = currentRow[col-1]
				}
				currentRow[col] = up + left
			}
		}
		prevRow = currentRow
	}
	return prevRow[n-1]
}
