package main

import (
	_ "data-structure/Array/input"
	"fmt"
	"math"
	"strconv"
)

func main() {
	// fmt.Println(superpalindromesInRange("40000000000000000", "50000000000000000"))
	// s := "barfoofoobarthefoobarman"
	// words := []string{"bar", "foo", "the"}
	// fmt.Println(findSubstring(s, words))
}

/*
*
* 906. Super Palindromes
* Let's say a positive integer is a super-palindrome if it is a palindrome, and it is also the square of a palindrome.
* Given two positive integers left and right represented as strings, return the number of super-palindromes integers in the inclusive range [left, right].
*
// Example:
// Input: left = "4", right = "1000"
// Output: 4
// Explanation: 4, 9, 121, and 484 are superpalindromes.
// Note that 676 is not a superpalindrome: 26 * 26 = 676, but 26 is not a palindrome.
*/

// func superpalindromesInRange(left string, right string) int {
// 	l, _ := strconv.ParseInt(left, 10, 64)
// 	r, _ := strconv.ParseInt(right, 10, 64)
// 	lSqrt := int(math.Sqrt(float64(l)))
// 	rSqrt := int(math.Sqrt(float64(r)))
// 	count := 0
// 	for i := lSqrt; i <= rSqrt; i++ {
// 		fmt.Println(i)
// 		sqrt := int(math.Sqrt(float64(i)))
// 		if sqrt*sqrt > int(rSqrt) {
// 			break
// 		}
// 		if !(sqrt*sqrt == int(i)) || !isPalindrome(fmt.Sprint(sqrt)) || !isPalindrome(fmt.Sprint(i)) {
// 			continue
// 		}

// 		count++
// 	}
// 	return count
// }

// func isPalindrome(s string) bool {
// 	i, j := 0, len(s)-1
// 	for i <= j {
// 		if string(s[i]) != string(s[j]) {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }

func superpalindromesInRange(left string, right string) int {
	result := 0
	// Calculate the square roots of the left and right bounds, rounding up for the left bound and down for the right bound
	l, r := sqrt(left, math.Ceil), sqrt(right, math.Floor)
	// Count super-palindromes with even length
	countSuperpalindromes(true, l, r, &result)
	// Count super-palindromes with odd length
	countSuperpalindromes(false, l, r, &result)

	return result
}

func countSuperpalindromes(isEven bool, l, r int, result *int) {
	for i := 1; ; i++ {
		pallindrome, reversed := i, i

		// fmt.Println(p, reversed)
		if !isEven {
			reversed /= 10
		}
		// Generate the palindrome by appending the reverse of the number
		for ; reversed > 0; reversed /= 10 {
			pallindrome *= 10
			pallindrome += reversed % 10
		}
		// If the palindrome exceeds the upper bound, stop the loop
		if pallindrome > r {
			break
		}
		// Check if the palindrome is within the lower bound and if its square is also a palindrome
		if pallindrome >= l && isPalindrome(pallindrome*pallindrome) {
			(*result)++
		}
	}
}

// isPalindrome checks if an integer is a palindrome.
func isPalindrome(x int) bool {
	y := 0
	for ; x > y; x /= 10 {
		y *= 10
		y += x % 10
	}

	return x == y || x == y/10
}

// sqrt parses a string to a float64, computes its square root, and applies a rounding function.
func sqrt(s string, roundFunc func(float64) float64) int {
	f, _ := strconv.ParseFloat(s, 64)
	return int(roundFunc(math.Sqrt(f)))
}

/**
* Problem: 41. First Missing Positives
* Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.
* You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.
* Input: nums = [1,2,0]
* Output: 3
* Explanation: The numbers in the range [1,2] are all in the array.
 */
// func firstMissingPositive(nums []int) int {
// 	smallestNumber := math.MaxInt64
// 	largestnumber := math.MinInt64
// 	for _, v := range nums {
// 		if v > 0 && smallestNumber > v {
// 			smallestNumber = v
// 		}

// 		if largestnumber < v {
// 			largestnumber = v
// 		}
// 	}

// 	if smallestNumber > 0 && smallestNumber != 1 {
// 		return 1
// 	}
// 	// fmt.Println(smallestNumber, largestnumber)
// 	// // smallestNumber = 1 , [1,2,0]
// 	for i := 0; i < largestnumber; i++ {
// 		if i > 0 && !contains(nums, i) {
// 			return i
// 		}
// 	}
// 	// fmt.Println(smallestNumber, largestnumber)
// 	return largestnumber + 1
// }

//	func contains(nums []int, target int) bool {
//		for _, v := range nums {
//			if v == target {
//				return true
//			}
//		}
//		return false
//	}
//
// Solution: Using boolean hash map
func firstMissingPositive(nums []int) int {
	boleanMap := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		boleanMap[nums[i]] = true
	}
	fmt.Println(boleanMap)
	for i := 0; i < len(nums); i++ {
		if _, ok := boleanMap[i]; !ok && i > 0 {
			return i
		}
	}

	return len(nums) + 1
}

// Solution: Using cyclic sort

func firstMissingPositiveUsingInPlaceAlgo(nums []int) int {
	n := len(nums)
	containesOne := false

	// Replace negative numbers, zeros,
	// and numbers larger than n with 1s.
	// After this nums contains only positive numbers.
	for i := 0; i < n; i++ {
		// Check whether 1 is in the original array
		if nums[i] == 1 {
			containesOne = true
		} else if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	if !containesOne {
		return 1
	}

	// Mark whether integers 1 to n are in nums
	// Use index as a hash key and negative sign as a presence detector.
	for i := 0; i < n; i++ {
		val := int(math.Abs(float64(nums[i])))
		if val == n {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[val] = -int(math.Abs(float64(nums[val])))
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	// nums[0] stores whether n is in nums
	if nums[0] > 0 {
		return n
	}

	// If nums contained all elements 1 to n
	// the smallest missing positive number is n + 1
	return n + 1
}

/**
* Problem: 30. Substring with Concatenation of All Words
* You are given a string s and an array of strings words. All the strings of words are of the same length.
* A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.
* For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
* Input: s = "barfoothefoobarman", words = ["foo","bar"]
* Output: [0,9]
* The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
* The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
* Solution: Using sliding window
 */

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordsCount := len(words)
	substrLen := wordLen * wordsCount

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	result := []int{}

	// Sliding Window Over the String:
	// The outer loop iterates over each possible starting index i in s where a valid substring could begin.
	// For each starting index i, a new map seenWords is created to track the words seen in the current window.
	// The inner loop iterates over each word in the words list, checking if the substring from s matches any word in words.
	// * wordIndex calculates the starting index of the current word in the substring.
	// * word extracts the current word from s.
	// * If the word exists in wordMap and hasn't been seen more times than it appears in words, it increments the count in seenWords.
	// * If the word doesn't exist in wordMap or has been seen too many times, it breaks out of the inner loop.
	// If all words are matched (j == wordsCount), the starting index i is added to the result list.
	// Iterate over each possible starting index within the first word length

	for i := 0; i <= len(s)-substrLen; i++ {
		seenWords := make(map[string]int)
		j := 0
		for ; j < wordsCount; j++ {
			wordIndex := i + j*wordLen
			word := s[wordIndex : wordIndex+wordLen]
			if count, exists := wordMap[word]; exists {
				seenWords[word]++
				if seenWords[word] > count {
					break
				}
			} else {
				break
			}
		}
		if j == wordsCount {
			result = append(result, i)
		}
	}
	return result
}

func findSubstringOptimal(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordsCount := len(words)

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	result := []int{}

	// 1. Outer Loop (Sliding Window Initialization)
	// This loop iterates over each possible starting index within the first word length. This ensures that all possible starting positions are considered.
	for offset := 0; offset < wordLen; offset++ {
		// Initialization
		// left and right pointers are initialized to the current offset.
		// seenWords is a map to keep track of the words seen in the current window.
		// count keeps track of the number of valid words seen in the current window.
		left := offset
		right := offset
		seenWords := make(map[string]int)
		count := 0

		// 3. Inner Loop (Sliding Window)
		// This loop slides the window over the string s in increments of wordLen
		for right+wordLen <= len(s) {
			// Extracting the Word
			// Extracts a word of length wordLen from the current position of right and then moves the right pointer forward by wordLen
			word := s[right : right+wordLen]
			right += wordLen

			// Checking the Word
			// Checks if the extracted word exists in the wordMap
			if _, exists := wordMap[word]; exists {
				// Valid Word Handling
				// If the word is valid, it increments its count in seenWords and increments the count
				seenWords[word]++
				count++
				// Adjusting the Window
				// If the count of the current word exceeds its allowed count in wordMap, it adjusts the window by moving the left pointer forward and decrementing the counts accordingly.
				for seenWords[word] > wordMap[word] {
					leftWord := s[left : left+wordLen]
					seenWords[leftWord]--
					count--
					left += wordLen
				}
				// Checking for Complete Match
				// If the count of valid words matches the total number of words, it means a valid substring is found, and the starting index left is added to the result
				if count == wordsCount {
					result = append(result, left)
				}
			} else {
				// Invalid Word Handling
				// If the word is not valid, it resets the seenWords map, count, and moves the left pointer to the current right position.
				seenWords = make(map[string]int)
				count = 0
				left = right
			}
		}
	}
	return result
}
