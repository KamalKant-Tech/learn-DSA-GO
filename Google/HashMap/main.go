package main

import "fmt"

type Pair struct {
	key   int
	value int
}

type MyHashMap struct {
	size    int
	buckets [][]Pair
}

func Constructor() MyHashMap {
	return MyHashMap{
		size:    1000,
		buckets: make([][]Pair, 1000),
	}
}

func (this *MyHashMap) hashFunction(key int) int {
	return key % this.size
}

func (this *MyHashMap) Put(key int, value int) {
	hashKey := this.hashFunction(key)
	bucket := this.buckets[hashKey]
	// Check if key already exists; if so, update its value
	for i, pair := range bucket {
		if pair.key == key {
			this.buckets[hashKey][i].value = value
			return
		}
	}

	// If key does not exist, append the new key-value pair
	this.buckets[hashKey] = append(this.buckets[hashKey], Pair{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	hashKey := this.hashFunction(key)
	bucket := this.buckets[hashKey]

	// Search for the key in the bucket and return its value if found
	for _, pair := range bucket {
		if pair.key == key {
			return pair.value
		}
	}
	return -1 // Key not found
}

func (this *MyHashMap) Remove(key int) {
	// hashkey := this.hashFunction(key)
	// bucket := this.buckets[hashkey]
	hashKey := this.hashFunction(key)
	bucket := this.buckets[hashKey]

	// Iterate over the bucket and remove the pair if found
	for i, pair := range bucket {
		if pair.key == key {
			// Remove the pair by slicing it out of the bucket
			this.buckets[hashKey] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func main() {
	obj := Constructor()
	obj.Put(1, 100)
	obj.Put(1, 300)
	param_2 := obj.Get(1)
	obj.Remove(1)
	fmt.Println(param_2, obj.buckets)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
