package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// nums := []int{1, 2, 4, 8, 9, -2, -7, 3}
	// nums1 := []int{10, 1, 2, 7, 6, 1, 5}
	// nums1 := []int{1, 2, 3, 5, 6, 7}
	// // n := 8
	// // m := 3
	// // nums2 := []int{2, 5, 6}
	// // n := 3
	// // ans := []int{}

	// // fmt.Println(permuteUnique(nums1))
	// fmt.Println(BinarySearchRecursive(nums1, 5, 0, len(nums1)-1))

	// matrix := [][]byte{
	// 	{5, 3, '.', '.', 7, '.', '.', '.', '.'},
	// 	{6, '.', '.', 1, 9, 5, '.', '.', '.'},
	// 	{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
	// 	{8, '.', '.', '.', 6, '.' ,'.', '.', 3},
	// 	{4, '.', '.', 8, '.', 3, '.', '.', 1},
	// 	{7, '.', '.', '.', 2, '.', '.', '.', 6},
	// 	{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
	// 	{'.', '.', '.', 4, 1, 9, '.', '.', 5},
	// 	{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	// }
	// fmt.Println(isValidSudoku(matrix))

	// box := [][]string{
	// 	{"#", "#", "*", ".", "*", "."},
	// 	{"#", "#", "#", "*", ".", "."},
	// 	{"#", "#", "#", ".", "#", "."},
	// }
	box := [][]byte{{'#', '.', '#'}}
	fmt.Println(rotateTheBoxs(box))
}

/** Print the all subsequences whose sum equal to k
* Using recusrion we can check use Pick or No pick approach. First function will be called to pick the element and add into sum.
* Second function call happens to not picking the value and subtract them from the sum.
 */

func printKsumSubsequences(nums []int, arr []int, i, sum, k int) {
	if i >= len(nums) {
		if sum == k {
			fmt.Println(arr)
		}
		return
	}

	arr = append(arr, nums[i])
	sum += nums[i]
	printKsumSubsequences(nums, arr, i+1, sum, k) // Pick element to add in sum

	arr = append([]int{}, arr[:len(arr)-1]...) //Remove the element from an array as we are not going pick this for subsequences.
	sum -= nums[i]
	printKsumSubsequences(nums, arr, i+1, sum, k) // Not pick the element

}

/** Print only one subsequences whose sum equal to k
* Print the array else return false
* Using recusrion we can check use Pick or No pick approach. First function will be called to pick the element and add into sum.
* Second function call happens to not picking the value and subtract them from the sum.
 */

func printKSumOneSubsequences(nums, arr []int, index, sum, k int) bool {
	// Base condition
	if index >= len(nums) {
		if sum == k {
			fmt.Println(arr)
			return true
		}
		return false
	}

	arr = append(arr, nums[index])
	sum += nums[index]

	if printKSumOneSubsequences(nums, arr, index+1, sum, k) {
		return true
	}

	arr = append([]int{}, arr[:len(arr)-1]...)
	sum -= nums[index]

	return printKSumOneSubsequences(nums, arr, index+1, sum, k)
}

/** Print count of subarray subsequences whose sum equal to k
* Print the count
* Using recusrion we can check use Pick or No pick approach. First function will be called to pick the element and add into sum.
* Second function call happens to not picking the value and subtract them from the sum.
 */

func printSubarrayCountSumEqualk(nums []int, arr []int, index, sum, k int) int {
	// Base condition
	if index >= len(nums) {
		if sum == k {
			fmt.Println(arr)
			return 1
		}
		return 0
	}
	arr = append(arr, nums[index])
	sum += nums[index]

	left := printSubarrayCountSumEqualk(nums, arr, index+1, sum, k)

	arr = append([]int{}, arr[:len(arr)-1]...)
	sum -= nums[index]

	right := printSubarrayCountSumEqualk(nums, arr, index+1, sum, k)

	return left + right
}

/** Merge Sort
* TC : O(nlogn), SC: O(n)
 */

func mergeSort(arr []int, low, high int) {

	if low == high {
		return
	}

	mid := (low + high) / 2

	mergeSort(arr, low, mid)    // first recursive call this will finish first and below code execute later
	mergeSort(arr, mid+1, high) // Second recirsive function call which call alway first recursive calls

	merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {

	left := low
	right := mid + 1
	tmp := []int{}
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			tmp = append(tmp, arr[left])
			left++
		} else {
			tmp = append(tmp, arr[right])
			right++
		}
	}

	for left <= mid {
		tmp = append(tmp, arr[left])
		left++
	}

	for right <= high {
		tmp = append(tmp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = tmp[i-low]
	}
	fmt.Println("Here I am", arr, tmp, low, high)
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
	mid := (m + n) / 2
	if (m+n)%2 == 0 {
		mid--
	}
	fmt.Println("Print merge inputs: ", nums1, 0, mid, (m+n)-1)
	merge(nums1, 0, mid, (m+n)-1)
	fmt.Println(nums1)
}

/**
* Quick Sort Algorithm
* TC(nlogn), SC(1)
* First we pick the pivot element and put this into its correct position
* Then do the recursive calls for left and right partition of the array untill we sort the array.
 */

func QuickSort(arr []int, low, high int) {
	//lets add the base condition
	if low < high {
		// find the pivot element from the array
		pivot := partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	i, j := low, high
	pivot := arr[low]
	for i < j {
		for arr[i] <= pivot && i <= high {
			// if the elements in lesser than pivot then move the i pointer towards right
			i++
		}
		for arr[j] > pivot && j >= low {
			// If the elements is greater then pivot then move the j pointer towards left
			j--
		}

		if i < j {
			// swap the elements for the above the i, j values
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// finally swap the pivot element to its correct order in the array
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

/**
* Combination of Sum
* Find the combination of given array elements which sum equal to the given sum
* You can repeat the element any number of times.
* Input: [2,3,6,7], target = 7
* Output: [[2,2,3], [7]]
 */

func combinationOfSum(arr []int, target int) [][]int {
	ans := [][]int{}
	findCombination(0, arr, target, &ans, []int{})
	return ans
}

func findCombination(index int, arr []int, target int, ans *[][]int, result []int) {
	if index == len(arr) {
		if target == 0 {
			*ans = append(*ans, append([]int{}, result...))
		}
		return
	}

	if arr[index] <= target {
		result = append(result, arr[index])
		findCombination(index, arr, target-arr[index], ans, result)
		result = result[:len(result)-1]
	}

	findCombination(index+1, arr, target, ans, result)
}

/**
* Combination of Sum
* Find the first combination of given array elements which sum equal to the given sum
* You can repeat the element any number of times.
* Input: [2,3,6,7], target = 7
* Output: [[2,2,3], [7]]
 */

func combinationOfSums(arr []int, target int) [][]int {
	ans := [][]int{}
	stopExec := false
	findCombinations(0, arr, target, &ans, []int{}, &stopExec)
	return ans
}

func findCombinations(index int, arr []int, target int, ans *[][]int, result []int, stopExec *bool) {
	if index == len(arr) {
		if target == 0 {
			*ans = append(*ans, result)
			*stopExec = true
		}
		return
	}

	if *stopExec {
		return
	}

	if arr[index] <= target {
		result = append(result, arr[index])
		findCombinations(index, arr, target-arr[index], ans, result, stopExec)
		result = result[:len(result)-1]
	}

	findCombinations(index+1, arr, target, ans, result, stopExec)
}

/**
* Combination of Sum ||
* Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
* Each number in candidates may only be used once in the combination.
* Note: The solution set must not contain duplicate combinations.
* Input: candidates = [10,1,2,7,6,1,5], target = 8
* Output: [[1,1,6],[1,2,5],[1,7],[2,6]]
 */

func combinationofSum2(candidates []int, target int) [][]int {
	fmt.Println("Given Array: ", candidates, target)
	ans := [][]int{}
	sort.Ints(candidates)
	findCombinationsum2(candidates, 0, target, []int{}, &ans)
	return ans
}

func findCombinationsum2(candidates []int, index, target int, ds []int, ans *[][]int) {

	if target == 0 {
		*ans = append(*ans, append([]int{}, ds...))
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		if candidates[i] > target {
			break
		}

		ds = append(ds, candidates[i])
		// fmt.Println(i, target)
		findCombinationsum2(candidates, i+1, target-candidates[i], ds, ans)
		ds = ds[:len(ds)-1]
	}
}

/**
* Subset of sum I
* Print all the possible sum of subset in a given array.
* input= {3,1,2}, Output = {6,4,5,3,3,1,2,0}
 */

func PrintOfSubsetSum(arr []int, index int, sum int, ans *[]int) {
	if index >= len(arr) {
		*ans = append(*ans, sum)
		return
	}

	sum += arr[index]
	PrintOfSubsetSum(arr, index+1, sum, ans)
	sum -= arr[index]

	PrintOfSubsetSum(arr, index+1, sum, ans)
}

/**
* Problem: 90. Subsets II
* Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
* The solution set must not contain duplicate subsets. Return the solution in any order.
* Input: nums = [1,2,2]
* Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
 */

func subsetsWithDup(num []int) [][]int {
	sort.Ints(num)
	ans := [][]int{}
	subsetsWithDupCalls(num, 0, []int{}, &ans)
	return ans
}

func subsetsWithDupCalls(num []int, index int, ds []int, ans *[][]int) {
	// if index > len(num) {
	// 	*ans = append(*ans, append([]int{}, ds...))
	// 	// return
	// }

	// if index == 0 || index == 1 {
	// 	*ans = append(*ans, append([]int{}, ds...))
	// }

	*ans = append(*ans, append([]int{}, ds...))

	for i := index; i < len(num); i++ {
		if i != index && num[i] == num[i-1] {
			continue
		}
		ds = append(ds, num[i])
		subsetsWithDupCalls(num, i+1, ds, ans)
		ds = ds[:len(ds)-1]
	}
}

/**
* Leetcode Problem 46: Permutations
* Given an array nums of distinct integers, return all the possible permutations. You can return array in any order.
* Example:
* Input: nums = [1,2,3]
* Output: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
* Solution: Recursive approach & Use map to keep track the picked index
 */

func permute(nums []int) [][]int {
	ans := [][]int{}
	trackDS := make(map[int]bool, 3)
	Permutations(nums, []int{}, &ans, trackDS)
	return ans
}

func Permutations(nums []int, ds []int, ans *[][]int, trackDS map[int]bool) {
	if len(ds) == len(nums) {
		*ans = append(*ans, append([]int{}, ds...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if trackDS[i] {
			continue
		}
		ds = append(ds, nums[i])
		trackDS[i] = true
		Permutations(nums, ds, ans, trackDS)
		ds = ds[:len(ds)-1]
		trackDS[i] = false
		fmt.Println("ith value", i)
	}
}

/**
* Leetcode Problem 46: Permutations
* Given an array nums of distinct integers, return all the possible permutations. You can return array in any order.
* Example:
* Input: nums = [1,2,3]
* Output: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
* Solution: Recursive approach & not going to use the extra map to keep record of array elements
* Instead will use the swapping approach to form permutations
 */

func permuteBetter(nums []int) [][]int {
	ans := [][]int{}
	// trackDS := make(map[int]bool, 3)
	PermutationsBetter(nums, 0, &ans)
	return ans
}

func PermutationsBetter(nums []int, index int, ans *[][]int) {
	if index >= len(nums) {
		*ans = append(*ans, append([]int{}, nums...))
		return
	}

	for i := index; i < len(nums); i++ {

		nums[index], nums[i] = nums[i], nums[index]
		PermutationsBetter(nums, index+1, ans)
		nums[index], nums[i] = nums[i], nums[index] // when you come back from the recursion we need to reswap it to make its original array
	}
}

/**
* Leetcode Problem : 47. Permutations II
* Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
* Input: nums = [1,1,2]
* Output:[[1,1,2],[1,2,1],[2,1,1]]
* Approach: to handle duplicate elements properly, we need to ensure that we're not skipping permutations that are actually distinct.
* In the case of duplicate elements, simply skipping when the current element is equal to the previous one may not be sufficient,
* as it could lead to skipping permutations that are valid.
 */

func permuteUnique(nums []int) [][]int {
	// sort.Ints(nums)
	ans := [][]int{}
	used := make(map[int]bool)
	track := []int{}
	// trackDS := make(map[int]bool, 3)
	PermutationsUnique(nums, &ans, used, track)
	fmt.Println(used, len(ans))
	return ans
}

func PermutationsUnique(nums []int, ans *[][]int, used map[int]bool, track []int) {
	if len(track) == len(nums) {
		*ans = append(*ans, append([]int{}, track...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		PermutationsUnique(nums, ans, used, track)
		track = track[:len(track)-1]
		used[i] = false
	}
}

/**
* Problem: 50. Pow(x, n)
* Implement pow(x, n), which calculates x raised to the power n (i.e., x^n).
* Input: x = 2.00000, n = 10
* Output: 1024.00000
 */

func myPow(x float64, n int) float64 {
	var powerSum float64
	var isNegative bool
	if n < 0 {
		isNegative = true
		n = int(math.Abs(float64(n)))
	}
	if n == 0 {
		return 1
	}
	for i := 0; i < n; i++ {
		if powerSum > 0 {
			powerSum *= x
		} else if powerSum < 0 {
			powerSum *= x
		} else {
			powerSum = x
		}

	}

	if isNegative {
		powerSum = 1 / powerSum
	}

	return powerSum
}

func myPowOptimalSolution(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	half := myPowOptimalSolution(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// Binary Search Using Recurssion

func BinarySearch(nums []int, n, startInd, EndInd int) bool {
	// length := len(nums)
	fmt.Println("Given Array", nums, startInd, EndInd)
	half := (startInd + EndInd) / 2
	fmt.Println(half, len(nums), EndInd)
	if half == 0 {
		return false
	}
	if nums[half] > n {
		EndInd = half
	} else {
		startInd = half
	}
	if startInd >= len(nums) || startInd >= EndInd {
		return false
	}

	if nums[startInd] == n || nums[EndInd] == n {
		return true
	} else {
		startInd++
		EndInd--
	}

	// if nums[EndInd] < n {
	// 	return false
	// }

	// if startInd+1 == EndInd {
	// 	return false
	// }

	return BinarySearch(nums, n, startInd, EndInd)
}

// binarySearchRecursive searches for the target value in the sorted slice arr
// recursively. It returns the index of the target in the slice if found,
// otherwise returns -1.
func BinarySearchRecursive(nums []int, target, startIndex, endIndex int) int {
	if startIndex > endIndex {
		return -1
	}

	mid := (startIndex + endIndex) / 2

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		startIndex = mid + 1
		return BinarySearchRecursive(nums, target, startIndex, endIndex)
	} else {
		endIndex = mid - 1
		return BinarySearchRecursive(nums, target, startIndex, endIndex)
	}
}

/** Problem: 36. Valid Sudoku
* Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
* Each row must contain the digits 1-9 without repetition.
* Each column must contain the digits 1-9 without repetition.
* Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
* Note:
* A Sudoku board (partially filled) could be valid but is not necessarily solvable.
* Only the filled cells need to be validated according to the mentioned rules.
 */

func isValidSudoku(board [][]byte) bool {
	return validSudoku(&board, 0, 0)
}

func validSudoku(board *[][]byte, rowStart int, colStart int) bool {
	for row := rowStart; row < len(*board); row++ {
		for col := colStart; col < len((*board)[0]); col++ {
			if (*board)[row][col] == '.' {
				continue
			}
			// fmt.Println("Values", (*board)[row][col], row, col)
			if !isValid(board, row, col, (*board)[row][col]) {
				return false
			} else {
				return validSudoku(board, row, col+1)
			}
		}
		colStart = 0
	}
	return true
}

func isValid(board *[][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {

		if (*board)[i][col] == val && i != row {
			return false
		}

		if (*board)[row][i] == val && i != col {
			return false
		}

		iRow := (3 * (row / 3)) + (i)/3
		iCol := (3 * (col / 3)) + (i)%3

		rowColStatus := iRow != row && iCol != col
		if (*board)[iRow][iCol] == val && rowColStatus {
			return false
		}
	}
	return true
}

func rotateTheBox(box [][]string) [][]string {
	// box := [][]byte{
	// 	{'#','#','*','.','*','.'},
	// 	{'#','#','#','*','.','.'},
	// 	{'#','#','#','.','#','.'},
	// }
	ansBox := make([][]string, len(box[0]))
	for row := 0; row < len(box); row++ {
		for col := 0; col < len(box[0]); col++ {
			if row == 0 {
				ansBox[col] = make([]string, len(box))
			}
			// fmt.Println("Ans row and col", col, len(box)-1-row, box[row][col])
			// ansBox[col][len(box)-1-row] = box[row][col]
			// if box[row][col] == "*" {
			// 	ansBox[col][len(box)-1-row] = box[row][col]
			// 	// breakCol = col
			// 	continue
			// }
			if box[row][col] == "." {
				// fmt.Println("Row and Col", row, col)
				for k := col; k >= 1; k-- {
					if box[row][k-1] == "*" {
						break
					}
					box[row][k], box[row][k-1] = box[row][k-1], box[row][k]
					// ansBox[row][k], ansBox[row][k -1] = ansBox[row][k -1], ansBox[row][k]
					// ansBox[col][len(box)-1-row] = box[row][k]
				}

			}
		}
	}

	for row := 0; row < len(ansBox); row++ {
		for col := 0; col < len(ansBox[0]); col++ {
			ansBox[row][len(ansBox[0])-1-col] = box[col][row]
		}
	}

	fmt.Println(box)
	// fmt.Println(ansBox)

	return ansBox
}

func rotateTheBoxs(box [][]byte) [][]byte {
	// Get dimensions of the box
	rows := len(box) //
	cols := len(box[0])

	// Create the rotated box with dimensions swapped
	rotatedBox := make([][]byte, cols)
	for col := range rotatedBox {
		rotatedBox[col] = make([]byte, rows)
	}

	// Rotate the box
	for row := 0; row < rows; row++ { // 0
		emptyCol := cols - 1 // Track the empty column position //

		// Move stones to the rightmost empty positions
		for col := cols - 1; col >= 0; col-- {
			rotatedBox[col][row] = '.'
			if box[row][col] == '#' {
				rotatedBox[emptyCol][rows-1-row] = '#'
				emptyCol--
			} else if box[row][col] == '*' {
				emptyCol = col - 1                // Reset empty column position
				rotatedBox[col][rows-1-row] = '*' // Place obstacle directly
			}
		}
	}

	return rotatedBox
}
