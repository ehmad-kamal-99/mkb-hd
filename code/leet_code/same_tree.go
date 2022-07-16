package main

import (
	"fmt"
)

// Problem Statement
// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

/*
Notes

Do some sort of traversal and store it in an array.
Then compare the arrays.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// call function from here
	treeOne := &TreeNode{
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
	treeTwo := &TreeNode{
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
	fmt.Println(isSameTree(treeOne, treeTwo))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	treeOne := make([]int, 0)
	treeTwo := make([]int, 0)

	traverseTree(p, &treeOne)
	traverseTree(q, &treeTwo)

	if len(treeOne) != len(treeTwo) {
		return false
	}

	for i := 0; i < len(treeOne); i++ {
		if treeOne[i] != treeTwo[i] {
			return false
		}
	}

	return true
}

func traverseTree(node *TreeNode, tree *[]int) {
	if node == nil {
		return
	}

	traverseTree(node.Left, tree)
	traverseTree(node.Right, tree)

	*tree = append(*tree, node.Val)
}

// Challenge:
