package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	// triangle := [][]int{
	// 	{2},
	// 	{3, 4},
	// 	{6, 5, 7},
	// 	{4, 1, 8, 3},
	// }
	// fmt.Println(minimumTotal(triangle))
	strs := []string{"10", "0001", "111001", "1", "0"}
	// nums := [][]int{
	// 	{-1},
	// 	{2},
	// 	{3},
	// }

	fmt.Println(findMaxFormRecursiveApproach(strs, 5, 3))
}

/**
* Problem : 120. Triangle
* Given a triangle array, return the minimum path sum from top to bottom.
* Solution: Use bottom up approach and update the value of upper array while adding the min of below array¯
**/
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += findMin(triangle[i+1][j], triangle[i+1][j+1])
		}
		fmt.Println(triangle)
	}
	return triangle[0][0]
}

func findMin(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

/**
* Problem : 15. 3Sum
* Problem: Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
* Notice that the solution set must not contain duplicate triplets.
* Input: nums = [-1,0,1,2,-1,-4]
* Output: [[-1,-1,2],[-1,0,1]]
 */
func threeSumBruteForceSolution(nums []int) [][]int {
	fmt.Println("Given Array: ", nums)
	//foundZeroSum := false
	mapStoreTwosum := make([][]int, int(len(nums)/3))
	// fmt.Println(mapStoreTwosum)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					//foundZeroSum = true
					fmt.Printf("i=%d, j=%d, k=%d \n", i, j, k)
					mapStoreTwosum[i] = []int{nums[k], nums[i], nums[j]}
				}
			}

		}
	}
	fmt.Println(mapStoreTwosum)
	return mapStoreTwosum
}

/**
* Problem : 15. 3Sum
* Problem: Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
* Notice that the solution set must not contain duplicate triplets.
* Input: nums = [-1,0,1,2,-1,-4]
* Output: [[-1,-1,2],[-1,0,1]]
 */
func threeSumOptimalSolution(nums []int) [][]int {
	fmt.Println("Given Array: ", nums)
	sort.Ints(nums)
	fmt.Println("Sorted Array: ", nums)
	//foundZeroSum := false
	mapStoreTwosum := [][]int{}
	low := 0
	right := len(nums) - 1
	sum := 0

	// fmt.Println(mapStoreTwosum)
	for i := 0; i < len(nums)-2; i++ {
		sum = -nums[i]
		low = i + 1
		fmt.Println(i, nums[low], nums[right], sum)
		for low < right {
			//fmt.Println(i, low, right, nums[low]+nums[right], sum)
			if nums[low]+nums[right] == sum {
				//mapStoreTwosum[i] = []int{nums[i], nums[low], nums[right]}
				mapStoreTwosum = append(mapStoreTwosum, []int{nums[i], nums[low], nums[right]})
				//fmt.Println(nums[i], nums[low], nums[right])
				low++
				right--
			} else {
				right--
			}

			if nums[low]+nums[right] < sum {
				low++
			}

			//fmt.Println(nums[i], nums[low], nums[right])
		}
	}
	fmt.Println(mapStoreTwosum, sum)
	return mapStoreTwosum
}

/**
* Problem: 189. Rotate Array
* Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
* Input: nums = [1,2,3,4,5,6,7], k = 3
* Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
*/

func BruteForceRotate(nums []int, k int) {
	fmt.Println("Given Array:", nums)
	// for i := len(nums); i > len(nums)-k; i-- {
	// 	fmt.Println(i)
	// 	//nums = append(nums[i:], nums[:i]...)
	// 	nums[i+1] = nums[i]
	// }
	/** Brute force 1
	for i := 0; i < k; i++ {
		var j, last int
		length := len(nums)
		last = nums[length-1]
		for j = length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
	*/
	arrLen := len(nums)
	k = k % arrLen
	tmp := make([]int, k)
	m := 0
	for i := (arrLen - k); i < len(nums); i++ {
		tmp[m] = nums[i]
		m++
	}
	// tmpElem := nums[k%arrLen]
	for j := (arrLen - 1); j >= k; j-- {
		nums[j] = nums[j-k]
	}
	//nums[arrLen-1] = tmpElem
	for l := 0; l < k; l++ {
		nums[l] = tmp[l]
	}
	fmt.Println(nums, tmp)
}

/**
* Problem: 189. Rotate Array
* Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
* Input: nums = [1,2,3,4,5,6,7], k = 3
* Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
*/

func OptimalRotate(nums []int, k int) {
	k = k % len(nums) // get the modulo of k because if the array lenth 4 and we want rotate this 10 times that means 4+4+2.
	// Basically we rotate array for 4 times its return to the same as original array. And we do this another 4 time then returns the same.
	// So this 4+4+2 means will have to rotate the array only 2
	reverse(nums, 0, len(nums)-1-k)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, startIndex int, endIndex int) {

	if startIndex == endIndex {
		return
	}
	for startIndex < endIndex {
		nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
		startIndex++
		endIndex--
	}
}

/**
	Problem : 523. Continuous Subarray Sum
	Given an integer array nums and an integer k, return true if nums has a good subarray or false otherwise.
	A good subarray is a subarray where:

	its length is at least two, and
	the sum of the elements of the subarray is a multiple of k.
	Note that:

	A subarray is a contiguous part of the array.
	An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k
**/

func checkSubarraySum(nums []int, k int) bool {
	tmpMod := 0
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
		if i < 1 {
			continue
		}
		tmpMod = nums[i] % k

		if (tmpMod+nums[i])%k == 0 {
			return true
		}
	}
	return (totalSum % k) == 0
}

/**
* 167. Two Sum II - Input Array Is Sorted
* Input: numbers = [2,7,11,15], target = 9
* Output: [1,2]
* Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
 */

func twoSum(numbers []int, target int) []int {
	fmt.Println("Given Array: ", numbers)
	var i, j = 0, len(numbers) - 1
	for i < len(numbers) && j >= 0 {
		fmt.Println(numbers[i], numbers[j], i, j)
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
		if i >= j {
			break
		}

		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

/**
* Problem: 53. Maximum Subarray
* Given an integer array nums, find the subarray with the largest sum, and return its sum.
* Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
* Output: 6
* Explanation: The subarray [4,-1,2,1] has the largest sum 6.
* Solution: Kadanes alogorithm to calculate the max sum. The intuition of the algorithm is not to consider the subarray as a part of the answer if its sum is less than 0
 */

func maxSubArray(nums []int) int {
	largestSum := math.MinInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > largestSum {
			largestSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return largestSum

}

/**
* Problem: Print the subarray with maximum sum
* Given an integer array nums, find the subarray with the largest sum, and return its sum.
* Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
* Output:
* The subarray is: [4 -1 2 1 ]
* The maximum subarray sum is: 6
 */

func PrintMaxSubarrayofSum(nums []int) []int {
	var maxSum, sum, start, startPos, endPos = math.MinInt16, 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if sum == 0 {
			start = i
		}

		sum += nums[i]
		if sum > maxSum {
			startPos = start
			endPos = i
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	fmt.Println(start, startPos, endPos)
	return nums[startPos : endPos+1]
}

/*
*
* Problem: 2149. Rearrange Array Elements by Sign
* You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.
* You should rearrange the elements of nums such that the modified array follows the given conditions:
* Every consecutive pair of integers have opposite signs.
* For all integers with the same sign, the order in which they were present in nums is preserved.
* The rearranged array begins with a positive integer.
* Return the modified array after rearranging the elements to satisfy the aforementioned conditions.
* Input: nums = [3,1,-2,-5,2,-4]
Output: [3,-2,1,-5,2,-4]
Explanation:
The positive integers in nums are [3,1,2]. The negative integers are [-2,-5,-4].
The only possible way to rearrange them such that they satisfy all conditions is [3,-2,1,-5,2,-4].
Other ways such as [1,-2,2,-5,3,-4], [3,1,2,-2,-5,-4], [-2,3,-5,1,-4,2] are incorrect because they do not satisfy one or more conditions.
*/
func rearrangeArrayBruteForceSolution(nums []int) []int {
	var posArray, negArray = []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			posArray = append(posArray, nums[i])
		} else {
			negArray = append(negArray, nums[i])
		}
	}
	for i := 0; i < int(len(nums)/2); i++ {
		nums[i*2] = posArray[i]
		nums[i*2+1] = negArray[i]
	}
	return nums
}

/**
* Input: nums = [3,1,-2,-5,2,-4]
Output: [3,-2,1,-5,2,-4]
Explanation:
The positive integers in nums are [3,1,2]. The negative integers are [-2,-5,-4].
The only possible way to rearrange them such that they satisfy all conditions is [3,-2,1,-5,2,-4].
Other ways such as [1,-2,2,-5,3,-4], [3,1,2,-2,-5,-4], [-2,3,-5,1,-4,2] are incorrect because they do not satisfy one or more conditions.
*/

func rearrangeArrayBetter(nums []int) []int {
	var ansArray, posIndex, negIndex = make([]int, len(nums)), 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			ansArray[negIndex*2+1] = nums[i]
			negIndex++
		} else {
			ansArray[posIndex*2] = nums[i]
			posIndex++
		}
	}
	fmt.Println(ansArray)
	return ansArray
}

/**
* Problem: 31. Next Permutation
* Input: nums = [1,2,3]
* Output: [1,3,2]
 */

func nextPermutation(nums []int) {
	fmt.Println("Input Array", nums)
	var index, arrLen, j, k = -1, len(nums), 0, len(nums) - 1
	for i := arrLen - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}
	if index == -1 {
		for j < k {
			nums[j], nums[k] = nums[k], nums[j]
			j++
			k--
		}
	} else {
		for i := arrLen - 1; i >= 0; i-- {
			if nums[i] > nums[index] {
				nums[i], nums[index] = nums[index], nums[i]
				break
			}
		}
		j = index + 1
		for j < k {
			nums[j], nums[k] = nums[k], nums[j]
			j++
			k--
		}
	}
}

/**
* Leaders in array problem
* Given an array, print all the elements which are leaders. A Leader is an element that is greater than all of the elements on its right side in the array.
* Input : arr = [4, 7, 1, 0]
* Output: arr = [7, 1, 0]
* Solution: Two pointer approach
 */

func leaderInArray(nums []int) []int {
	fmt.Println("Given array: ", nums)
	leaderArray := []int{}
	//leaderArray = append(leaderArray, nums[len(nums)-1])
	var i, j = 0, len(nums) - 1
	for i < j {
		isLeader := false
		if nums[i] > nums[j] {
			j--
			isLeader = true
		} else {
			i++
			j = len(nums) - 1
			isLeader = false
		}
		if i == j {
			if isLeader {
				leaderArray = append(leaderArray, nums[i])
			}
			i++
			j = len(nums) - 1
		}
	}
	leaderArray = append(leaderArray, nums[j])
	return leaderArray
}

func leaderInArrayBetter(nums []int) []int {
	fmt.Println("Given array: ", nums)
	leaderArray := []int{}
	max := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > max {
			leaderArray = append(leaderArray, nums[i])
			max = nums[i]
		}
	}
	leaderArray = append(leaderArray, nums[len(nums)-1])
	return leaderArray
}

/*
* Problem: 128. Longest Consecutive Sequence
* Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
* You must write an algorithm that runs in O(n) time.
* Input: nums = [100,4,200,1,3,2]
* Output: 4
* Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
 */
func longestConsecutiveBruteForce(nums []int) int {
	longestStreak := 1
	if len(nums) == 0 {
		return len(nums)
	}
	for i := 0; i < len(nums); i++ {
		currentNumber := nums[i]
		count := 1
		for findInArray(currentNumber+1, nums) {
			currentNumber += 1
			count += 1
		}
		if count > longestStreak {
			longestStreak = count
		}
	}
	return longestStreak
}

func findInArray(elem int, nums []int) bool {
	for j := 0; j < len(nums); j++ {
		if nums[j] == elem {
			return true
		}
	}
	return false
}
func longestConsecutiveBetter(nums []int) int {
	//fmt.Println("Given array: ", nums)
	var longestStreak, tmpStreak = 1, 1
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	sort.Ints(nums)
	for k := 1; k < len(nums); k++ {
		if nums[k] == nums[k-1]+1 {
			tmpStreak++
		} else if nums[k] != nums[k-1] {
			tmpStreak = 1
		}
		if tmpStreak > longestStreak {
			longestStreak = tmpStreak
		}
	}
	fmt.Println(nums, longestStreak)
	return longestStreak
}

/*
* Solution: Have unordered set or map to store all the elem of the given array and try to to least number from the map.
* If this exist then start the iteration number + 1 if there increase the count and longeststreak as well.
 */
func longestConsecutiveOptimalSolution(nums []int) int {
	fmt.Println("Given array: ", nums)
	longestStreak := 1
	if len(nums) == 0 {
		return 0
	}
	numMap := make(map[int]struct{})
	for _, v := range nums {
		numMap[v] = struct{}{}
	}

	for k := range numMap {
		if _, ok := numMap[k-1]; !ok {
			leastNumber := k
			count := 1
			_, findNext := numMap[leastNumber+1]
			for findNext {
				leastNumber += 1
				count += 1
				_, findNext = numMap[leastNumber+1]
			}
			if count > longestStreak {
				longestStreak = count
			}
		}
	}
	//fmt.Println(numMap)
	return longestStreak
}

/*
* Problem : 73. Set Matrix Zeroes
* Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
* Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
* Output: [[1,0,1],[0,0,0],[1,0,1]]
* Solution Time Compexity: O(n^3), space complexity: 0(1)
 */

func setZeroesBruteForceSolution(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				markRow(i, n, matrix)
				markCol(j, m, matrix)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == math.MinInt64 {
				matrix[i][j] = 0
			}
		}
	}
}

func markRow(row int, n int, matrix [][]int) {
	for i := 0; i < n; i++ {
		if matrix[row][i] != 0 {
			matrix[row][i] = math.MinInt64
		}
	}
}

func markCol(col int, m int, matrix [][]int) {
	for i := 0; i < m; i++ {
		fmt.Println(i, col)
		if matrix[i][col] != 0 {
			matrix[i][col] = math.MinInt64
		}
	}
}

// Solution: Time complexity: O(n^2), Space Complexity: O(m) + O(n)
func setZeroesBetterSolution(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row := make([]int, m) // Take row array of m length
	col := make([]int, n) // Take col array of n length
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroesOptimalSolution(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	col0 := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // mark row as 0
				if j != 0 {
					matrix[0][j] = 0 // mark col as 0
				} else {
					col0 = 0
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != 0 {
				//check for row and column
				if matrix[0][j] == 0 || matrix[i][0] == 0 {
					matrix[i][j] = 0
				}
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if col0 == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

/**
* Problem: 48. Rotate Image
* You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise)
* You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
* Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
* Output: [[7,4,1],[8,5,2],[9,6,3]]
**/

func rotateBruteForceSolution(matrix [][]int) {
	var m, n = len(matrix), len(matrix[0])
	dummyMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		dummyMatrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dummyMatrix[j][n-1-i] = matrix[i][j]
			// matrix[j][n-1-i] = matrix[i][j]

			// if j == n-1 || (i == m-1 && j > 0) {
			// 	matrix[j][n-1-i] = dummyMatrix[i][j]
			// }

			// if i > 0 && j > i {
			// 	matrix[j][n-1-i] = dummyMatrix[i][j]
			// }
			// if j >= n-1-i {
			// 	fmt.Println(j, dummyMatrix[i][j], matrix[i][j])
			// }
			// if n-1-i == j && j >= i {
			// 	matrix[j][n-1-i] = dummyMatrix[i][j]
			// }

		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(" ", dummyMatrix[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(" ", matrix[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println(matrix)
}

/**
* Solution:
* 1. Transpose the matrix first which means col becomes row and row becomes col
* 2. Reverse the earch row and you will be achieve the answer.
 */
func rotateOptimalSolution(matrix [][]int) {
	var m, n = len(matrix), len(matrix[0])

	for i := 0; i < m-1; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j] // Transpose the matrix logic
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j] // reverse the matrix
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(" ", matrix[i][j])
		}
		fmt.Print("\n")
	}
}

/**
* Problem: 54. Spiral Matrix
* Given an m x n matrix, return all elements of the matrix in spiral order.
* Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
* Output: [1,2,3,6,9,8,7,4,5]
* Solution: right->bottom->left->top approach to traverse the matrix
 */

func spiralOrder(matrix [][]int) []int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(" ", matrix[i][j])
		}
		fmt.Print("\n")
	}
	var top, bottom, left, right, matrixVal = 0, len(matrix) - 1, 0, len(matrix[0]) - 1, []int{}
	for left <= right && top <= bottom {

		for i := left; i <= right; i++ {
			matrixVal = append(matrixVal, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			matrixVal = append(matrixVal, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrixVal = append(matrixVal, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrixVal = append(matrixVal, matrix[i][left])
			}
			left++
		}
	}
	return matrixVal
}

/**
* Problem: 560. Subarray Sum Equals K
* Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
* A subarray is a contiguous non-empty sequence of elements within an array.
* Input: nums = [1,1,1], k = 2
* Output: 2
 */
func subarraySumBruteForce(nums []int, k int) int {
	var count int
	//var totalSum int
	for i := 0; i < len(nums); i++ {
		var sum = nums[i]
		if sum == k {
			count += 1
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]

			if sum == k {
				count += 1
				if (i == 0 && j == len(nums)-1) && sum-nums[i] <= k {
					i = j
					break
				}
			}
		}
	}
	return count
}

/**
* Optimal solution using two pointer equidorectional approach
 */

func subarraySumOptimal(nums []int, k int) int {
	var count, prefixSum = 0, 0
	unorderedmMap := make(map[int]int, len(nums))
	unorderedmMap[0] = 1
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		remove := prefixSum - k
		count += unorderedmMap[remove]
		unorderedmMap[prefixSum] += 1
	}
	return count
}

/**
* Problem: 215. Kth Largest Element in an Array
* Given an integer array nums and an integer k, return the kth largest element in the array.
* Note that it is the kth largest element in the sorted order, not the kth distinct element.
* Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
* Output: 4
* This solution work only incase of positive integers
 */
func findKthLargestOptimal(nums []int, k int) int {
	fmt.Println("Given Array: ", nums)
	maxElement := nums[0]
	var count int
	for i := 0; i < len(nums); i++ {
		if maxElement < nums[i] {
			maxElement = nums[i]
		}
	}
	fmt.Println(maxElement)
	freqArray := make([]int, maxElement+1)

	for i := 0; i < len(nums); i++ {
		freqArray[nums[i]]++
	}

	for i := maxElement; i > 0; i-- {
		if freqArray[i] != 0 {
			count += freqArray[i]
			if int(count) >= k {
				return i
			}
		}
	}
	return 0
}

func findKthLargest(nums []int, k int) int {
	fmt.Println("Given Array: ", nums)
	if k < 1 || k > len(nums) {
		return -1 // Invalid k value
	}

	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}

	pivotIndex := partition(nums, start, end)

	if pivotIndex == k {
		return nums[pivotIndex]
	} else if k < pivotIndex {
		return quickSelect(nums, start, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, end, k)
	}
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start - 1

	for j := start; j < end; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[end] = nums[end], nums[i+1]
	fmt.Println(nums)
	return i + 1
}

func findPivotIndex(nums []int, start, end int) int {
	pivot := nums[end] // find the pivot and place its correct order
	j := start - 1

	for i := start; i < end; i++ {
		if nums[i] <= pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[j+1], nums[end] = nums[end], nums[j+1]
	//fmt.Println(nums)
	return j + 1
}

/**
* NAother solution is to check the constraint for elements val and create a slice with 2*num[i] contraint value.
* Next, store all the numbers there with constaraint addition nums[i] + 10^4 in this question.
* Next, check the frequency for each element and sum it up, untill you get sum >= k
 */

func findKthLargestAnotherOptimalSolution(nums []int, k int) int {
	count := make([]int, 20001)
	for _, num := range nums {
		count[num+10000]++
	}
	running := 0
	for i := len(count) - 1; i >= 0; i-- {
		if count[i] > 0 {
			running += count[i]
			if running >= k {
				return i - 10000
			}
		}
	}
	return -1
}

/**
* Problem: 347. Top K Frequent Elements
* Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
* Input: nums = [1,1,1,2,2,3], k = 2
* Output: [1,2]
* Intuition: counting sort algorithm
 */

func topKFrequent(nums []int, k int) []int {
	fmt.Println("Given Array: ", nums)
	ans := []int{}
	countArray := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		countArray[nums[i]]++
	}

	countFreq := make(map[int][]int, len(countArray))

	for key, freq := range countArray {
		countFreq[freq] = append(countFreq[freq], key)
	}

	for i := len(countArray) + 1; len(ans) <= k; i-- {
		ans = append(ans, countFreq[i]...)
		if len(ans) >= k {
			return ans[0:k]
		}
	}
	return ans
}

func findKthLargestHeap(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

type IntHeap []int
type any = interface{}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

/** TopKFrequentElements using Priority Queue Start here */

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value int, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

func topKFrequentUsingPriorityQueue(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(counter))
	i := 0
	for value, count := range counter {
		pq[i] = &Item{
			value:    value,
			priority: count,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	ans := []int{}
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		ans = append(ans, item.value)
	}
	return ans
}

/** TopKFrequentElements using Priority Queue Ends here */

/**
* Problem: 659. Split Array into Consecutive Subsequences
* You are given an integer array nums that is sorted in non-decreasing order.
* Determine if it is possible to split nums into one or more subsequences such that both of the following conditions are true:
* 1. Each subsequence is a consecutive increasing sequence (i.e. each integer is exactly one more than the previous integer).
* 2. All subsequences have a length of 3 or more.
* Return true if you can split nums according to the above conditions, or false otherwise.
* Example Input: nums = [1,2,3,3,4,5]
* Output: true
* Explanation: nums can be split into the following subsequences:
* [1,2,3,3,4,5] --> 1, 2, 3
* [1,2,3,3,4,5] --> 3, 4, 5
 */

func isPossible(nums []int) bool {
	var freqMap = map[int]int{}
	var vacancyMap = map[int]int{}
	if len(nums) < 3 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		if freqMap[nums[i]] <= 0 {
			continue
		} else if vacancyMap[nums[i]] > 0 {
			freqMap[nums[i]]--
			vacancyMap[nums[i]]--
			vacancyMap[nums[i]+1] = vacancyMap[nums[i]+1] + 1
		} else if freqMap[nums[i]] > 0 && freqMap[nums[i]+1] > 0 && freqMap[nums[i]+2] > 0 {
			freqMap[nums[i]]--
			freqMap[nums[i]+1]--
			freqMap[nums[i]+2]--
			vacancyMap[nums[i]+3] = vacancyMap[nums[i]+3] + 1
		} else {
			return false
		}
	}
	//fmt.Println(freqMap, vacancyMap)
	return true
}

/**
* Problem: 1296. Divide Array in Sets of K Consecutive Numbers
* Given an array of integers nums and a positive integer k, check whether it is possible to divide this array into sets of k consecutive numbers.
* Return true if it is possible. Otherwise, return false.
* Input: nums = [1,2,3,3,4,4,5,6], k = 4
* Output: true
* Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].
* Greedy Approach
 */

func isPossibleDivide(nums []int, k int) bool {
	arrLen := len(nums)
	var freqMap = map[int]int{}
	if arrLen < k*2 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < arrLen; i++ {
		freqMap[nums[i]]++
	}

	for i := 0; i < arrLen; i++ {
		if freqMap[nums[i]] <= 0 {
			continue
		} else if !CheckIfKthConsecutiveElem(nums, i, k, freqMap) {
			return false
		}
	}
	return true
}

func CheckIfKthConsecutiveElem(arr []int, j, k int, freqMap map[int]int) bool {
	for i := 0; i < k; i++ {
		if freqMap[arr[j]+i] <= 0 {
			return false
		}
		freqMap[arr[j]+i]--
	}
	return true
}

/**
* Problem: 2155. All Divisions With the Highest Score of a Binary Array
* You are given a 0-indexed binary array nums of length n. nums can be divided at index i (where 0 <= i <= n) into two arrays (possibly empty) numsleft and numsright:
* numsleft has all the elements of nums between index 0 and i - 1 (inclusive), while numsright has all the elements of nums between index i and n - 1 (inclusive).
*   If i == 0, numsleft is empty, while numsright has all the elements of nums.
*   If i == n, numsleft has all the elements of nums, while numsright is empty.
* The division score of an index i is the sum of the number of 0's in numsleft and the number of 1's in numsright.
* Return all distinct indices that have the highest possible division score. You may return the answer in any order.
* Input: nums = [0,0,1,0]
* Output: [2,4]
* Explanation: Division at index
- 0: numsleft is []. numsright is [0,0,1,0]. The score is 0 + 1 = 1.
- 1: numsleft is [0]. numsright is [0,1,0]. The score is 1 + 1 = 2.
- 2: numsleft is [0,0]. numsright is [1,0]. The score is 2 + 1 = 3.
- 3: numsleft is [0,0,1]. numsright is [0]. The score is 2 + 0 = 2.
- 4: numsleft is [0,0,1,0]. numsright is []. The score is 3 + 0 = 3.
* Indices 2 and 4 both have the highest possible division score 3.
* Note the answer [4,2] would also be accepted.
*/

func maxScoreIndices(nums []int) []int {
	fmt.Println("Given Array: ", nums)
	maxIndices := []int{}
	trackOs1sMap := make(map[int][]int)
	countZero, countOne, maxCountCompare := 0, 0, 0
	trackOs1sMap[0] = append(trackOs1sMap[0], countZero, countOne)
	if len(nums) < 2 && nums[0] == 1 {
		return []int{0}
	}
	if len(nums) < 2 && nums[0] == 0 {
		return []int{1}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countZero++
		} else {
			countOne++
		}
		trackOs1sMap[i+1] = append(trackOs1sMap[i+1], countZero, countOne)
	}

	// for i := 0; i < len(trackOs1sMap); i++ {
	// 	currentScore := trackOs1sMap[i][0] + trackOs1sMap[len(trackOs1sMap)-1][1] - trackOs1sMap[i][1]
	// 	if maxCountCompare < currentScore {
	// 		maxCountCompare = currentScore
	// 	}
	// }

	for i := 0; i <= len(trackOs1sMap)-1; i++ {
		currentScore := trackOs1sMap[i][0] + trackOs1sMap[len(trackOs1sMap)-1][1] - trackOs1sMap[i][1]
		//fmt.Println(i, currentScore, startIndex, endIndex)
		if maxCountCompare < currentScore {
			maxCountCompare = currentScore
			maxIndices = []int{i}
		} else if maxCountCompare == currentScore {
			maxIndices = append(maxIndices, i)
		}
	}
	//fmt.Println(trackOs1sMap)
	return maxIndices
}

func maxScoreIndicesBetterApproach(nums []int) []int {
	fmt.Println("Given Array: ", nums)
	maxIndices := []int{0}
	leftScore, rightScore, noOfOnes, maxScore := 0, 0, 0, 0

	for _, num := range nums {
		if num == 1 {
			noOfOnes++
		}
	}
	rightScore = noOfOnes
	maxScore = leftScore + rightScore

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			leftScore++
		} else if nums[i] == 1 {
			rightScore--
		}

		currentScore := leftScore + rightScore
		if currentScore > maxScore {
			maxScore = currentScore
			maxIndices = []int{i + 1}
		} else if currentScore == maxScore {
			maxIndices = append(maxIndices, i+1)
		}
	}
	//fmt.Println(trackOs1sMap)
	return maxIndices
}

/**
* Problem: 474. Ones and Zeroes
* You are given an array of binary strings strs and two integers m and n.
* Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
* A set x is a subset of a set y if all elements of x are also elements of y.
* Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
* Output: 4
* Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
 */

func findMaxForm(strs []string, m int, n int) int {
	fmt.Println("Given Array: ", strs, m, n)
	i, j, remaining0s, remaining1s, countMaxForm, tmp := 0, len(strs)-1, m, n, 0, 0
	for i < len(strs) && j >= 0 {
		countZero, countOne := countZerosOnesInString(strs[j])
		if countZero <= remaining0s && countOne <= remaining1s {
			remaining0s = remaining0s - countZero
			remaining1s = remaining1s - countOne
			tmp++
			j--
		} else if countZero > remaining0s || countOne > remaining1s {
			remaining0s = m
			remaining1s = n
			tmp = 0
			j = len(strs) - 1 - j
		} else {
			j--
		}
		if tmp > countMaxForm {
			countMaxForm = tmp
		}

		if j < i {
			j = len(strs) - 1
			i++
			remaining0s = m
			remaining1s = n
			tmp = 0
		}
	}
	return countMaxForm
}

func countZerosOnesInString(strs string) (int, int) {
	countZero, countOne := 0, 0
	for _, v := range strs {
		if v == '0' {
			countZero++
		} else {
			countOne++
		}
	}
	return countZero, countOne
}

func findMaxFormRecursiveApproach(strs []string, m int, n int) int {
	fmt.Println("Given Array: ", strs, m, n)
	var countValues = []int{}
	fmt.Println(helper(strs, m, n, 0, &countValues), len(countValues))
	return 0
}

func helper(strs []string, zero, one, index int, countValues *[]int) int {
	*countValues = append(*countValues, 1)
	var countStrZeroOnes = make([]int, 2)
	if index == len(strs) || zero+one == 0 {
		return 0
	}
	//fmt.Println(strings.Count(strs[index], "0"))
	countStrZeroOnes[0] += strings.Count(strs[index], "0")
	countStrZeroOnes[1] += (len(strs[index]) - countStrZeroOnes[0])

	var accept, reject int
	if zero >= countStrZeroOnes[0] && one >= countStrZeroOnes[1] {
		accept = helper(strs, zero-countStrZeroOnes[0], one-countStrZeroOnes[1], index+1, countValues) + 1
	} else {
		// reject the string
		reject = helper(strs, zero, one, index+1, countValues)
	}

	if reject > accept {
		return reject
	}

	return accept
}
