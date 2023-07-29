package main

import "fmt"

// Node represents a node of linked list
type Node struct {
	Val  int
	Next *Node
}

func main() {
	var list *Node // null node
	list = list.insertAtLast(1)
	head := list
	head.Next = head
	head = head.insertAtLast(2)
	head = head.insertAtLast(3)
	head = head.insertAtLast(4)
	head = head.insertAtLast(5)
	head = head.insertAtLast(6)
	head.Traverse()

	head = delete(head, 1)

	head.Traverse()
}

func (head *Node) insertAtLast(val int) *Node {
	node := &Node{Val: val}
	if head == nil {
		head = node
		return head
	}
	currHead := head

	for currHead != nil {
		if currHead.Next == head {
			currHead.Next = node
			node.Next = head
			break
		}
		currHead = currHead.Next
	}
	fmt.Println(currHead.Next)
	return head
}

func delete(head *Node, val int) *Node {
	currPtr := head
	for currPtr != nil {
		if currPtr.Next == head {
			break
		}
		if currPtr.Next.Val == val {
			currPtr.Next = currPtr.Next.Next
			break
		}
		currPtr = currPtr.Next
	}

	// if the deleted value is head then shift the head to the next
	if head.Val == val {
		currPtr.Next = head.Next
		head = head.Next
	}
	return head
}

func (head *Node) Traverse() {
	ptr := head
	fmt.Printf("Linked List: ")
	for ptr != nil {
		if ptr.Next == head {
			fmt.Printf("%d -> ", ptr.Val)
			break
		}
		fmt.Printf("%d -> ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Printf("%d \n", ptr.Next.Val)
}
