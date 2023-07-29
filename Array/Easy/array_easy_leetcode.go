package main

import (
	_ "data-structure/Array/input"
	"fmt"
	"math"
)

func main() {
	// var num = []int{23, 2, 6, 4, 7}
	// num[0] = 23
	// num[1] = 2
	// num[2] = 4
	// num[3] = 6
	// num[4] = 7
	// fmt.Println(checkSubarraySum(num, 6))
	// arr := []int{-1, 2, 3, 234, -56, 8, 989, 76}
	// FindTheLargestThreeDistinctElem(arr)
	// arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// fmt.Println("Given Array: ", arr)
	// fmt.Println(removeElementSolution2(arr, 2))
	// arr := []int{1, 2, 13, 8, 9, 74, 83, 15, 75}
	//fmt.Println(removeDuplicates(arr))
	// fmt.Println(FindTheSecondOrderElem(arr))
	// arr := []int{1, 2, 2, 3, 3, 4, 5, 5, 6}
	// fmt.Println(CheckArrayIfSortedInNonDescOrder(arr))
	// arr := []int{3, 4, 5, 1, 2}
	// fmt.Println(check(arr))
	// arr := []int{9}
	// fmt.Println(plusOne(arr))
	// arr := []int{1, 2, 3, 4, 1}
	// fmt.Println(containsDuplicate(arr))
	// nums1 := []int{1, 2, 3, 4, 0, 0, 0}
	// m := 4
	// nums2 := []int{2, 5, 6}
	// n := 3
	// mergeTwoSortedArray(nums1, m, nums2, n)
	arr1 := []int{1, -2147483648, 2}
	// arr2 := []int{2, 3, 3, 5, 6, 6, 7}
	// MergeSort(arr)
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	// sortColors(arr1)
	// fmt.Println()
	fmt.Println(thirdMax(arr1))
	//fmt.Println()
	// fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
}

// Continuous Subarray Sum
// Given an integer array nums and an integer k, return true if nums has a good subarray or false otherwise.

// A good subarray is a subarray where:

// its length is at least two, and
// the sum of the elements of the subarray is a multiple of k.
// Note that:

// A subarray is a contiguous part of the array.
// An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.

func checkSubarraySum(nums []int, k int) bool {
	sum := 0
	prev_idx := map[int]int{0: -1}

	for idx, num := range nums {
		sum += num
		sum = sum % k
		fmt.Println(sum, num)
		if val, ok := prev_idx[sum]; ok {
			fmt.Println(prev_idx, sum, idx, val)
			if (idx - val) >= 2 {
				return true
			}
		} else {
			prev_idx[sum] = idx
		}
	}

	return false
}

// Problem: Find the largest three distinct elements in an array
func FindTheLargestThreeDistinctElem(arr []int) {
	//create three elements and assign -∞
	var first int32 = math.MinInt32
	var second int32 = math.MinInt32
	var third int32 = math.MinInt32

	// iterate through all the elements

	for i := 0; i < len(arr); i++ {
		// check arr[i] > first
		// assign third = second
		// second = third
		// first = arr[i]
		if arr[i] > int(first) {
			third = second
			second = first
			first = int32(arr[i])
		} else if arr[i] > int(second) && arr[i] != int(first) {
			// check arr[i] > second && arr[i] != first
			// third = second
			// second = arr[i]
			third = second
			second = int32(arr[i])
		} else if arr[i] > int(third) && arr[i] != int(second) {
			// check arr[i] > third && arr[i] != second
			// third = second
			// second = arr[i]
			third = int32(arr[i])
		}
	}
	fmt.Println(first, second, third)
}

// Problem 27. Remove Element
// Solution: 1 - 2 ms
func removeElement(nums []int, val int) int {
	repeatCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], append(nums[i+1:], '_')...)
			repeatCount++
			i = i - 1
		}
	}
	return len(nums) - repeatCount
}

// Solution: 2 - 0ms
func removeElementSolution2(nums []int, val int) int {
	var i, j = 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	fmt.Println(nums)
	return i
}

// Leetcode Problem 26. Remove Duplicates from Sorted Array
// [0,0,1,1,1,2,2,3,3,4]
func removeDuplicates(nums []int) int {
	var j = 0
	for i := range nums {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	fmt.Println(nums)
	return j + 1
}

// Problem 40. Combination Sum II
// Input: candidates = [10,1,2,7,6,1,5], target = 8
/*
	Output:
	[
	[1,1,6],
	[1,2,5],
	[1,7],
	[2,6]
	]
*/
func combinationSum2(candidates []int, target int) [][]int {
	candidatesSumArray := [][]int{}
	return candidatesSumArray
}

// arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
func removeDuplicatesBruteForceSolution(nums []int) []int {
	i := 0
	j := 0
	for i < len(nums) && j < len(nums) {
		if nums[i] != nums[j] { // compare each element with others if not equal then increament i else i will be same
			nums[i+1] = nums[j] // if not equal swap i+1 to j
			i++
		}
		j++
	}
	return nums[:i+1]
}

// Fins second largest element from an array of integers
/**
* Brute force solution: sort the array and find the second largest
* Optimal: Lets have two variable firstLargest and seconLarget and assign min negative integers
* Iterate the loop
* - get first max elment and assign second largest to Flargest
* - get second max elment if second not equal to firstLargest
*
 */
// nums := []int{1,2, 13, 8, 9, 74, 83, 15, 75}
func FindTheSecondLargestElem(nums []int) int {
	fLargest := math.MinInt32
	sLargest := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if fLargest < nums[i] {
			sLargest = fLargest
			fLargest = nums[i]
		}
		if sLargest < nums[i] && fLargest > nums[i] && sLargest != fLargest {
			sLargest = nums[i]
		}
	}
	fmt.Println(fLargest, sLargest)
	return sLargest
}

// Fins second largest and smallest element from an array of integers
/**
* Brute force solution: sort the array and find the second largest
* Optimal: Lets have two variable firstLargest and seconLarget and assign min negative integers
* Iterate the loop
* - get first max elment and assign second largest to Flargest
* - get second max elment if second not equal to firstLargest
*
 */
// nums := []int{1,2, 13, 8, 9, 74, 83, 15, 75}
func FindTheSecondOrderElem(nums []int) []int {
	fLargest := math.MinInt32
	sLargest := math.MinInt32
	fSmallest := math.MaxInt32
	sSmallest := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if fLargest < nums[i] {
			sLargest = fLargest
			fLargest = nums[i]
		}
		if sLargest < nums[i] && fLargest > nums[i] && sLargest != fLargest {
			sLargest = nums[i]
		}

		// find the second smallest elem
		if fSmallest > nums[i] {
			sSmallest = fSmallest
			fSmallest = nums[i]
		}
		if sSmallest > nums[i] && fSmallest < nums[i] && sSmallest != fSmallest {
			sSmallest = nums[i]
		}
	}
	fmt.Println(fLargest, sLargest)
	return []int{sLargest, sSmallest}
}

// Check if a given array is sorted in non descending order
// Means: arr := [1,2,2,3,3,4,5,5] This should return in this case because all the elems are >= to each elems
// Return false in arr := [1,2,1,3,3,4]
func CheckArrayIfSortedInNonDescOrder(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

// 1752. Check if Array Is Sorted and Rotated
// arr := []int{3, 4, 5, 1, 2}
func check(nums []int) bool {
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		tmp := []int{}
		tmp = append(append(tmp, nums[i:]...), nums[:i]...)
		isSorted := true
		for j := 1; j < len(nums); j++ {
			if tmp[j] < tmp[j-1] {
				isSorted = false
				break
			}
		}
		if isSorted {
			return true
		}
	}
	return false
}

//You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.
// Problem : 66. Plus One - leetcode
func plusOne(digits []int) []int {
	for i := (len(digits) - 1); i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}

	tmpDigits := make([]int, len(digits)+1) // In case of 9 this one generate array of 2 elements with 0 default value
	tmpDigits[0] = 1
	return tmpDigits
}

// 217. Contains Duplicate - Leetcode
// arr := []int{1, 2, 3, 4, 1}
func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	storeNumbers := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := storeNumbers[nums[i]]; ok {
			return true
		}
		storeNumbers[nums[i]] += 1
	}
	return false
}

// Problem: 88. Merge Sorted Array
//
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	m := 3
//	nums2 := []int{2, 5, 6}
//	n := 3
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("nums1 Before: ", nums1)
	fmt.Println("nums2 Before: ", nums2)
	nums1 = append(nums1[:m], nums2...)
	fmt.Println(nums1)
	for i := 0; i < m; i++ {
		for j := len(nums1) - 1; j >= m; j-- {
			fmt.Println(i, j, nums1)
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	fmt.Println(nums1)
}

// Problem: 88. Merge Sorted Array
// MErge two sorted array with Gap analysis algorithm
func mergeTwoSortedArray(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("nums1 Before: ", nums1)
	fmt.Println("nums2 Before: ", nums2)
	nums1 = append(nums1[:m], nums2...)
	fmt.Println(nums1)
	//gap := int(math.Ceil(float64(m+n) / 2))
	// i, j := 0, gap

	// for gap > 1 {
	// 	fmt.Println(i, j, nums1[i], nums1[j])
	// 	if nums1[i] > nums1[j] {
	// 		nums1[i], nums1[j] = nums1[j], nums1[i]
	// 	}
	// 	j++
	// 	i++
	// 	if j > len(nums1)-1 && gap == 1 {
	// 		break
	// 	}
	// 	if j > len(nums1)-1 {
	// 		gap = int(math.Ceil(float64(gap) / 2))
	// 		j = gap
	// 		fmt.Println("J:", j)
	// 		i = 0
	// 	}

	// }

	for i := 0; i < m; i++ {
		for j := len(nums1) - 1; j >= m; j-- {
			fmt.Println(i, j, nums1)
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	fmt.Println(nums1)
}

/**
* Problem: 118. Pascal's Triangle
* Solution: First declare a 2-dimension array with N rows, then fill each row, both ends are 1 and the internal value is equal to the sum of both numbers on its shoulder (Dynamic Programming Algorithm).
* Input: numRows = 5
* Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 */

func generate(numRows int) [][]int {
	pasacelTriangleArray := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		pasacelTriangleArray[i] = make([]int, i+1)
		pasacelTriangleArray[i][0], pasacelTriangleArray[i][i] = 1, 1
		for j := 1; j < i; j++ {
			pasacelTriangleArray[i][j] = pasacelTriangleArray[i-1][j] + pasacelTriangleArray[i-1][j-1]
		}
	}
	return pasacelTriangleArray
}

/**
* Problem: 119. Pascal's Triangle II
* Solution: First declare a 2-dimension array with N rows, then fill each row, both ends are 1 and the internal value is equal to the sum of both numbers on its shoulder (Dynamic Programming Algorithm).
* Input: numRows = 5
* Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 */
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	pascelTriangleArr := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		pascelTriangleArr[i] = make([]int, i+1)
		pascelTriangleArr[i][0], pascelTriangleArr[i][i] = 1, 1
		for j := 1; j < i; j++ {
			pascelTriangleArr[i][j] = pascelTriangleArr[i-1][j] + pascelTriangleArr[i-1][j-1]
		}
	}
	return pascelTriangleArr[rowIndex]
}

/**
* Problem: 121. Best Time to Buy and Sell Stock
* Solution: .
* Input: prices = [7, 1, 5, 3, 6, 4, 1, 4, 5, 1, 2]
* Output: 8
 */
func maxProfit(prices []int) int {
	fmt.Println("Given Array: ", prices)
	minPrice := prices[0]
	maxPrice := math.MinInt32
	hasFoundMax := false
	for i := 1; i < len(prices); i++ {
		if minPrice >= prices[i] {
			minPrice = prices[i]
		}
		if prices[i] != minPrice {
			if maxPrice < prices[i]-minPrice {
				hasFoundMax = true
				maxPrice = prices[i] - minPrice
			}
		}
		fmt.Println(maxPrice, minPrice)
	}
	if !hasFoundMax {
		return 0
	}
	fmt.Println(maxPrice, minPrice)
	return maxPrice
}

/**
* Problem: 136. Single Number
* Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
* You must implement a solution with a linear runtime complexity and use only constant extra space.
* Input: nums = [4,1,2,1,2]
* Output: 4
 */

func singleNumber(nums []int) int {
	xor_result := 0

	for _, number := range nums {

		fmt.Println("Before change:", xor_result, number)
		xor_result ^= number
		fmt.Println("After change:", xor_result, number)
	}

	return xor_result
}

/**
* Problem: Longerst subarray with sum of k
* Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
* You must implement a solution with a linear runtime complexity and use only constant extra space.
* Input: nums = [4, 4, 1, 2, 1, 1, 1, 2], k = 3
* Output: 3
 */

func BruteForceLongestSubarrayofKSum(nums []int, k int) int {
	fmt.Println("Given Array: ", nums)
	totalLengthArr := len(nums)
	length := 0
	for i := 0; i < totalLengthArr; i++ {
		sum := 0
		for j := i; j < totalLengthArr; j++ {
			sum += nums[j]
			if sum == k && length < j-i+1 {
				length = j - i + 1
			}
		}
	}
	return length
}

// Solution: have a map and push the sum of the elem into it with i as value and sum as key
// check whether the sum exist in map then do not store
// check if the sum-k value is exist then increase length if length in less than sumOfMap[sum]-v then assign the length to sumOfMap[sum] - v
// Input: nums = [4, 4, 1, 2, 1, 1, 1, 2], k = 3
// Output: 3
func BetterSolutionLongestSubarrayofKSum(nums []int, k int) int {
	fmt.Println("Given Array: ", nums)
	sumOfMap := map[int]int{}
	totalLengthArr := len(nums)
	length := 0
	sum := 0
	for i := 0; i < totalLengthArr; i++ {
		sum += nums[i]
		if _, ok := sumOfMap[sum]; !ok {
			sumOfMap[sum] = i
		}
		if v, ok := sumOfMap[sum-k]; ok {
			if length < sumOfMap[sum]-v {
				fmt.Println(i)
				length = sumOfMap[sum] - v
			}
		}

	}
	fmt.Println(sumOfMap, sum)

	return length
}

/**
* Problem : Two Sum
* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
* Input: nums = [2,7,11,15], target = 9
* Output: [0,1]
 */
func twoSum(nums []int, target int) []int {
	fmt.Printf("Given Array: %d and target is: %d\n", nums, target)
	// numsMapStoreDiff := map[int]int{}
	// for i := 1; i < len(nums); i++ {
	// 	if _, ok := numsMapStoreDiff[target-nums[i-1]]; !ok {
	// 		numsMapStoreDiff[target-nums[i-1]] = i - 1
	// 	}
	// 	if v, ok := numsMapStoreDiff[nums[i]]; ok {
	// 		return []int{v, i}
	// 	}
	// }
	// return []int{0, 0}
	diffMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := diffMap[target-nums[i]]; !ok {
			diffMap[target-nums[i]] = i
		}
		fmt.Println(diffMap)
		if Val, ok := diffMap[nums[i+1]]; ok {
			return []int{Val, i + 1}
		}
	}

	return []int{0, 0}
}

/**
* BruteForce Solution
* Problem: 1752 - Check if Array Is Sorted and Rotated
* Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.There may be duplicates in the original array.
* Input: nums = [3,4,5,1,2]
* Output: true
* Explanation: [1,2,3,4,5] is the original sorted array.
* You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].
 */

func BruteForceCheckArraySortedRotated(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		tmp := append(nums[i:], nums[:i]...)
		isSortedRotated := true
		for j := 1; j < len(nums); j++ {
			if tmp[j-1] > tmp[j] {
				isSortedRotated = false
			}
		}
		if isSortedRotated {
			return true
		}
	}
	return false
}

/**
* Optimal Solution
* Problem: 1752 - Check if Array Is Sorted and Rotated
* Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.There may be duplicates in the original array.
* Input: nums = [3,4,5,1,2]
* Output: true
* Explanation: [1,2,3,4,5] is the original sorted array.
* You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].
 */

func checkArraySortedRotated(nums []int) bool {
	fmt.Println("Given Array: ", nums)
	// for i := 1; i < len(nums); i++ {
	// 	j := 0
	// 	// tmp := append(nums[i:], nums[:i]...)
	// 	isSortedRotated := true
	// 	// // for j := 1; j < len(nums); j++ {
	// 	// 	if tmp[j-1] > tmp[j] {
	// 	// 		isSortedRotated = false
	// 	// 	}
	// 	// }
	// 	// if isSortedRotated {
	// 	// 	return true
	// 	// }

	// 	if nums[i] < nums[i-1] {
	// 		isSortedRotated = false
	// 	}
	// 	for j < i {

	// 		if nums[j] < nums[i] && nums[j] < nums[j+1] {
	// 			fmt.Println(j, i, nums[j], nums[i], nums[j], nums[j+1])
	// 			isSortedRotated = false
	// 		}
	// 		j++
	// 	}
	// 	if isSortedRotated {
	// 		return true
	// 	}
	// }
	// return false
	// decreased := 0

	// for i, num := range nums {
	//     if num > nums[(i + 1) % len(nums)] {
	//         decreased++
	//     }

	//     if decreased > 1 {
	//         return false
	//     }
	// }

	// return true
	decreament := 0
	for i, num := range nums {
		fmt.Println(num, nums[(i+1)%len(nums)], (i+1)%len(nums))
		if num > nums[(i+1)%len(nums)] {
			decreament++
		}
		if decreament > 1 {
			return false
		}
	}
	return true
}

/**
* Problem: 283. Move Zeroes
* Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
* Note that you must do this in-place without making a copy of the array.
 */
func moveZeroes(nums []int) {
	fmt.Println("Given Array: ", nums)
	foundNonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[foundNonZero] = nums[i]
			foundNonZero++
		}
	}

	for i := foundNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums, foundNonZero)
}

/**
* Problem: 283. Move Zeroes
* Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
* Note that you must do this in-place without making a copy of the array.
* Solution: Two points approach check if the number is non-zero then swap the element and increament i and keep the j pointer at 0 number
 */
func moveZeroesOptimal(nums []int) {
	var i int
	var j = -1
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
			break
		}
	}
	if j < 0 {
		return // if no zero found returns here
	}

	for i = j + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums, i, j)
}

/**
* Union of two sorted array
* You are given a two sorted arrays and you have find the union of the those two arrays
* Example: arr1 = [1,2,2,3,4,5], arr2 = [3,4,5,6]
* Output: [1,2,3,4,5,6]
 */

func unionOfTwoSortedArrays(arr1 []int, arr2 []int) {
	var i, j int
	tmp := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			if !findInSlice(tmp, arr1[i]) {
				tmp = append(tmp, arr1[i])
			}
			i++
		} else {
			if !findInSlice(tmp, arr2[j]) {
				tmp = append(tmp, arr2[j])
			}
			j++
		}
	}
	for i < len(arr1) {
		tmp = append(tmp, arr1[i])
		i++
	}
	for j < len(arr2) {
		tmp = append(tmp, arr2[j])
		j++
	}

	fmt.Println(tmp)
}

/**
* Intersection of two sorted array
* You are given a two sorted arrays and you have find the union of the those two arrays
* Example: arr1 = [1,2,2,3,4,5], arr2 = [3,3,4,5,6]
* Output: [3,4,5]
 */
func intersectionOfTwoSortedArraysBruteForce(arr1 []int, arr2 []int) {
	intersectionArray := []int{}
	visitedArray := make([]int, len(arr2))
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			fmt.Println(arr1[i], arr2[j])
			if arr1[i] == arr2[j] && visitedArray[j] == 0 {
				intersectionArray = append(intersectionArray, arr1[i])
				visitedArray[j]++
				break
			}
			if arr2[j] > arr1[i] {
				break
			}
		}
	}
	fmt.Println(intersectionArray)
}

/**
* Intersection of two sorted array
* You are given a two sorted arrays and you have find the union of the those two arrays
* Example: arr1 = [1,2,2,3,4,5], arr2 = [3,3,4,5,6]
* Output: [3,4,5]
* Solution: Two pointer approach
 */
func intersectionOfTwoSortedArraysOptimal(arr1 []int, arr2 []int) {
	intersectionArray := []int{}
	var i, j int
	for i < len(arr1) && j < len(arr2) {

		if arr2[j] > arr1[i] {
			i++
		}

		if arr1[i] == arr2[j] {
			intersectionArray = append(intersectionArray, arr1[i])
			i++
			j++
		}
	}
	fmt.Println(intersectionArray)
}

func findInSlice(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

/**
* Problem: 268. Missing Number
* Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
* Input: nums = [3,0,1]
* Output: 2
* Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
 */
func missingNumberBruteForce(nums []int) int {
	maxNumber := math.MinInt32
	minNumber := math.MaxInt16

	// for i := 0; i < len(nums); i++ {
	// 	if maxNumber < nums[i] {
	// 		maxNumber = nums[i]
	// 	}
	// 	if minNumber > nums[i] {
	// 		minNumber = nums[i]
	// 	}
	// }

	// for i := minNumber; i < maxNumber; i++ {
	// 	if !findInSlice(nums, i) {
	// 		return i
	// 	}
	// }
	lenArr := len(nums)
	for i := 0; i <= lenArr; i++ {
		flag := 0
		for j := 0; j < lenArr; j++ {
			if nums[j] == i {
				flag = 1
				break
			}

		}
		if flag == 0 {
			return i
		}

	}
	fmt.Println(maxNumber, minNumber)
	return 0
}

func missingNumberBetter(nums []int) int {
	visitedNumbersArr := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		visitedNumbersArr[nums[i]]++
	}
	for i := 0; i < len(visitedNumbersArr); i++ {
		if visitedNumbersArr[i] == 0 {
			return i
		}
	}
	fmt.Println(visitedNumbersArr)
	return 0
}

// Solution: using natural numbers sum formula
// This can be achieved using bitwise operator
func missingNumberOptimal(nums []int) int {
	sum := (len(nums) * (len(nums) + 1)) / 2
	sumOfNumbers := 0
	for i := 0; i < len(nums); i++ {
		sumOfNumbers += nums[i]
	}
	return sum - sumOfNumbers
}

// This can be achieved using bitwise operator
func missingNumberOptimalUsingXor(nums []int) int {
	xor1 := 0
	xor2 := 0
	for i := 0; i < len(nums); i++ {
		xor2 = xor2 ^ nums[i]
		xor1 = xor1 ^ (i + 1)
	}
	return xor1 ^ xor2
}

/**
* Problem: 747. Largest Number At Least Twice of Others
* You are given an integer array nums where the largest integer is unique.
* Determine whether the largest element in the array is at least twice as much as every other number in the array. If it is, return the index of the largest element, or return -1 otherwise.
 */

func dominantIndex(nums []int) int {
	fmt.Println("GIven Array: ", nums)
	largestElem := nums[0]
	//largetTwice := math.MinInt32
	j := 0
	for i := 0; i < len(nums); i++ {
		// if nums[i] > largestElem {

		// }
		// if largetTwice < nums[i]*2 {
		// 	largetTwice = nums[i] * 2
		// }
		//fmt.Println(nums[i], largestElem, largetTwice)
		if nums[i] > largestElem {
			fmt.Println("Largest element twice:", largestElem*2, largestElem, nums[i])
			if largestElem*2 <= nums[i] {
				largestElem = nums[i]
				j++
			}

		}
		// if largestElem < nums[i]*2 {
		// 	j--
		// }
	}
	if j > 0 {
		return j
	}
	return -1
}

func findMaxSubarryOfSumK(nums []int, k int) []int {
	sumOfNUmbers := 0
	i, j := 0, 0
	for i < len(nums) && j < len(nums) {
		sumOfNUmbers += nums[j]
		if sumOfNUmbers <= k {
			j++
		}

		if sumOfNUmbers == k {
			sumOfNUmbers = 0
			i++
		}
		if sumOfNUmbers > k {
			sumOfNUmbers = 0
			i++
			j = i
		}
	}
	fmt.Println(i, j, sumOfNUmbers)
	return nums
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	res = append(res, []int{})

	var nums1Set, nums2Set [1001][2]int

	for _, num := range nums1 {
		if num >= 0 {
			nums1Set[num][0]++
		} else {
			nums1Set[-num][1]++
		}
	}

	for _, num := range nums2 {
		if num >= 0 {
			nums2Set[num][0]++
		} else {
			nums2Set[-num][1]++
		}
	}

	for i := -1000; i < 1001; i++ {
		if i >= 0 {
			if nums1Set[i][0] > 0 && nums2Set[i][0] == 0 {
				res[0] = append(res[0], i)
			}
		} else {
			if nums1Set[-i][1] > 0 && nums2Set[-i][1] == 0 {
				res[0] = append(res[0], i)
			}
		}
	}

	for i := -1000; i < 1001; i++ {
		if i >= 0 {
			if nums2Set[i][0] > 0 && nums1Set[i][0] == 0 {
				res[1] = append(res[1], i)
			}
		} else {
			if nums2Set[-i][1] > 0 && nums1Set[-i][1] == 0 {
				res[1] = append(res[1], i)
			}
		}
	}

	return res
}

/**
 	Problem: Given a binary array nums, return the maximum number of consecutive 1's in the array.
 	Input: nums = [1,1,0,1,1,1]
	Output: 3
	Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
*/

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter++
		} else {
			counter = 0
		}
		if counter > maxCount {
			maxCount = counter
		}
	}
	return maxCount
}

/**
* Problem: Longest Subarray with given Sum K
* k = 5, array[] = {2,3,5}
* Explanation: The longest subarray with sum 5 is {2, 3}. And its length is 2.
* Solution: 2 pointer approach
 */

func findLongestSubarrayWithSumK(nums []int, k int) int {
	fmt.Println("Given Array: ", nums)
	var i, j, counter, sum = 0, 0, 0, 0
	for i < len(nums) && j < len(nums) {
		sum += nums[j]
		if sum < k {
			j++
		}

		if sum >= k {
			if sum == k {
				if counter < j-i {
					counter = (j - i) + 1
				}
				fmt.Println("I am here", i, j, counter)
			}
			i++
			j = i
			sum = 0
		}
	}
	return counter
}

/**
* Problem: 75. Sort Colors
* Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
* We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
* You must solve this problem without using the library's sort function.
* Input: nums = [2,0,2,1,1,0]
* Output: [0,0,1,1,2,2]
* Two Poiters Opposite direction approch
 */

func sortColors(nums []int) {
	var i, j = 0, len(nums) - 1
	for i < j {
		fmt.Println("I am here", nums[i], nums[j])
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		j--
		if i == j {
			i++
			if nums[i-1] == nums[i] {
				i++
			}
			j = len(nums) - 1
		}

	}
	fmt.Println(nums)
}

/**
* Problem: 169. Majority Element
* Given an array nums of size n, return the majority element.
* The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
* Input: nums = [3,2,3]
* Output: 3
 */

func majorityElementBruteForce(nums []int) int {
	majorityCounter := 0
	var majorityElem int
	for i := 0; i < len(nums); i++ {
		tmpCounter := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				tmpCounter++
			}
		}
		if majorityCounter < tmpCounter {
			majorityElem = nums[i]
			majorityCounter = tmpCounter
		}
	}

	if majorityCounter > int(len(nums)/2) {
		return majorityElem
	}
	return 0
}

func majorityElementBetterSolution(nums []int) int {
	majorityElemCountMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := majorityElemCountMap[nums[i]]; !ok {
			majorityElemCountMap[nums[i]] = 0
		}
		majorityElemCountMap[nums[i]]++
		fmt.Println(int(float64(len(nums) / 2)))
		if majorityElemCountMap[nums[i]] > int(float64(len(nums)/2)) {
			return nums[i]
		}
	}

	return 0
}

/**
* Moores Voting algorithm
 */
func majorityElementOptimalSolution(nums []int) int {
	// var el, counter, arrLength, counter2 = 0, 0, len(nums), 0
	// for i := 0; i < arrLength; i++ {
	// 	if counter == 0 {
	// 		counter = 1
	// 		el = nums[i]
	// 	} else if nums[i] == el {
	// 		counter++
	// 	} else {
	// 		counter--
	// 	}
	// }
	// for i := 0; i < arrLength; i++ {
	// 	if el == nums[i] {
	// 		counter2++
	// 	}
	// }
	// if counter2 > int(arrLength/2) {
	// 	return el
	// }
	// return 0
	leader := nums[0]
	appearances := 0

	for i := 0; i < len(nums); i++ {
		if appearances == 0 {
			leader = nums[i]
		}

		if nums[i] == leader {
			appearances++
		} else {
			appearances--
		}
	}

	return leader
}

/**
* Problem: 121. Best Time to Buy and Sell Stock
* You are given an array prices where prices[i] is the price of a given stock on the ith day.
* Input: prices = [7,1,5,3,6,4]
* Output: 5
* Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
* Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
 */
func maxProfitSolution(nums []int) int {
	var minPrice, maxProfit, arrLength = nums[0], 0, len(nums)
	for i := 1; i < arrLength; i++ {
		cost := nums[i] - minPrice
		maxProfit = int(math.Max(float64(maxProfit), float64(cost)))
		minPrice = int(math.Min(float64(minPrice), float64(nums[i])))
	}
	return maxProfit
}

/**
* Problem: 414. Third Maximum Number
* Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.
* Input: nums = [3,2,1]
* Output: 1
 */

func thirdMax(nums []int) int {
	firstMax := math.MinInt64
	secondMax := math.MinInt64
	thirdMax := math.MinInt64
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = nums[i]
		}
		if nums[i] > secondMax && nums[i] < firstMax {
			thirdMax = secondMax
			secondMax = nums[i]
		}

		if nums[i] > thirdMax && nums[i] < secondMax {
			thirdMax = nums[i]
		}
	}
	if thirdMax == math.MinInt64 {
		return firstMax
	}

	return thirdMax
}
