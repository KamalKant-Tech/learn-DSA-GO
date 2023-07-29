package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 2, 10, 8, 3}
	fmt.Println("Given Array :", arr)
	low := 0
	high := len(arr) - 1
	Partition(arr, low, high)
	fmt.Println(arr)

}

// Selection Sort Pick an element and compare with all the remaining elements in an array
// Time complexity : O(n*2)
func SelectionSort() {
	arr := []int{1, 6, 2, 9, 23, 56, 78, 65, 3, 4}
	fmt.Println("Before Selection Sorting: ", arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fmt.Printf("i=%d , j=%d, Array : %d \n", i, j, arr)
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

// Bubble Sort Pick an element and compare that element with its adjacent elems
// Time complexity in worst case : O(n*2)
// Time complexity in best case : O(n) if array is sorted
func BubbleSort() {
	arr := []int{1, 6, 2, 9, 23, 56, 78, 65, 3, 4}
	// arr := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr); i++ {
		swapStatus := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapStatus = true
			}
		}
		if !swapStatus {
			break
		}
	}
	fmt.Println(arr)
}

// Insertion sort: pick an lowest or largest element and put this into their respective index
// Time complexity in worst case: O(n*2)
// Best It would be O(n) if the given array is sorted
func InsertionSort() {
	arr := []int{1, 6, 2, 9, 23, 56, 78, 65, 3, 4}
	// arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before Sorting: ", arr)
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			fmt.Println("i, j", i, j)
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	fmt.Println(arr)
}

// Merge sort works on divide and merge techniques
// Time complexity in worst case: O(nlogn)
// Best It would be O(n) if the given array is sorted
// Space comlexity would be O(n)
func MergeSort(arr *[]int, low int, high int) {
	if low >= high {
		return
	}
	mid := int((low + high) / 2)
	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)
	Merge(arr, low, mid, high)
}

func Merge(arr *[]int, low, mid, high int) {
	tmp := []int{}
	left := low
	right := mid + 1

	for left <= mid && right <= high {

		if (*arr)[left] <= (*arr)[right] {
			tmp = append(tmp, (*arr)[left])
			left++
		} else {
			tmp = append(tmp, (*arr)[right])
			right++

		}
	}
	for left <= mid {
		tmp = append(tmp, (*arr)[left])
		left++
	}

	for right <= high {
		tmp = append(tmp, (*arr)[right])

		right++
	}
	for i := low; i <= high; i++ {
		(*arr)[i] = tmp[i-low]
	}
}

// Quick sort works finding the pivot element and put this into correct position
// Time complexity in worst case: O(nlogn)
// Best It would be O(n) if the given array is sorted
// Space comlexity would be O(1)
func QuickSort(arr []int, low int, high int) {
	if low < high {
		pIndex := Partition(arr, low, high)
		QuickSort(arr, low, pIndex-1)
		QuickSort(arr, pIndex+1, high)
	}
}
func Partition(arr []int, low int, high int) int {
	i := low
	j := high
	pivot := arr[low]
	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			// fmt.Printf("Inside first for loop: value- %d, pivot- %d, ith val- %d \n ", arr[i], pivot, i)
			i++
		}
		for arr[j] > pivot && j >= low {
			// fmt.Printf("Inside second for loop: value- %d, pivot- %d, ith val- %d \n ", arr[j], pivot, j)
			j--
		}
		if i < j {
			// fmt.Printf("Swap values: %d, %d, ith value- %d, jth value- %d \n ", arr[i], arr[j], i, j)
			arr[i], arr[j] = arr[j], arr[i]
			//fmt.Println("Sorted Array: ", arr)
		}
	}
	//fmt.Println(arr[i], arr[j], low, j)
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
