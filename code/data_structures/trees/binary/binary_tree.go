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

1- Insert an element (In level order [first available position])
2- Remove an element (Replace with deepest right-most node)
3- Search for an element
4- Traversing an element
	a- PreOrder Traversal [root, left, right]
	b- InOrder Traversal [left, root, right]
	c- PostOrder Traversal [left, right, root]

Auxiliary Ops

1- Find Height of Tree
2- Find Level of Tree
3- Find Size of Tree

*/

func levelOrderTraversalQueue(root *node) {
	queue := make([]*node, 0)

	nodePtr := root

	for nodePtr != nil {
		fmt.Print(nodePtr.data)

		queue = append(queue, nodePtr.lChild)
		queue = append(queue, nodePtr.rChild)

		nodePtr = queue[0]
		queue = queue[1:]
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

	// fmt.Println(heightRecursive(tree))

	levelOrderTraversalQueue(tree)
}
