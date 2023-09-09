package main

import "fmt"

func main() {
	fmt.Println(SumOfNNumbersSolution(3))
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
