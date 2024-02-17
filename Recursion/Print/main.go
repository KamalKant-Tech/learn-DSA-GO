package main

import (
	"fmt"
)

// Node represents a node in an N-ary tree
type Node struct {
	Value    int
	Children []*Node
}

// PrintNaryTree prints an N-ary tree
func PrintNaryTree(root *Node, level int) {
	if root == nil {
		return
	}
	//fmt.Println("Print level", level)
	fmt.Printf("%*d\n", level, root.Value)

	for _, child := range root.Children {
		//fmt.Println("Printing Level", level)
		PrintNaryTree(child, 8/(level+1))
	}
}

// getIndentation generates indentation based on the level
func getIndentation(level int) string {
	return fmt.Sprintf("%*d", level) // Adjust the number of spaces for indentation
}

func main() {
	// Example N-ary tree
	root := &Node{
		Value: 1,
		Children: []*Node{
			{
				Value: 2,
				Children: []*Node{
					{Value: 5},
					{Value: 6},
				},
			},
			{
				Value: 3,
				Children: []*Node{
					{Value: 7},
					{Value: 8},
				},
			},
			{
				Value: 4,
				Children: []*Node{
					{Value: 9},
					{Value: 10},
				},
			},
		},
	}

	// Print the N-ary tree
	fmt.Println("N-ary Tree:")
	PrintNaryTree(root, 0)
}
