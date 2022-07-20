package main

import (
	"fmt"
)

// Problem Statement
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.

/*
Notes



*/

func main() {
	listOne := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	}

	// call function from here
	fmt.Println(hasCycle(listOne))
	fmt.Println(hasCycleOptimized(listOne))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	validator := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := validator[head]; !ok {
			validator[head] = true
			head = head.Next
		} else {
			return true
		}
	}

	return false
}

func hasCycleOptimized(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != nil || fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next

		if fast != nil {
			fast = fast.Next

			if fast != nil {
				fast = fast.Next
			}
		}
	}

	return false
}

// Challenge: Do it with O(1) Space Complexity [Slow + Fast Pointers]
