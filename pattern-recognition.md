## 🎯 **Daily Pattern Recognition Training Plan**

## 🎯 **Pattern Recognition Strategy**

### **Step 1: Read the Problem & Identify Keywords**

| **Keywords** | **Pattern** | **Data Structure** |
|--------------|-------------|-------------------|
| "subarray", "substring", "consecutive" | Sliding Window | Array/String |
| "two elements sum to target" | Two Pointers | Sorted Array |
| "frequency", "anagram", "duplicate" | Hash Map | Map/Set |
| "path", "ways", "maximum/minimum" | Dynamic Programming | DP Array |
| "sorted array", "search", "find element" | Binary Search | Array |
| "level order", "shortest path" | BFS | Queue |
| "all paths", "combinations", "permutations" | DFS/Backtracking | Stack/Recursion |

### **Step 2: Ask These Questions**

1. **Input Type**: Array? String? Tree? Graph?
2. **Output Type**: Single value? Array? Boolean?
3. **Constraints**: Sorted? Unique elements? Size limits?
4. **Optimization**: Time vs Space tradeoff?

### **Step 3: Pattern Matching Flow**

```go
// Decision Tree for Pattern Selection
func identifyPattern(problemDescription string) string {
    if contains(problemDescription, []string{"subarray", "substring"}) {
        if contains(problemDescription, []string{"maximum", "minimum", "longest"}) {
            return "Sliding Window"
        }
        if contains(problemDescription, []string{"sum equals k", "target sum"}) {
            return "Prefix Sum + Hash Map"
        }
    }
    
    if contains(problemDescription, []string{"two elements", "pair", "sorted array"}) {
        return "Two Pointers"
    }
    
    if contains(problemDescription, []string{"ways", "paths", "maximum", "minimum"}) &&
       contains(problemDescription, []string{"recursive", "overlapping"}) {
        return "Dynamic Programming"
    }
    
    // ... more pattern matching logic
    return "Unknown Pattern"
}
```

### **Today's Practice Session - 3 Problems**

Let me walk you through **3 carefully chosen problems** to practice pattern recognition step by step:

---

## **Problem 1: LeetCode 121 - Best Time to Buy and Sell Stock**

**Problem Statement:**
You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

**Example:** `Input: prices = [7,1,5,3,6,4]`, `Output: 5`

### **🔍 Pattern Recognition Process:**

**Step 1: Identify Keywords & Constraints**
- ✅ **"array"** → Array-based problem
- ✅ **"maximum profit"** → Optimization problem
- ✅ **"single day to buy, different day to sell"** → Need to track state
- ✅ **"future day"** → Order matters, can't go backwards

**Step 2: Ask Pattern Questions**
1. **Input/Output:** Array → Single value (max profit)
2. **Optimization:** Looking for maximum
3. **Dependencies:** Future decisions depend on past choices
4. **State:** Need to track "minimum price seen so far"

**Step 3: Pattern Match**
- ❌ **Two Pointers?** No, we need to track minimum over time
- ❌ **Sliding Window?** No fixed/variable window size
- ❌ **Hash Map?** No lookup or frequency counting
- ✅ **Single Pass Tracking** → Keep track of minimum price and maximum profit

**Pattern Identified: Single Pass Optimization**

```go
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    
    minPrice := prices[0]
    maxProfit := 0
    
    for i := 1; i < len(prices); i++ {
        // Update minimum price seen so far
        if prices[i] < minPrice {
            minPrice = prices[i]
        }
        
        // Calculate profit if we sell today
        profit := prices[i] - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    
    return maxProfit
}
```

---

## **Problem 2: LeetCode 3 - Longest Substring Without Repeating Characters**

**Problem Statement:**
Given a string `s`, find the length of the longest substring without repeating characters.

**Example:** `Input: s = "abcabcbb"`, `Output: 3` (substring "abc")

### **🔍 Pattern Recognition Process:**

**Step 1: Identify Keywords**
- ✅ **"substring"** → Contiguous elements
- ✅ **"longest"** → Optimization (maximum)
- ✅ **"without repeating"** → Need to track uniqueness
- ✅ **"characters"** → Frequency/existence tracking

**Step 2: Ask Pattern Questions**
1. **Window Nature:** Contiguous substring → Variable size window
2. **Constraint:** No repeating characters → Shrink when duplicate found
3. **Optimization:** Find maximum length → Track max as we go
4. **Tracking:** Need to know character positions → Hash map

**Step 3: Pattern Match**
- ❌ **Two Pointers?** Not exactly, need window expansion/contraction
- ✅ **Sliding Window + Hash Map** → Variable size window with constraint
- ❌ **DP?** No overlapping subproblems
- ❌ **Binary Search?** No sorted property

**Pattern Identified: Sliding Window (Variable Size) + Hash Map**

```go
func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    left := 0
    maxLen := 0
    
    for right := 0; right < len(s); right++ {
        // If character already exists and is in current window
        if lastIndex, exists := charIndex[s[right]]; exists && lastIndex >= left {
            left = lastIndex + 1  // Shrink window
        }
        
        charIndex[s[right]] = right  // Update character position
        maxLen = max(maxLen, right-left+1)  // Update max length
    }
    
    return maxLen
}
```

---

## **Problem 3: LeetCode 70 - Climbing Stairs**

**Problem Statement:**
You are climbing a staircase. It takes `n` steps to reach the top. Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

**Example:** `Input: n = 3`, `Output: 3` (1+1+1, 1+2, 2+1)

### **🔍 Pattern Recognition Process:**

**Step 1: Identify Keywords**
- ✅ **"distinct ways"** → Counting problem
- ✅ **"can either climb 1 or 2"** → Multiple choices at each step
- ✅ **"reach the top"** → Target state
- ✅ **"staircase"** → Sequential/ordered problem

**Step 2: Ask Pattern Questions**
1. **Problem Type:** Count total ways → DP indicator
2. **Recurrence:** Ways to reach step n = ways to reach (n-1) + ways to reach (n-2)
3. **Base Cases:** Need starting conditions
4. **Overlap:** Same subproblems solved multiple times

**Step 3: Pattern Match**
- ❌ **Array Traversal?** No array given
- ❌ **Hash Map?** No key-value mapping needed
- ❌ **Sliding Window?** No window concept
- ✅ **Dynamic Programming** → Classic DP pattern (Fibonacci-like)

**Pattern Identified: Linear DP (Fibonacci Pattern)**

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    // Bottom-up DP approach
    prev2 := 1  // ways to reach step 1
    prev1 := 2  // ways to reach step 2
    
    for i := 3; i <= n; i++ {
        current := prev1 + prev2  // ways to reach step i
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}
```

---

## 📋 **Self-Assessment Checklist**

After each problem, ask yourself:

### **Recognition Questions:**
1. ✅ Did I identify the key indicators correctly?
2. ✅ Did I ask the right pattern questions?
3. ✅ Did I eliminate wrong patterns systematically?
4. ✅ Did I arrive at the correct pattern?

### **Implementation Questions:**
1. ✅ Does my solution match the identified pattern?
2. ✅ Did I handle edge cases appropriately?
3. ✅ Is the time/space complexity optimal for this pattern?

---

## 🎯 **Your Daily Practice Framework**

### **Week 1 Schedule (Starting Tomorrow):**

**Monday - Array Patterns:**
- Two Pointers: LeetCode 167 (Two Sum II)
- Sliding Window: LeetCode 209 (Minimum Size Subarray Sum)

**Tuesday - Hash Map Patterns:**
- Frequency Counter: LeetCode 242 (Valid Anagram)
- Lookup Optimization: LeetCode 1 (Two Sum)

**Wednesday - DP Patterns:**
- Linear DP: LeetCode 198 (House Robber)
- Grid DP: LeetCode 62 (Unique Paths)

**Thursday - Tree Patterns:**
- DFS: LeetCode 104 (Maximum Depth of Binary Tree)
- BFS: LeetCode 102 (Binary Tree Level Order Traversal)

**Friday - Mixed Review:**
- Pick one problem from each day and solve without looking at notes

### **Daily Process (20-30 minutes):**

1. **Read Problem (2 mins):** Identify keywords and constraints
2. **Pattern Analysis (5 mins):** Use the recognition process above
3. **Write Approach (3 mins):** Outline solution in comments
4. **Implement (10 mins):** Code the solution
5. **Review (5 mins):** Check against pattern and optimize

---

## 🚀 **Pattern Recognition Cheat Sheet**

Keep this handy for quick reference:

```go
/*
PATTERN DECISION TREE:

1. ARRAY/STRING + "subarray/substring" + "maximum/minimum" → Sliding Window
2. ARRAY + "two elements" + "target sum" → Two Pointers
3. ANY + "frequency/count/anagram" → Hash Map
4. ANY + "ways to do something" + "optimal choice" → Dynamic Programming
5. TREE/GRAPH + "traverse all nodes" → DFS/BFS
6. SORTED ARRAY + "find element" → Binary Search
7. "shortest path" + "unweighted graph" → BFS
8. "all combinations/permutations" → Backtracking
9. "merge/sort" + "divide and conquer" → Merge Sort pattern
10. "minimum/maximum in range" → Segment Tree/Heap
*/
```

---

## 📅 **Tomorrow's Problems**

I'll give you these 2 problems for tomorrow (I can help you recognize patterns):

1. **LeetCode 217 - Contains Duplicate** (Hash Set Pattern)
2. **LeetCode 53 - Maximum Subarray** (Kadane's Algorithm)

**Would you like me to:**
1. Walk through tomorrow's problems right now?
2. Create a weekly schedule with specific problems?
3. Focus on a particular pattern you find most difficult?
4. Help you analyze a specific problem you're currently stuck on?

The key is **consistent daily practice** with **guided pattern recognition**. After 2-3 weeks, you'll start seeing patterns automatically! 🎯