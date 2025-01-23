## What is HashMap?

### A HashMap is a data structure that maps keys to values using a hash table as its underlying implementation. It allows for fast lookups, insertions, and deletions by leveraging the key's hash code to determine its location in memory.

### Key Characteristics of HashMaps
- <b>Key-Value Pairs</b>:
    - A HashMap stores data in the form of key-value pairs.
    - Each key in a HashMap is unique, but values can be duplicated.
- <b>Fast Operations</b>: 
    - Average-case O(1) time complexity for operations like `get`, `put`, and `remove`, assuming a good hash function.
- <b>Unordered</b>:
    - The order of keys and values in a HashMap is not guaranteed because the internal storage depends on hash codes.
- <b>Collision Handling</b>:
    - When multiple keys generate the same hash code (a collision), HashMaps use techniques like **chaining** or **open addressing** to resolve conflicts.
- <b>Hashing</b>: 
    - A hash function is used to compute a hash code for each key, which determines its "bucket" location in memory.

### How a HashMap Works?
- When a key-value pair is inserted, the key is passed through a hash function, generating a hash code.
- The hash code determines the bucket (location) in the hash table where the value will be stored.
- If a **collision** occurs (two keys have the same hash code), the HashMap uses a collision resolution strategy (like linked list chaining) to store multiple values in the same bucket.
- When retrieving a value, the key is hashed again, and the corresponding bucket is checked for the value.

### Use Cases for HashMaps:
- <b>Caching</b>: 
    - Store frequently accessed data for quick retrieval.
- <b>Database Indexing</b>: 
    - Implement lookup tables to access records efficiently.
- <b>Counting Frequencies</b>:
    - Count occurrences of elements in a dataset.
- <b>Deduplication</b>: 
    - Identify duplicates using keys.
- <b>Implementing LRU Caches</b>: 
    - Combine a HashMap with other data structures (e.g., linked lists) to create efficient LRU caches.
    
### Exapmle:
Go's map is a built-in data structure that functions like a HashMap.

```go
package main

import "fmt"

func main() {
    myMap := make(map[string]int)
    myMap["apple"] = 3
    myMap["banana"] = 2

    fmt.Println(myMap["apple"])  // Output: 3
    delete(myMap, "banana")
    fmt.Println(myMap)           // Output: map[apple:3]
}
```
