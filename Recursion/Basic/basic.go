package main

import "fmt"

func main() {
	//str := "kamal"
	fmt.Println(fibonacciNumbers(4))
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
