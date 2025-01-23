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
	//strs := []string{"10", "0001", "111001", "1", "0"}
	// nums := [][]int{
	// 	{-1},
	// 	{2},
	// 	{3},
	// }
	// nums := []int{1, 2, 4, 8, 9, -2, -7, 3}
	// fmt.Println(letterCombinations("234"))
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// // fmt.Println(groupAnagrams(strs))
	// fmt.Println(groupAnagramsUsingHashMap(strs))
	// fmt.Println(checkTwoStringAnnagram("eat", "tea"))
	// fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	s := "aaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa"}
	fmt.Println(wordBreak(s, wordDict))
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
	var storeZeroOneAtIndex = make([][][]int, m+1)
	for i := 0; i < len(storeZeroOneAtIndex); i++ {
		storeZeroOneAtIndex[i] = make([][]int, n+1)
		for j := 0; j < len(storeZeroOneAtIndex[i]); j++ {
			storeZeroOneAtIndex[i][j] = make([]int, len(strs))
		}
	}
	return helper(strs, m, n, 0, storeZeroOneAtIndex)
}

func helper(strs []string, zero, one, index int, storeZeroOneAtIndex [][][]int) int {
	//*storeZeroOneAtIndex = append(*storeZeroOneAtIndex, 1)
	var countStrZeroOnes = make([]int, 2)
	if index == len(strs) || zero+one == 0 {
		return 0
	}

	if storeZeroOneAtIndex[zero][one][index] > 0 {
		return storeZeroOneAtIndex[zero][one][index]
	}
	//fmt.Println(strings.Count(strs[index], "0"))
	countStrZeroOnes[0] += strings.Count(strs[index], "0")
	countStrZeroOnes[1] += (len(strs[index]) - countStrZeroOnes[0])

	var accept, reject int
	if zero >= countStrZeroOnes[0] && one >= countStrZeroOnes[1] {
		accept = helper(strs, zero-countStrZeroOnes[0], one-countStrZeroOnes[1], index+1, storeZeroOneAtIndex) + 1
	}

	// reject the string
	reject = helper(strs, zero, one, index+1, storeZeroOneAtIndex)
	storeZeroOneAtIndex[zero][one][index] = int(math.Max(float64(reject), float64(accept)))
	return storeZeroOneAtIndex[zero][one][index]
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

/**
* Using Dynamic Programming
 */
func findMaxFormOptimalApproach(strs []string, m int, n int) int {
	fmt.Println("Given Array:", strs, m, n)
	var storeZeroOnemaxForm = [101][101]int{}
	// for i := 0; i < len(storeZeroOnemaxForm); i++ {
	// 	storeZeroOnemaxForm[i] = make([]int, n+1)
	// }

	for _, str := range strs {
		countZero := strings.Count(str, "0")
		countOne := (len(str) - countZero)
		for zero := m; zero >= countZero; zero-- {
			for one := n; one >= countOne; one-- {
				storeZeroOnemaxForm[zero][one] = max(storeZeroOnemaxForm[zero][one], storeZeroOnemaxForm[zero-countZero][one-countOne]+1)
			}
		}
	}
	return storeZeroOnemaxForm[m][n]
}

/**
* Problem: 712. Minimum ASCII Delete Sum for Two Strings
* Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.
* Input: s1 = "sea", s2 = "eat"
* Output: 231
* Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
* Deleting "t" from "eat" adds 116 to the sum.
* At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
 */
func minimumDeleteSum(s1 string, s2 string) int {
	//loop through both the string and get ascii char values for each character
	// add the ascii char values for each char
	fmt.Println("Here")
	minSum := 0
	tmpCount := 0
	strTmp := ""
	for i, str := range s1 {
		strTmp = s1[i:]
		if !strings.Contains(s2, string(str)) {
			minSum += int(str)
		} else if strings.Count(strTmp, string(str)) > strings.Count(s2, string(str)) {
			minSum += int(str)
		} else {
			//fmt.Println(s1[i:])
			if len([]rune(strTmp)) <= i {
				strTmp = s1[i-1:]
			}
			if tmpCount < int(str) && strings.Replace(strTmp, string(str), "", 1) == strings.Replace(s2, string(str), "", 1) {
				tmpCount = int(str)
			}
		}
	}
	for _, str := range s2 {
		if !strings.Contains(s1, string(str)) {
			minSum += int(str)
		}
	}
	fmt.Println(tmpCount, minSum)
	return minSum
}

// Compare string at the end
func minimumDeleteSumBetterApproach(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		dp[i][n] = dp[i+1][n] + int(s1[i])
	}
	for j := n - 1; j >= 0; j-- {
		dp[m][j] = dp[m][j+1] + int(s2[j])
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}

	return dp[0][0]
}

// Compare string at the start
func minimumDeleteSumAnotherApproach(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		dp[i+1][0] = dp[(i+1)-1][0] + int(s1[i])
	}
	for j := 0; j < n; j++ {
		dp[0][j+1] = dp[0][(j+1)-1] + int(s2[j])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[m][n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Find min delete sum using recursive approach
func minimumDeleteSumRecApproach(s1 string, s2 string) int {
	return MinRec(s1, s2, len(s1), len(s2))
}

func MinRec(s1 string, s2 string, n, m int) int {
	if m == 0 {
		return getAtIndex(s1, n)
	}
	if n == 0 {
		return getAtIndex(s2, m)
	}
	fmt.Println(string(s1[n-1]), string(s2[m-1]), n, m)
	if s1[n-1] == s2[m-1] {
		return MinRec(s1, s2, n-1, m-1)
	} else {
		//fmt.Println("Print S1 : ", int(s1[n-1]))
		ans1 := int(s1[n-1]) + MinRec(s1, s2, n-1, m)
		//fmt.Println("Print ans1 : ", ans1)
		ans2 := int(s2[m-1]) + MinRec(s1, s2, n, m-1)
		fmt.Println("Print ans1 : ", ans1)
		fmt.Println("Print ans2 : ", ans2)
		return min(ans1, ans2)
	}
}

func getAtIndex(s string, idx int) int {
	sum := 0
	for i := 0; i < idx; i++ {
		sum += int(s[i])
	}
	return sum
}

/**
* Problem: 300. Longest Increasing Subsequence
* Given an integer array nums, return the length of the longest strictly increasing subsequence
* Input: nums = [10,9,2,5,3,7,101,18]
* Output: 4
* Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
 */

// func lengthOfLIS(nums []int) int {
// 	fmt.Println("Given Array: ", nums)
// 	mapOfNums := make(map[int][]int)
// 	trackElemCmp := make(map[int]bool)
// 	totalLen := 0
// 	i := 0
// 	j := 1
// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	for i < len(nums) && j < len(nums) {

// 		if _, ok := mapOfNums[nums[i]]; !ok {
// 			mapOfNums[nums[i]] = []int{nums[i]}
// 		}

// 		if nums[i] < nums[j] {
// 			//fmt.Println(mapOfNums[nums[i]][len(mapOfNums[nums[i]])-1], nums[j], mapOfNums[nums[i]][len(mapOfNums[nums[i]])-1] < nums[j])
// 			if len(mapOfNums[nums[i]]) > 0 && mapOfNums[nums[i]][len(mapOfNums[nums[i]])-1] < nums[j] {
// 				mapOfNums[nums[i]] = append(mapOfNums[nums[i]], nums[j])
// 			} else if !trackElemCmp[nums[i]] {
// 				//fmt.Println(mapOfNums[nums[i]][:len(mapOfNums[nums[i]])])
// 				mapOfNums[nums[i]] = append(mapOfNums[nums[i]][:len(mapOfNums[nums[i]])-1], nums[j])
// 			}
// 		}
// 		if totalLen < len(mapOfNums[nums[i]]) {
// 			totalLen = len(mapOfNums[nums[i]])
// 		}
// 		if j >= len(nums)-1 {
// 			trackElemCmp[nums[i]] = true
// 			i = i + 1
// 			j = i + 1
// 		} else {
// 			j++
// 		}
// 	}
// 	fmt.Println(mapOfNums)
// 	return totalLen
// }

/**
* 55. Jump Game
* You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
* Return true if you can reach the last index, or false otherwise.
* Input: nums = [2,3,1,1,4]
* Output: true
* Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
 */

func canJump(nums []int) bool {
	arrLen := len(nums)
	maxIndex := 0
	if arrLen == 1 {
		return true
	}

	for i := 0; i < arrLen-1 && maxIndex >= i; i++ {

		if maxIndex < i+nums[i] {
			maxIndex = i + nums[i]
		}

		if maxIndex >= arrLen-1 {
			return true
		}
	}
	return false
}

/**
* 45. Jump Game II
* You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
* Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:
* 0 <= j <= nums[i] and i + j < n
* Input: nums = [2,3,1,1,4]
* Output: 2
* Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
 */

func jump(nums []int) int {
	return jumpRecursive(nums, 0, len(nums)-1)
}

func jumpRecursive(nums []int, currentIndex, dest int) int {
	if currentIndex >= dest {
		return 0
	}

	tmpMinJump := math.MaxInt32

	for i := 1; i <= nums[currentIndex]; i++ {
		tmpJump := 1 + jumpRecursive(nums, currentIndex+i, dest)
		if tmpMinJump > tmpJump {
			tmpMinJump = tmpJump
		}
	}

	return tmpMinJump
}

/**
* Problem: 17. Letter Combinations of a Phone Number
* Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
* A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
* Image: https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png
* Example:
* Input: digits = "23"
* Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 */
func letterCombinations(digits string) []string {
	var result []string
	phoneMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return result
	}
	var backtrack func(string, string)
	backtrack = func(combination string, nextDigits string) {
		if len(nextDigits) == 0 {
			result = append(result, combination)
		} else {
			for _, letter := range phoneMap[string(nextDigits[0])] {
				backtrack(combination+letter, nextDigits[1:])
			}
		}
	}
	backtrack("", digits)
	return result
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

/*
*
* Problem: 49. Group Anagrams
* Given an array of strings strs, group the anagrams together. You can return the answer in any order.
* A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
* Image: https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png
* Example:
* Input: strs = ["eat","tea","tan","ate","nat","bat"]
* Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
* Explanation:
  - There is no string in strs that can be rearranged to form "bat".
  - The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
  - The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
*/
func groupAnagrams(strs []string) [][]string {
	fmt.Println("Given Array: ", strs)
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		fmt.Println("Sorted string: ", str, sortedStr)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	fmt.Println(anagramMap)
	result := [][]string{}
	for _, anagrams := range anagramMap {
		result = append(result, anagrams)
	}
	return result
}

// another solution using hashmap

func groupAnagramsUsingHashMap(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		// fmt.Println(k, s)
		mp[k] = append(mp[k], s)
	}
	fmt.Println(mp)
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

// checkTwoStringAnnagram checks if two strings are anagrams of each other.
// An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Parameters:
//   - s: The first string to be checked.
//   - t: The second string to be checked.
//
// Input: s = "anagram", t = "nagaram"
// Output: true
// Returns: bool
func checkTwoStringAnnagram(s string, t string) bool {
	if len(s) != len(t) {
		fmt.Println("The strings are not anagrams of each other.")
		return false
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		tMap[v]++
	}
	for k, v := range sMap {
		if tMap[k] != v {
			fmt.Println("The strings are not anagrams of each other.")
			return false
		}
	}
	fmt.Println("The strings are anagrams of each other.")
	return true
}

/**
* Problem: 907. Sum of Subarray Minimums
* Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr. Since the answer may be large, return the answer modulo 10^9 + 7
* Input: arr = [3,1,2,4]
* Output: 17
* Explanation:
* Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
* Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
* Sum is 17.
 */

func sumSubarrayMinsBruteForce(arr []int) int {
	const mod = 1e9 + 7
	// result := [][]int{}
	// 0, 0,0 | 0,1 | 0,2 | 0,3 => [[3], [3,1], [3,1,2], [3,1,2,4]]
	// 1, 1,1 | 1,2,| 1,3 => [[1], [1,2], [1,2,4]]
	// 2, 2,2 | 2,3 => [[2], [2,4]]
	// 3, 3,3 => [[4]]
	minSum := 0
	for i := range arr {
		minimum := math.MaxInt64
		if i == 0 {
			fmt.Println(minimum)
		}
		for j := i + 1; j <= len(arr); j++ {
			// result = append(result, arr[i:j])
			tmpArr := arr[i:j]
			if len(tmpArr) > 1 {

				for k := 0; k < len(tmpArr); k++ {
					if tmpArr[k] < minimum {
						minimum = tmpArr[k]
					}
				}
				minSum += minimum
			} else {
				minSum += tmpArr[0]
			}
		}
		// fmt.Println(minSum)
		// fmt.Println(arr)
		// minSum += arr[i]
	}

	// fmt.Println(result, minSum)
	return minSum % int(mod)
}

func sumSubarrayMins(arr []int) int {

	// mod is a constant used for taking the result modulo (10^9 + 7).
	// n is the length of the input array arr.
	// stack is used to keep track of indices while processing the array.
	// prev and next arrays are used to store the distances to the previous and next smaller elements for each element in arr.

	const mod = 1e9 + 7
	n := len(arr)
	stack := []int{}
	prev := make([]int, n)
	next := make([]int, n)

	// Initialize prev and next arrays. Initially, prev[i] is set to i + 1 and next[i] is set to n - i.
	for i := 0; i < n; i++ {
		prev[i] = i + 1
		next[i] = n - i
	}

	// This loop calculates the prev array. For each element arr[i], it finds the distance to the previous smaller element.
	// If the stack is not empty and the top element of the stack is greater than arr[i], pop the stack.
	// If the stack is not empty after popping, set prev[i] to the distance between i and the top element of the stack.
	// If the stack is empty, set prev[i] to i + 1.
	// Push the current index i onto the stack.

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prev[i] = i - stack[len(stack)-1]
		} else {
			prev[i] = i + 1
		}
		stack = append(stack, i)
	}

	// This loop calculates the next array. For each element arr[i], it finds the distance to the next smaller element.
	// If the stack is not empty and the top element of the stack is greater than or equal to arr[i], pop the stack.
	// If the stack is not empty after popping, set next[i] to the distance between the top element of the stack and i.
	// If the stack is empty, set next[i] to n - i.
	// Push the current index i onto the stack.

	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			next[i] = stack[len(stack)-1] - i
		} else {
			next[i] = n - i
		}
		stack = append(stack, i)
	}

	// Initialize result to 0.
	// For each element arr[i], add the product of arr[i], prev[i], and next[i] to result, taking modulo (10^9 + 7).
	result := 0
	for i := 0; i < n; i++ {
		result = (result + arr[i]*prev[i]*next[i]) % int(mod)
	}

	return result
}

/*
*
* Problem: 139. Word Break
* Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
* Note that the same word in the dictionary may be reused multiple times in the segmentation.
* Example:
* Input: s = "leetcode", wordDict = ["leet","code"]
* Output: true
* Explanation: Return true because "leetcode" can be segmented as "leet code".

* Input: s = "applepenapple", wordDict = ["apple","pen"]
* Output: truex

* Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
* Output: false
 */
func wordBreak(s string, wordDict []string) bool {
	fmt.Println("Given Array: ", s, wordDict)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	return wordBreakRecursive(s, wordSet, 0, make(map[int]bool))
}

func wordBreakRecursive(s string, wordSet map[string]bool, start int, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	// fmt.Println(memo)
	if val, exists := memo[start]; exists {
		return val
	}
	// start = 0, end = 1 to 8
	// 0,1 | 0,2 | 0,3 | 0,4 | 0,5 | 0,6 | 0,7 | 0,8
	// start = 1, end = 2 to 8
	// 1,2 | 1,3 | 1,4 | 1,5 | 1,6 | 1,7 | 1,8
	// start = 2, end = 3 to 8
	// 2,3 | 2,4 | 2,5 | 2,6 | 2,7 | 2,8
	// start = 3, end = 4 to 8
	// 3,4 | 3,5 | 3,6 | 3,7 | 3,8
	// start = 4, end = 5 to 8
	// 4,5 | 4,6 | 4,7 | 4,8
	// start = 5, end = 6 to 8
	// 5,6 | 5,7 | 5,8
	// start = 6, end = 7 to 8
	// 6,7 | 6,8
	// start =7, end = 8 to 8
	// 7,8

	// wordBreakRecursive(0)
	// wordBreakRecursive(1)
	for end := start + 1; end <= len(s); end++ {
		fmt.Println(start, end, wordSet[s[start:end]])
		if wordSet[s[start:end]] && wordBreakRecursive(s, wordSet, end, memo) {
			fmt.Println("String: ", s[start:end])
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}

func wordBreakOptimal(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// dp[j] - true
			// wordset["c"] - false
			// i = 2, j = 0
			// dp[0] - true
			// wordset["ca"] - false
			// dp[1] - true
			// wordset["a"] - false

			// i = 3, j =0
			// dp[0] - true
			// wordset["cat"] - true
			// dp[3] = true

			// i = 4, j =0
			// dp[0] - true
			// wordset["cats"] - true
			// dp[4] = true

			// i =5 , j = 0
			// dp[0] - true
			// wordset["catsa"] - false
			// dp[1] - false
			// wordset["atsa"] - false
			// dp[2] - false
			// wordset["tsa"] - false
			// dp[3] - true
			// wordset["sa"] - false
			// dp[4] - true
			// wordset["a"] - false

			// i = 5, j= 0
			// dp[0] - true
			// wordset["catsan"] - false
			// dp[1] - false
			// wordset["atsan"] - false
			// dp[2] - false
			// wordset["tsan"] - false
			// dp[3] - true
			// wordset["san"] - false
			// dp[4] - true
			// wordset["an"] - false
			// dp[5] - false
			// wordset["n"] - false
			// *
			// *
			// *
			// *
			// *
			// i =10, j =0
			// dp[0] - true
			// wordset["catsanddog"] - false
			// dp[1] - false
			// wordset["atsanddog"] - false
			// dp[2] - false
			// wordset["tsanddog"] - false
			// dp[3] - true
			// wordset["sanddog"] - false
			// dp[4] - true
			// wordset["anddog"] - false
			// dp[5] - false
			// wordset["nddog"] - false
			// dp[6] - false
			// wordset["ddog"] - false
			// dp[7] - true
			// wordset["dog"] - true

			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// func wordBreak(s string, wordDict []string) bool {
// 	fmt.Println("Given Array: ", s, wordDict)
// 	// cats -> index 0 + 4 = 4
// 	// dog -> index 6 =  4 + 3 = 7
// 	// dog -> index 3 + 3 = 7 + 3 = 10
// 	start := 0
// 	end := 0
// 	// 2D array to store the result of the word break
// 	dp := make([][]bool, len(s))
// 	for i := 0; i < len(s); i++ {
// 		dp[i] = make([]bool, len(s))
// 	}
// 	return isWordBreak(s, wordDict, start, end, &dp)
// }

// func isWordBreak(s string, wordDict []string, start, end int, dp *[][]bool) bool {
// 	if end == len(s) {
// 		return true
// 	}

// 	if (*dp)[start][end] {
// 		(*dp)[start][end] = true
// 		return (*dp)[start][end]
// 	}

// 	if isWordDictContains(wordDict, s[start:end+1]) && isWordBreak(s, wordDict, end+1, end+1, dp) {
// 		(*dp)[start][end] = true
// 		return true
// 	}

// 	if isWordBreak(s, wordDict, start, end+1, dp) {
// 		(*dp)[start][end] = true
// 		return true
// 	}

// 	(*dp)[start][end] = false
// 	return (*dp)[start][end]
// }

// func isWordDictContains(wordDict []string, word string) bool {
// 	for _, v := range wordDict {
// 		if v == word {
// 			return true
// 		}
// 	}
// 	return false

// }
