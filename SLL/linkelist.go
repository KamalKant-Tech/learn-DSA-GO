package main

import (
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Insert inserts new node at the end of  from linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func AddTwoNumber(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	// byteArr := make(map[int]int)
	l3 := &LinkedList{}
	ptr := l1.head
	ptr2 := l2.head
	// carry := 0
	// j := 0

	for ptr != nil || ptr2 != nil {
		var i, j int
		// byteArr[j] = ptr.value
		// ptr = ptr.next
		// j++

		if ptr == nil {
			i = 0
		} else {
			i = ptr.value
		}
		if ptr2 == nil {
			j = 0
		} else {
			j = ptr2.value
		}
		sum := i + j
		fmt.Printf("Sum of %d + %d = %d", i, j, sum)
		// l3.head.value = sum % 10
		// l3.head = l3.head.next

		if ptr != nil {
			// fmt.Println("Here Pointer 1", ptr.value)
			ptr = ptr.next
		}

		if ptr2 != nil {
			// fmt.Println("Here Pointer 1", ptr.value)
			ptr2 = ptr2.next
		}
		// carry = sum / 10

	}

	// if carry != 0 {
	// 	current.Next = &ListNode{1, nil}
	// }

	// i := 0
	// for ptr2 != nil {
	// 	rem := (byteArr[i] + ptr2.value) / 10

	// 	if (byteArr[i] + ptr2.value) > 9 {
	// 		byteArr[i] = rem
	// 	} else {
	// 		byteArr[i] = (byteArr[i] + ptr2.value) % 10
	// 	}
	// 	fmt.Println(byteArr[i], ptr2.value, rem, byteArr, math.Mod(11, 10))
	// 	// fmt.Printf(" Sum of %d + %d = %d, %d , ")

	// 	if _, found := byteArr[i+1]; found {
	// 		byteArr[i+1] = byteArr[i+1] + rem
	// 	}

	// 	ptr2 = ptr2.next
	// 	i++
	// }

	// for k := 0; k < len(byteArr); k++ {
	// 	l3.Insert(byteArr[k])
	// }
	return l3
}

func TraverseLinkedList(l *LinkedList) {
	fmt.Printf("Linked List: ")
	temp := l.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.value)
		temp = temp.next
	}
	fmt.Printf("NULL")
}

func TraverseNodeList(l *Node) {
	fmt.Printf("Linked List: ")
	temp := l
	for temp != nil {
		fmt.Printf("%d -> ", temp.value)
		temp = temp.next
	}
	fmt.Printf("NULL")
}

func main() {
	l1 := &LinkedList{}
	InsertAtTail(l1, 1)
	InsertAtTail(l1, 2)
	InsertAtTail(l1, 3)
	InsertAtTail(l1, 4)
	InsertAtTail(l1, 5)
	InsertAtTail(l1, 6)

	l2 := &LinkedList{}
	InsertAtTail(l2, 1)
	InsertAtTail(l2, 3)
	InsertAtTail(l2, 4)
	InsertAtTail(l2, 5)
	TraverseLinkedList(l1)
	// TraverseLinkedList(l2)
	// fmt.Println("Before Delete ")
	// TraverseLinkedList(l1)
	// fmt.Println("\nAfter Delete ")
	// deleteAtIthPosition(l1, 4)
	// TraverseLinkedList(l1)
	// deleteAtGivenVal(l1, 23)
	// fmt.Println("\nAfter Delete by given value of node : ")
	// TraverseLinkedList(l1)

	// reverse(l1)
	// TraverseLinkedList(mergeTwoLists(l1, l2))

	// l1Result := deleteAtGivenVal(l1, 1)

	// TraverseNodeList(l1Result)
	// creating a loop in the above linked list
	// l1.head.next.next.next.next = l1.head // comment this then for no loop
	// fmt.Println(l1.head.next.next.next.next)
	fmt.Println(middleNode(l1))
	// if hasCycle(l1) {
	// 	fmt.Println("found loop")
	// } else {
	// 	fmt.Println("no loop")
	// }

	//fmt.Println(hasCycle(l1))

	// Example:
	// listA: 1 -> 3 -> 5 -> 7 -> 9 -> 11
	// listB: 2 -> 4 -> 9 -> 11
	// intersection: 9 -> 11
	listA := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9, Next: &ListNode{Val: 11}}}}}}
	listB := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: listA.Next.Next}}
	intersection := getIntersectionNode(listA, listB)
	fmt.Println("Intersection:", intersection.Val) // Output: 9

}

func InsertAtHead(l *LinkedList, val int) {
	node := Node{value: val}
	if l.head == nil {
		l.head = &node
		l.len++
	} else {
		tmp := l.head
		l.head = &node
		l.head.next = tmp
		l.len++
	}
	// fmt.Println(l.head.value)
}

func InsertAtTail(l *LinkedList, val int) {
	node := &Node{value: val}
	if l.len == 0 {
		l.head = node
		l.len++
		return
	} else {
		ptr := l.head
		for ptr != nil {
			if ptr.next == nil {
				ptr.next = node
				l.len++
				return
			}
			ptr = ptr.next
		}
	}
}

func InsertAtIthPosition(l *LinkedList, val int, pos int) {
	node := &Node{value: val}
	if l.len == 0 {
		l.head = node
		l.len++
		return
	}

	if pos == 0 {
		tmp := l.head
		l.head = node
		l.head.next = tmp
		l.len++
	} else {
		ptr := l.head
		preNode := ptr
		i := 1

		for ptr != nil {
			if pos == i {
				preNode.next = node
				node.next = ptr
				break
			}
			preNode = ptr
			i++
			l.len++
			ptr = ptr.next
		}
	}
}

func (l *LinkedList) InsertRec(val, index int) *Node {
	head := InsertRecursion(val, index, l.head)
	l.head = head
	return l.head
}
func InsertRecursion(val, index int, node *Node) *Node {
	if index == 0 {
		newNode := &Node{value: val, next: node}
		return newNode
	}
	node.next = InsertRecursion(val, index-1, node.next)
	fmt.Println(node, "\n")
	return node
}

func deleteAtIthPosition(l *LinkedList, pos int) {
	if pos < 0 {
		fmt.Println("Pos cant be empty")
		return
	}

	if l.len == 0 {
		fmt.Println("No list in the node")
		return
	}

	if pos == 0 {
		l.head = l.head.next
		l.len--
	}

	prevNode := l.GetAt(pos - 1)
	if prevNode != nil {
		prevNode.next = prevNode.next.next
		l.len--
	}

	// ptr := l.head
	// i := 0
	// for ptr != nil {
	// 	if i == (pos - 2) {
	// 		fmt.Println(i, ptr)
	// 		ptr.next = ptr.next.next
	// 	}
	// 	ptr = ptr.next
	// 	l.len--
	// 	i++
	// }
}

func deleteAtGivenVal(l *LinkedList, val int) *Node {
	if l.len == 0 {
		fmt.Println("No list in the node")
		return &Node{}
	}

	// traverse the node and find the val of the element
	ptr := l.head
	for ptr.next != nil {
		fmt.Println(ptr.value)

		if ptr.next.value == val {
			ptr.next = ptr.next.next
		} else {
			ptr = ptr.next
		}
		l.len--
	}

	if l.head.value == val {
		fmt.Println(l.head.next)
		return l.head.next
	}
	return l.head
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// Golang Program to reverse a given linked list
func reverse(l *LinkedList) {
	var prev *Node
	ptr := l.head
	next := &Node{}
	i := 0
	for ptr != nil {
		next = ptr.next //find the next of current pointer
		ptr.next = prev // assing prev to next the current pointer
		if i == 0 {
			ptr.next = nil
		}
		prev = ptr // assign current pointer as prev of the pointer
		ptr = next // assign current poitner to next to loop through the list
		i++
	}
	l.head = prev
}

func reverseBetween(l1 *LinkedList, left int, right int) {
	var prev *Node
	var tmp *Node
	var headPointer *Node
	leftPtr := l1.head // get the left position of list
	if leftPtr.next == nil || (left == 1 && right == 1) {
		return
	}

	for i := 0; i < left-1; i++ {
		headPointer = leftPtr
		leftPtr = headPointer.next
	}

	curr := leftPtr

	for i := 0; i < (right-left)+1; i++ {
		tmp = curr.next
		fmt.Println("\n Printing tmp", tmp)
		curr.next = prev
		prev = curr
		curr = tmp
	}
	fmt.Println("\n ", headPointer, leftPtr, prev, tmp, curr)

	if headPointer != nil {
		headPointer.next = prev
	} else {
		l1.head = prev
	}

	if tmp != nil {
		leftPtr.next = tmp
	}
}

func hasCycle(l1 *LinkedList) bool {
	// Implementation of Floyd’s Cycle-Finding Algorithm
	slowPtr := l1.head
	fastPtr := l1.head
	fmt.Println("\n Print Pointers : ", slowPtr, fastPtr)
	for slowPtr != nil && fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		fmt.Println("\n Print Pointers : ", slowPtr, fastPtr)
		if slowPtr == fastPtr {
			return true
		}
	}
	return false
}

func (l1 *LinkedList) findNode(val int) *Node {
	ptr := l1.head

	for ptr != nil {
		if ptr.value == val {
			return ptr
		}
		ptr = ptr.next
	}
	return &Node{}
}

func isListPallindrome(l1 *LinkedList) bool {
	tmp := l1.head
	// l2 := l1.head
	_byteArr := []int{}
	next := &Node{}
	prev := &Node{}
	i := 0
	for tmp != nil {
		_byteArr = append(_byteArr, tmp.value)
		next = tmp.next
		tmp.next = prev
		if i == 0 {
			tmp.next = nil
		}
		prev = tmp
		tmp = next
		i++
	}
	l1.head = prev

	tmpHead := l1.head
	fmt.Println(tmpHead)
	j := 0
	for tmpHead != nil {
		fmt.Println(tmpHead.value, _byteArr[j])
		if tmpHead.value != _byteArr[j] {
			return false
		}
		tmpHead = tmpHead.next
		j++
	}
	return true
	// Or we can dot this as below
	// res := []int{}
	// for head != nil {
	//     res = append(res, head.Val)
	//     head = head.Next
	// }
	// for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
	//     if res[i] != res[j] { return false }
	// }
	// return true
}

/**
* Remove Duplicates from Sorted List
* Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
 */

func (l *LinkedList) deleteDuplicates() *LinkedList {
	ptr := l.head
	if ptr == nil {
		return l
	}
	for ptr.next != nil {
		if ptr.next.value == ptr.value {
			ptr.next = ptr.next.next
		} else {
			ptr = ptr.next
		}
	}
	fmt.Println(l.head)

	return l
}

/**
* Remove Zero Sum Consecutive Nodes from Linked List
* Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.
* After doing so, return the head of the final linked list.  You may return any such answer.
 */

func (l *LinkedList) removeZeroSumSublists() *LinkedList {
	ptr := l.head
	preNode := ptr
	if ptr == nil {
		return l
	}
	for ptr != nil {
		fmt.Println()
		fmt.Println("Cuur Pointer", ptr, preNode)
		if ptr.next != nil && (ptr.value+ptr.next.value) == 0 {
			preNode.next = ptr.next.next
		}
		preNode = ptr
		ptr = ptr.next
	}

	return l
}

/**
* Remove Zero Sum Consecutive Nodes from Linked List
* Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.
* After doing so, return the head of the final linked list.  You may return any such answer.
 */
func mergeTwoLists(l *LinkedList, l2 *LinkedList) *Node {
	head1 := l.head
	head2 := l2.head
	ans := &Node{}
	tail := ans
	for head1 != nil && head2 != nil {
		if head1.value > head2.value {
			// InsertAtTail(ans, head2.value)
			tail.next = head2
			head2 = head2.next
			tail = tail.next
		} else {
			fmt.Println()
			fmt.Println(head1, head2)
			tail.next = head1
			head1 = head1.next
			tail = tail.next
		}
	}

	if head1 != nil {
		tail.next = head1
	}
	if head2 != nil {
		tail.next = head2
	}
	return ans.next
}

func hasCycleManual(l *LinkedList) bool {
	//slow and fastpointer implementation
	slow := l.head
	fast := l.head

	for slow != nil && fast != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			fmt.Println(slow, fast)
			return true
		}
	}
	return false
}

func findlengthCycle(l *LinkedList) int {
	//slow and fastpointer implementation
	slow := l.head
	fast := l.head
	counter := 0
	for slow != nil && fast != nil {
		slow = slow.next
		fmt.Println(slow)
		fast = fast.next.next

		if slow == fast {
			break
		}
	}
	temp := slow.next
	fmt.Println(temp)
	for temp != nil && temp != slow {
		temp = temp.next
		counter++
	}
	return counter + 1
}

func detectNodeCycle(l *LinkedList) (*Node, interface{}) {
	//slow and fastpointer implementation
	slow := l.head
	fast := l.head
	counter := 0
	for slow != nil && fast != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}
	temp := slow.next

	for temp != nil && temp != slow {
		temp = temp.next
		counter++
	}

	if counter == 0 {
		return nil, nil
	}
	start := l.head
	end := l.head

	for i := 0; i < (counter + 1); i++ {
		start = start.next
	}

	for start != end {
		fmt.Println("Here", counter, start, end)
		start = start.next
		end = end.next
	}

	return end, nil
}

// https://leetcode.com/problems/happy-number/
// Google questions
func isHappy(n int) bool {
	slow := n
	fast := n

	for slow > 1 {
		slow = getSquareRoot(slow)
		fast = getSquareRoot(getSquareRoot(fast))
	}

	if slow == 1 {
		return true
	}

	return false
}

func getSquareRoot(n int) int {
	ans := 0
	for n > 0 {
		rem := n % 10
		ans += rem * rem
		n /= 10
	}
	return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * Que: Middle of the Linked List
 */
func middleNode(l1 *LinkedList) *Node {

	slow := l1.head
	fast := l1.head
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// if counter%2 == 0 {
	// 	return slow
	// }
	return slow
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Initialize two pointers, a and b, to the heads of the linked lists
	a, b := headA, headB

	// Traverse both linked lists with two pointers
	for a != b {
		// When one of the pointers reaches the end of the list, redirect it to the head of the other list
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
