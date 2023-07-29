package main

import (
	"fmt"

	"github.com/imgoogege/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

// Node represents a node of linked list
//
//	type TreeNode struct {
//		Val   int
//		Left  *TreeNode
//		Right *TreeNode
//	}
// var result = []int{}

func main() {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 3}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 5}
	tree.Right.Left = &TreeNode{Val: 6}
	tree.Right.Right = &TreeNode{Val: 7}
	tree.Right.Right.Left = &TreeNode{Val: 8}
	//nums := []int{-10, -3, 0, 5, 9}
	// PreOrderTraverse(tree)
	fmt.Println(PostOrderTraverse(tree))
	// fmt.Println(printbinarytree.PrintTree(tree))
}

func PreOrderTraverse(root *TreeNode) *TreeNode {
	// Preorder travers towards Root Left Right
	if root == nil {
		return root
	}

	fmt.Println(root.Val)
	PreOrderTraverse(root.Left)
	PreOrderTraverse(root.Right)
	return root
}

func InOrderTraverse(root *TreeNode) *TreeNode {

	// Preorder travers towards Left Root Right
	if root == nil {
		return root
	}
	InOrderTraverse(root.Left)
	fmt.Println("Inorder Traversal :", root.Val)
	InOrderTraverse(root.Right)
	return root
}

func PostOrderTraverse(root *TreeNode) *TreeNode {

	// Preorder travers towards Left Root Right
	if root == nil {
		return root
	}
	PostOrderTraverse(root.Left)
	PostOrderTraverse(root.Right)
	fmt.Println("Postorder Traversal :", root.Val)
	return root
}

// Que: Given the root of a binary tree, return the preorder traversal of its nodes' values.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	// initiate int type of slice
	// Write preorder formula for traverse i.e. Root Left Right
	// Push node values to result slice
	// Write a logic to end the recursive function
	// Traverse trough all the way to left and right

	// if root == nil {
	// 	return result
	// }

	// result = append(result, root.Val)
	// preorderTraversal(root.Left)
	// preorderTraversal(root.Right)
	// return result
	result := []int{}
	return recOrderTraversal(root, &result)
}

func recOrderTraversal(root *TreeNode, result *[]int) []int {
	// initiate int type of slice
	// Write preorder formula for traverse i.e. Root Left Right
	// Push node values to result slice
	// Write a logic to end the recursive function
	// Traverse trough all the way to left and right
	if root == nil {
		return []int{}
	}
	fmt.Println(result)
	*result = append(*result, root.Val)
	recOrderTraversal(root.Left, result)
	recOrderTraversal(root.Right, result)
	return *result
}

/**
 * Problem 108. Convert Sorted Array to Binary Search Tree
 * Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced  binary search tree.
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := int(len(nums) / 2)

	root := &TreeNode{Val: nums[midIndex]}
	if midIndex != 0 {
		fmt.Println(root.Val)
		root.Left = sortedArrayToBST(nums[:midIndex])
	}

	if midIndex != len(nums)-1 {
		fmt.Println(root.Val)
		root.Right = sortedArrayToBST(nums[midIndex+1:])
	}
	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
