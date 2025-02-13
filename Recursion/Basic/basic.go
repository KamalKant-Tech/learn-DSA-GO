package main

import (
	"fmt"
)

func main() {
	//str := "kamal"
	// nums := []int{3, 1, 2}
	// arr := []int{}
	// printSubsequences(0, nums, arr)
	matrix := [][]int{
		{1,2,3},
		{3,1,2},
		{2,3,1},
	}
	fmt.Println(checkValid(matrix))
	// matrix := [][]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }
	// fmt.Println(diagonalSumBetter(matrix))
	// matrix := [][]int{
	// 	{1,2,3,4},
	// 	{5,6,7,8},
	// 	{9,10,11,12},
	// 	{13,14,15,16},
	// }
	// printMatrixSpiralOrder(matrix)
}

/**
* Print name N times using recursion
* TC -> O(n), SC -> O(n)
 */

func printName(n int, count int) {
	if count > n {
		return
	}
	fmt.Println("John Doe")
	printName(n, count+1)
}

/**
* Print numbets 1 -> N using recursion
* TC -> O(n), SC -> O(n)
 */

func printNumbers(n int, count int) {
	if count > n {
		return
	}
	fmt.Println(count)
	printNumbers(n, count+1)
}

/**
* Print Numbers 1 -> N using recursion but you are allowed to do the increment in numbers
* TC -> O(n), SC -> O(n)
* Approch: We are going to print numbers after the function call.
* So, when func invoke firt time it should not print number immediataly rather it should wait for the another function call.
 */

func printNumberAnotherSolution(i, n int) {
	if i < 1 {
		return
	}

	printNumberAnotherSolution(i-1, n)
	fmt.Println(i)
}

/**
* Sum of first N numbers
* Approach: Parametrised way
 */

func SumOfNNumbers(n, sum int) {
	if n < 1 {
		fmt.Println(sum)
		return
	}

	SumOfNNumbers(n-1, sum+n)
}

/**
* Sum of first N numbers
* Approach: Functional way
 */

func SumOfNNumbersSolution(n int) int {
	if n == 0 {
		return 0
	}

	return n + SumOfNNumbersSolution(n-1)
}

/** Reverse an array using recursion
* Approach: Two Pointer
 */

func reverseArray(nums []int, start, end int) {
	if start >= end {
		return
	}
	nums[start], nums[end] = nums[end], nums[start]
	reverseArray(nums, start+1, end-1)
}

/** Reverse an array using recursion
* Approach: Using single pointer
 */

func reverseArraySinglePointer(nums []int, start int) {
	n := len(nums)
	if start >= n/2 {
		return
	}
	nums[start], nums[n-start-1] = nums[n-start-1], nums[start]
	reverseArraySinglePointer(nums, start+1)
}

/**
* Check if a given string in pallindrome or not
* Return true or false
 */

func isPallindrome(str string, i int) bool {
	n := len(str)
	if i >= n/2 {
		return true
	}

	if str[i] != str[n-i-1] {
		return false
	}

	return isPallindrome(str, i+1)
}

/** Multiple Recursion Calls
* Point to remember: In multiple recursion calls, the first calls always finish first untill base condition then second recursive call happens.
 */

func fibonacciNumbers(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciNumbers(n-2) + fibonacciNumbers(n-1) // here is the multiple call for recursion
}

/**
* Print All the subsequences of an array
* Approach: Recursive
* Example: {3,1,2}
 */

func printSubsequences(i int, nums []int, arr []int) {
	if i >= len(nums) {
		fmt.Println(arr)
		return
	}

	// Pick the element into the subsequences
	arr = append(arr, nums[i])
	printSubsequences(i+1, nums, arr)

	// Not pick the element into subsequences
	arr = append([]int{}, arr[:len(arr)-1]...)
	printSubsequences(i+1, nums, arr)
}

/**
* Problem: 2133. Check if Every Row and Column Contains All Numbers
* An n x n matrix is valid if every row and every column contains all the integers from 1 to n (inclusive).
* Given an n x n integer matrix matrix, return true if the matrix is valid. Otherwise, return false.
* Input: matrix = [[1,2,3],[3,1,2],[2,3,1]]
* Output: true
* Explanation: In this case, n = 3, and every row and column contains the numbers 1, 2, and 3. Hence, we return true.
*/

func checkValid(matrix [][]int) bool {
	return isCheckValid(matrix, 0, 0)
}

func isCheckValid(matrix [][]int, irow int, icol int) bool {
	mflatMatrix := []int{}
	for row := irow; row < len(matrix); row++ {
		for col := icol; col < len(matrix[0]); col++ {
			// mflatMatrix = append(mflatMatrix, matrix[row][col])
			if !isValidRowCol(matrix, row, col, matrix[row][col]) {
				return false
			} else {
				return isCheckValid(matrix, row, col + 1)
			}
		}
		icol = 0
	}
	fmt.Println(mflatMatrix)
	return true
}

func isValidRowCol(matrix [][]int, row int, col int, val int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[row][i] == val && i != col {
			return false
		}

		if matrix[i][col] == val && i != row {
			return false
		}
	}
	return true
}
