package main

import (
	"fmt"
)

/*

Some imp. properties

- Maximum number of node at level `l` of BT is 2^l. level of root node = 0. [2^0 = 1]
- Maximum number of nodes in a binary tree of height ‘h’ is (2^h) – 1. height of root node = 1. [2^1 - 1 = 1]
- In Binary tree where every node has 0 or 2 children, the number of leaf nodes is always one more than nodes with
  two children.
- In a non-empty binary tree, if n is the total number of nodes and e is the total number of edges, then e = n-1.
*/

/*

TODO: Calculate diameter of trees in O(n^2) & O(n).

*/

type node struct {
	data   int
	lChild *node
	rChild *node
}

func newBinaryTree(data int) *node {
	return &node{
		data:   data,
		lChild: nil,
		rChild: nil,
	}
}

/*

Basic Operations on Binary Tree

			1
		  /	  \
	  	 2	   3
	    /  \  /  \
       4    5 6	  7
           / \     \
		  8   9     10
				   /
				 11

1- Insert an element (In level order [first available position])
2- Remove an element (Replace with deepest right-most node)
3- Search for an element
4- Depth-First Traversal
	a- PreOrder Traversal [root, left, right] (Done)
	b- InOrder Traversal [left, root, right]
	c- PostOrder Traversal [left, right, root] (Done)
5- Breadth-First Traversal [Level Order Traversal] (Done)

Auxiliary Ops

1- Find Height of Tree
2- Find Level of Tree
3- Find Size of Tree

*/

// Can't be done recursively.
// levelOrderTraversalQueue - breadth-first traversal using queue.
func levelOrderTraversalQueue(root *node) {
	queue := make([]*node, 0)

	nodePtr := root

	for nodePtr != nil {
		fmt.Print(nodePtr.data)

		if nodePtr.lChild != nil {
			queue = append(queue, nodePtr.lChild)
		}

		if nodePtr.rChild != nil {
			queue = append(queue, nodePtr.rChild)
		}

		if len(queue) == 0 {
			break
		}

		nodePtr = queue[0]
		queue = queue[1:]
	}

	fmt.Println()
}

func preOrderTraversal(root *node) {
	stack := make([]*node, 0)

	nodePtr := root

	for nodePtr != nil {
		fmt.Print(nodePtr.data)

		if nodePtr.rChild != nil {
			stack = append(stack, nodePtr.rChild)
		}

		if nodePtr.lChild != nil {
			stack = append(stack, nodePtr.lChild)
		}

		if len(stack) == 0 {
			break
		}

		nodePtr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	fmt.Println()
}

func preOrderTraversalR(root *node) {
	if root == nil {
		return
	}

	fmt.Print(root.data)

	preOrderTraversalR(root.lChild)
	preOrderTraversalR(root.rChild)
}

func inOrderTraversalR(root *node) {
	if root == nil {
		return
	}

	inOrderTraversalR(root.lChild)

	fmt.Print(root.data)

	inOrderTraversalR(root.rChild)
}

func postOrderTraversalR(root *node) {
	if root == nil {
		return
	}

	postOrderTraversalR(root.lChild)

	postOrderTraversalR(root.rChild)

	fmt.Print(root.data)
}

//func postOrderTraversal(root *node) {
//	stack := make([]*node, 0)
//
//	nodePtr := root
//
//	for nodePtr != nil {
//		if nodePtr.lChild != nil {
//			stack = append(stack, nodePtr.lChild)
//		}
//
//		if nodePtr.rChild != nil {
//			stack = append(stack, nodePtr.rChild)
//		}
//
//		if len(stack) == 0 {
//			break
//		}
//
//		nodePtr = stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//	}
//
//	fmt.Println()
//}

func printStack(stack []*node) {
	for _, node := range stack {
		fmt.Print(node, " ")
	}
	fmt.Println()
}

func heightRecursive(root *node) int {
	if root == nil {
		return 0
	} else {
		lHeight := heightRecursive(root.lChild)
		rHeight := heightRecursive(root.rChild)

		if lHeight > rHeight {
			return lHeight + 1
		}

		return rHeight + 1
	}
}

func main() {
	tree := newBinaryTree(1)
	tree.lChild = newBinaryTree(2)
	tree.rChild = newBinaryTree(3)
	tree.lChild.lChild = newBinaryTree(4)
	tree.lChild.rChild = newBinaryTree(5)
	tree.rChild.lChild = newBinaryTree(6)
	tree.rChild.rChild = newBinaryTree(7)
	tree.lChild.rChild.lChild = newBinaryTree(8)
	tree.lChild.rChild.rChild = newBinaryTree(9)
	tree.rChild.rChild.rChild = newBinaryTree(10)
	tree.rChild.rChild.rChild.lChild = newBinaryTree(11)

	tree2 := newBinaryTree(1)
	tree2.lChild = newBinaryTree(2)
	tree2.rChild = newBinaryTree(3)
	tree2.lChild.lChild = newBinaryTree(4)
	tree2.lChild.rChild = newBinaryTree(5)
	//tree2.rChild.lChild = newBinaryTree(6)
	tree2.rChild.rChild = newBinaryTree(8)
	tree2.lChild.rChild.lChild = newBinaryTree(6)
	tree2.lChild.rChild.rChild = newBinaryTree(7)
	tree2.rChild.rChild.lChild = newBinaryTree(9)
	//tree2.rChild.rChild.rChild.lChild = newBinaryTree(11)

	// fmt.Println(heightRecursive(tree))

	levelOrderTraversalQueue(tree)
	preOrderTraversal(tree)
	preOrderTraversalR(tree)
	fmt.Println()
	postOrderTraversalR(tree)
	fmt.Println()
	inOrderTraversalR(tree)
	fmt.Println()

	fmt.Println("--------")
	fmt.Println()
	levelOrderTraversalQueue(tree2)
	preOrderTraversalR(tree2)
	fmt.Println()
	inOrderTraversalR(tree2)
	fmt.Println()
	postOrderTraversalR(tree2)
	fmt.Println()
}
