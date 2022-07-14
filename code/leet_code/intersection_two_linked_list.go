package main

// Problem Statement
// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
// If the two linked lists have no intersection at all, return null.

/*
Notes

I cheated only for Time O(m+n) & Space O(1) but found mostly O(m+n+k) where k is iteration took to reach intersection.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// call function from here
	// fmt.Println(test())
}

func intersectionTwoLinkedLists(headA, headB *ListNode) *ListNode {
	lenOne, lenTwo := getLength(headA, headB)

	if lenOne > lenTwo {
		for lenOne-lenTwo != 0 {
			lenOne--
			headA = headA.Next
		}
	} else if lenOne < lenTwo {
		for lenTwo-lenOne != 0 {
			lenTwo--
			headB = headB.Next
		}
	}

	iterA, iterB := headA, headB

	for i := 0; i < lenOne; i++ {
		if iterA == iterB {
			return iterA
		}

		iterA = iterA.Next
		iterB = iterB.Next
	}

	return nil
}

func getLength(listOne, listTwo *ListNode) (int, int) {
	i, j := 0, 0
	headOne, headTwo := listOne, listTwo

	for {
		if headOne != nil {
			i++
			headOne = headOne.Next
		}

		if headTwo != nil {
			j++
			headTwo = headTwo.Next
		}

		if headOne == nil && headTwo == nil {
			break
		}
	}

	return i, j
}

// Challenge:
