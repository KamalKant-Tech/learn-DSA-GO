package main

import (
	"container/heap"
	"fmt"
)

// This package it to implement priority queue using heap data structure

type Item struct {
	Val      int
	Priority int
	Index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func main() {
	// Create the item map items
	// items := map[string]int{
	// 	"apple":      3,
	// 	"banana":     2,
	// 	"orange":     6,
	// 	"strawberry": 4,
	// }
	// create the priority queue
	// pq := make(PriorityQueue, len(items))

	// i := 0
	// // for item, priority := range items {
	// // 	pq[i] = &Item{
	// // 		Val:      item,
	// // 		Priority: priority,
	// // 		Index:    i,
	// // 	}
	// // 	i++
	// // }
	// // intialixe the heap queue
	// heap.Init(&pq)
	// Insert a new element to the queue

	// item := &Item{
	// 	Val:      "Pineapple",
	// 	Priority: 10,
	// }

	// heap.Push(&pq, item)
	// pq.update(item, item.Val, 10)

	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*Item)
	// 	fmt.Printf("%.2d %s %d \n", item.Priority, item.Val, item.Index)
	// }
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(TopKFrequentElements(nums, 2))

}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	oldQueue := *pq
	n := len(oldQueue)
	item := oldQueue[n-1]
	oldQueue[n-1] = nil // memory leackage
	//item.Index = -1
	*pq = oldQueue[:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, val int, priority int) {
	item.Val = val
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

/**
* Problem: 347. Top K Frequent Elements
* Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
* Input: nums = [1,1,1,2,2,3], k = 2
* Output: [1,2]
* Intuition: Using Priority Queue
 */

func TopKFrequentElements(nums []int, k int) []int {
	var result = make([]int, k)
	var freqMap = map[int]int{}
	for _, v := range nums {
		freqMap[v]++
	}
	fmt.Println(len(result))
	var pq = make(PriorityQueue, len(freqMap))
	i := 0
	for val, priority := range freqMap {
		pq[i] = &Item{
			Val:      val,
			Priority: priority,
		}
		i++
	}
	heap.Init(&pq)

	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.Val)
	}
	return result
}
