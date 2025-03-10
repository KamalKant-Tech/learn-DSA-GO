package main

import (
	_ "data-structure/Array/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(superpalindromesInRange("40000000000000000", "50000000000000000"))
	// s := "catsanddog"
	// words := []string{"cat", "cats", "and", "sand", "dog"}
	// fmt.Println(wordBreakUsingDP(s, words))

	// graph := [][]int{{0, 1, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 1}, {1, 0, 1, 0}}
	// graphInput := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}}
	// graphInput := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}}
	// // create a graph len of four because the vertices are 4 which is 0 -> 3
	// graph := make([][]int, 4)
	// for i := range graph {
	// 	graph[i] = make([]int, 4)
	// }
	// for _, edge := range graphInput {
	// 	graph[edge[0]][edge[1]] = 1
	// 	graph[edge[1]][edge[0]] = 1
	// }
	// fmt.Println(graph)
	// fmt.Println(graphColoring(graph, 3))

	fmt.Println(minCutOptimalUsingDp("abcdefghi"))
}

/*
*
* Problem: 906. Super Palindromes
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

func wordBreakBruteForceUsingBacktracking(s string, wordDict []string) []string {
	fmt.Println("Given Array: ", s, wordDict)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	result := []string{}

	tmp := []string{}
	wordBreakRecursive(s, wordSet, 0, &result, &tmp)
	return result
}

func wordBreakRecursive(s string, wordDict map[string]bool, start int, result *[]string, tmp *[]string) {
	if start == len(s) {
		// If we reached the end, join the temporary results and add to final results
		*result = append(*result, strings.Join(*tmp, " "))
		return
	}

	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		// Check if the current substring is a valid word
		if wordDict[word] {
			// Add the current word to the temporary result
			*tmp = append(*tmp, s[start:end])
			// Recur with the next part of the string
			wordBreakRecursive(s, wordDict, end, result, tmp)
			// Backtrack: remove the last word to explore other possibilities
			*tmp = (*tmp)[:len(*tmp)-1]
		}
	}
}

func wordBreakUsingDP(s string, wordDict []string) []string {
	fmt.Println("Given :", s, wordDict)
	dp := map[int][]string{}
	isExistWordExistMap := make(map[string]bool)
	// fmt.Println(len(s))
	for startIndex := len(s); startIndex >= 0; startIndex-- {
		validSentences := []string{}
		for endIndex := startIndex; endIndex < len(s); endIndex++ {
			currentWord := s[startIndex : endIndex+1]
			// fmt.Println(currentWord, dp, startIndex, endIndex+1)
			// 9,10 g
			// 8, 9 0
			// 8, 10 og
			// 7, 8 d
			// 7, 9 do
			// 7, 10 dog
			// 6, 7 d
			// 6, 8 dd
			// 6, 9 ddo
			// 6, 10 ddog

			// 5, 6 n
			// 5, 7 nd
			// 5, 8 ndd
			// 5, 9 nddo
			// 5, 10 nddog

			// 4, 5 a
			// 5, 6 an
			// 5, 7 and
			if isExistWordExistMap[currentWord] {
				continue
			}
			if isInWordDict(wordDict, currentWord) {
				if endIndex == len(s)-1 {
					// fmt.Println(" Inside:", endIndex, currentWord)
					validSentences = append(validSentences, currentWord)
				} else {
					// If it's not the last word, append it to each sentence formed by the remaining substring
					sentencesFromNextIndex, exists := dp[endIndex+1]
					// fmt.Println("Inside Else: ", sentencesFromNextIndex, endIndex, validSentences)
					if exists {
						for _, sentence := range sentencesFromNextIndex {
							// fmt.Println(currentWord, sentence, endIndex+1)
							validSentences = append(validSentences, currentWord+" "+sentence)
						}
					}
				}
			} else {
				isExistWordExistMap[currentWord] = true
			}

		}
		fmt.Println("Valid Sentences:", isExistWordExistMap)
		// Store the valid sentences in dp
		dp[startIndex] = validSentences
	}
	// fmt.Println(dp)
	// Return the sentences formed from the entire string
	return dp[0]
	// return []string{}
}

func isInWordDict(wordDict []string, s string) bool {
	for _, v := range wordDict {
		if v == s {
			return true
		}
	}
	return false
}

// M Coloring Problem
// Given an undirected graph and a number m, determine if the graph can be colored with at most m colors such that no two adjacent vertices of the graph are colored with the same color. Here coloring of a graph means the assignment of colors to all vertices. Print Yes if the graph can be colored, No otherwise.
// Input: graph = [[0, 1, 1, 1], [1, 0, 1, 0], [1, 1, 0, 1], [1, 0, 1, 0]], m = 3
// Output: Yes
// Explanation: The given graph can be colored with 3 colors such that no two adjacent vertices are colored with the same color.
// 0	1	2	3
// 0	0	1	1	1
// 1	1	0	1	0
// 2	1	1	0	1
// 3	1	0	1	0
// Solution: Backtracking

// To visualize the graph, you can use a simple text-based representation or a graphical tool. Here's a text-based representation:

// Graph:
// 0 -- 1
// | \  |
// |  \ |
// 3 -- 2

// Adjacency Matrix:
// 0 1 1 1
// 1 0 1 0
// 1 1 0 1
// 1 0 1 0

// The graph can be visualized as a square with diagonals connecting opposite corners.
func graphColoring(graph [][]int, m int) bool {
	colors := make([]int, len(graph))
	return graphColoringRecursive(graph, m, colors, 0)
}

func graphColoringRecursive(graph [][]int, m int, colors []int, i int) bool {
	if i == len(graph) {
		return true
	}
	for color := 1; color <= m; color++ {
		if isSafe(graph, colors, i, color) {
			colors[i] = color
			if graphColoringRecursive(graph, m, colors, i+1) {
				return true
			}
			colors[i] = 0
		}
	}
	return false
}

func isSafe(graph [][]int, colors []int, i, color int) bool {
	for j := 0; j < len(graph); j++ {
		if graph[i][j] == 1 && colors[j] == color {
			return false
		}
	}
	return true
}

// 132. Palindrome Partitioning II
// Given a string s, partition s such that every substring of the partition is a palindrome Palindrome.
// A palindrome is a string that reads the same forward and backward.
// Return the minimum cuts needed for a palindrome partitioning of s.
// Input: s = "aab"
// Output: 1
// Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

func minCutBruteForce(s string) int {
	cuts := 0
	minCutRecursive(s, 0, []string{}, &cuts)
	// fmt.Println(result)
	return cuts
}

func minCutRecursive(s string, index int, pallindromString []string, cuts *int) {
	if index == len(s) {
		if *cuts == 0 {
			*cuts = len(pallindromString) - 1
		}

		if len(pallindromString)-1 < *cuts {
			*cuts = len(pallindromString) - 1
		}
		return
	}
	for i := index; i < len(s); i++ {
		if isPalindromeStr(s[index : i+1]) {
			pallindromString = append(pallindromString, s[index:i+1])
			minCutRecursive(s, i+1, pallindromString, cuts)
			pallindromString = pallindromString[:len(pallindromString)-1]
		}
	}
}

func isPalindromeStr(s string) bool {
	i := 0
	j := len(s) - 1
	if len(s) == 0 {
		return false
	}
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func minCutBetterUsingDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Create a 2D slice to store whether a substring is a palindrome
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Initialize the isPalindrome table
	for i := 0; i < n; i++ {
		isPalindrome[i][i] = true
	}
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if length == 2 {
				isPalindrome[i][j] = (s[i] == s[j])
			} else {
				isPalindrome[i][j] = (s[i] == s[j] && isPalindrome[i+1][j-1])
			}
		}
	}

	// Create a slice to store the minimum cuts needed for each prefix of the string
	minCuts := make([]int, n)
	for i := range minCuts {
		if isPalindrome[0][i] {
			minCuts[i] = 0
		} else {
			minCuts[i] = i
			for j := 0; j < i; j++ {
				if isPalindrome[j+1][i] {
					minCuts[i] = min(minCuts[i], minCuts[j]+1)
				}
			}
		}
	}

	return minCuts[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCutOptimalUsingDp(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Create a slice to store the minimum cuts needed for each prefix of the string
	minCuts := make([]int, n)

	// Initialize minCuts with the worst case: cutting between every character
	for i := range minCuts {
		minCuts[i] = i
	}

	// Helper function to expand around the center and check for palindromes
	checkPalindrome := func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			if left == 0 {
				minCuts[right] = 0
			} else {
				minCuts[right] = min(minCuts[right], minCuts[left-1]+1)
			}
			left--
			right++
		}
	}

	// Iterate over each character as a potential center for odd and even palindromes
	for i := 0; i < n; i++ {
		// Check odd-length palindromes
		checkPalindrome(i, i)

		// Check even-length palindromes
		checkPalindrome(i, i+1)
	}
	fmt.Println(minCuts)
	return minCuts[n-1]
}
