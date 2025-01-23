package main

import "fmt"

type MyHashSet struct {
	Size   int
	Bucket [][]int
}

func main() {
	/**
	 * Your MyHashSet object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Add(key);
	 * obj.Remove(key);
	 * param_3 := obj.Contains(key);
	 */

	obj := Constructor()
	obj.Add(2)
	fmt.Println(obj.Bucket)
}

func Constructor() MyHashSet {
	size := 1000
	bucket := make([][]int, size)
	return MyHashSet{
		Size: size, Bucket: bucket,
	}
}

func (this *MyHashSet) hashFunction(key int) int {
	return key % this.Size
}

func (this *MyHashSet) Add(key int) {
	hashKey := this.hashFunction(key)
	if !this.Contains(key) {
		this.Bucket[hashKey] = append(this.Bucket[hashKey], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	hashkey := this.hashFunction(key)
	bucket := this.Bucket[hashkey]
	for i, val := range bucket {
		if val == key {
			this.Bucket[hashkey] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	hashKey := this.hashFunction(key)
	bucket := this.Bucket[hashKey]
	for _, val := range bucket {
		if val == key {
			return true
		}
	}
	return false
}
