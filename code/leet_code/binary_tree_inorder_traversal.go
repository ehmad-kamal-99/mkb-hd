package main

import (
	"fmt"
)

// Problem Statement
// Given the root of a binary tree, return the inorder traversal of its nodes' values.

/*
Notes



*/

func main() {
	// call function from here
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	inOrderTraversalR(root, &result)

	return result
}

func inOrderTraversalR(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inOrderTraversalR(root.Left, result)
	*result = append(*result, root.Val)
	inOrderTraversalR(root.Right, result)
}

//func iterative(root *TreeNode, result []int) {
//	stack := make([]int, 0)
//
//}

// Challenge: Do iterative solution.
