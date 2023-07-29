package main

import "fmt"

type Node struct {
	Prev *Node
	Val  int
	Next *Node
}

func main() {

	head := &Node{Val: 1}
	// tail := head

	// head.InsertAtHead(3)
	// head.InsertAtHead(4)
	// head.InsertAtHead(5)

	head = head.InsertAtHead(3)
	head = head.InsertAtHead(4)
	head = head.insertAtLast(5)
	// // list = list.Prev
	// head = head.InsertAtHead(4).Prev
	// head = head.InsertAtHead(5).Prev
	// fmt.Println(tail.Prev.Prev.Prev.Val)
	TraverseLinkedList(head)
	head = head.insertAfterGivenValue(5, 6)
	TraverseLinkedList(head)
	TraverseLinkedListReverse(head)
}

func (head *Node) InsertAtHead(val int) *Node {

	newNode := &Node{Val: val}
	ptr := head

	if ptr == nil {
		ptr = newNode
		return ptr
	}

	newNode.Prev = nil
	newNode.Next = ptr
	ptr.Prev = newNode
	ptr.Next = head.Next
	// *head = *ptr.Prev
	return newNode
}

func (head *Node) insertAtLast(val int) *Node {
	newNode := &Node{Val: val}
	ptr := head
	var lastNode *Node
	if ptr == nil {
		return newNode
	}

	for ptr != nil {
		lastNode = ptr
		ptr = ptr.Next
	}
	newNode.Prev = lastNode
	lastNode.Next = newNode
	return head
}

func (head *Node) insertAfterGivenValue(nodeValue int, numberToInsert int) *Node {
	newNode := &Node{Val: numberToInsert}
	var prevPointer *Node
	if head == nil {
		return newNode
	}

	ptr := head
	for ptr != nil {
		if ptr.Val == nodeValue {
			prevPointer = ptr
			break
		}
		ptr = ptr.Next
	}

	if prevPointer == nil {
		fmt.Println("Node value does not exist")
		return prevPointer
	}

	newNode.Next = prevPointer.Next
	prevPointer.Next = newNode
	newNode.Prev = prevPointer
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
	return head
}

func TraverseLinkedList(n *Node) {
	ptr := n
	fmt.Printf("Linked List: ")
	for ptr != nil {
		fmt.Printf("%d -> ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Printf("NULL \n")
}

func TraverseLinkedListReverse(n *Node) {
	ptr := n
	var tailNode *Node
	fmt.Printf("Linked List: ")
	for ptr != nil {
		tailNode = ptr
		ptr = ptr.Next
	}
	for tailNode != nil {
		fmt.Printf("%d -> ", tailNode.Val)
		tailNode = tailNode.Prev
	}
	fmt.Printf("Start \n")
}
