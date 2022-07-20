package main

import (
	"fmt"
)

// Problem Statement
//

/*
Notes

//[1,2,2,3,null,null,3,4,null,null,4]


				1
			  /   \
			2   	2
		   /  \	   /  \
		  3    n  n	   3
		 /	\ / \
		 4	 n 	 4


*/

func main() {
	// call function from here
	treeOne := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	treeTwo := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isBalanced(treeOne))
	fmt.Println(isBalanced(treeTwo))
	fmt.Println(isBalanced(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftTree := root.Left
	rightTree := root.Right

	leftHeight := height(leftTree)
	rightHeight := height(rightTree)

	if leftHeight < rightHeight {
		leftHeight, rightHeight = rightHeight, leftHeight
	}

	if leftHeight-rightHeight > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if lh > rh {
		return lh + 1
	}

	return rh + 1
}

// Challenge:
