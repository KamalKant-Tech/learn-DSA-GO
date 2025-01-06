package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))
}
func isHappy(n int) bool {
	numbers := map[int]int{}
	return isHappyExt(n, numbers)
}

func isHappyExt(n int, numbers map[int]int) bool {
	sum := 0

	if n == 1 {
		return true
	}

	if exist, ok := numbers[n]; ok && exist > 1 {
		return false
	}

	for n > 0 {
		modN := n % 10
		sum += modN * modN
		n = n / 10
		fmt.Println(sum, n)
	}

	numbers[sum]++
	return isHappyExt(sum, numbers)
}

func hashSetsImlementation() {
	set := make(map[string]struct{})
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["apple"] = struct{}{} // Duplicate, no effect

	for key := range set {
		fmt.Println(key)
	}
}
