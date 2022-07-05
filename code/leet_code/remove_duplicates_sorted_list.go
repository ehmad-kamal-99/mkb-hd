package main

// Problem Statement
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

/*
Notes

No need.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// call function from here
	// fmt.Println(test())
}

func removeDuplicatesSortedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	next := prev.Next

	for next != nil {
		if prev.Val == next.Val {
			prev.Next = next.Next

			next = next.Next
			continue
		}

		prev = prev.Next
		next = next.Next
	}

	return head
}

// Challenge:
