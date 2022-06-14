package main

//
// import "fmt"
//
// // Problem Statement
// // You are given the heads of two sorted linked lists list1 and list2.
// // Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
// // Return the head of the merged linked list.
//
// /*
// Notes
//
// I have head already. Let's checkNext the data/value between both head.
// Whichever one is lesser, add that item new linked list,
// Move the pointer to next node in the linked list with smaller data value.
// Continue, till both or one list is empty & then append all values to new list.
//
// Problem
// LinkedList initiated with a val of 0 & next == nil
// If a value is added for the first time, how do I verify that the list is empty or not?
// If I check that ListNode.Val == 0, then add value for example.
// In next iteration, this check is going to pass and head will be replaced again. As long as I keep adding 0, it's going to get overridden.
// And -ve number is also not an option like, if I want to set it -1 and then add 0, it will work.
// But what if my linked list can have -ve numbers?
// */
//
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
//
// func initiateListNode() *ListNode {
// 	return &ListNode{
// 		Val:  -101,
// 		Next: nil,
// 	}
// }
//
// func main() {
// 	listOne := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 4,
// 				Next: &ListNode{
// 					Val: 5,
// 					Next: &ListNode{
// 						Val:  6,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
//
// 	listTwo := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 3,
// 			Next: &ListNode{
// 				Val:  4,
// 				Next: nil,
// 			},
// 		},
// 	}
//
// 	// call function from here
// 	// fmt.Println(mergeTwoSortedLists(listOne, listTwo))
// 	newList := mergeTwoSortedLists(listOne, listTwo)
// 	newList.printAll()
// }
//
// func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	newList := initiateListNode()
//
// 	for {
// 		list1, list2 = newList.compareAndUpdate(list1, list2)
//
// 		if list1 == nil && list2 == nil {
// 			return newList
// 		}
//
// 		if list1 == nil || list2 == nil {
// 			break
// 		}
// 	}
//
// 	if list1 == nil {
// 		newList.appendToLast(list2)
// 	} else if list2 == nil {
// 		newList.appendToLast(list1)
// 	}
//
// 	return newList
// }
//
// func (l *ListNode) compareAndUpdate(nodeOne, nodeTwo *ListNode) (*ListNode, *ListNode) {
// 	if nodeOne.Val > nodeTwo.Val {
// 		l.addToList(nodeTwo.Val)
//
// 		nodeTwo = nodeTwo.Next
//
// 		return nodeOne, nodeTwo
// 	} else if nodeOne.Val < nodeTwo.Val {
// 		l.addToList(nodeOne.Val)
//
// 		nodeOne = nodeOne.Next
//
// 		return nodeOne, nodeTwo
// 	} else {
// 		l.addToList(nodeOne.Val)
// 		l.addToList(nodeTwo.Val)
//
// 		nodeOne = nodeOne.Next
// 		nodeTwo = nodeTwo.Next
//
// 		return nodeOne, nodeTwo
// 	}
// }
//
// func (l *ListNode) addToList(val int) {
// 	if l.Val == -101 {
// 		l.Val = val
// 		l.Next = nil
//
// 		return
// 	}
//
// 	head := l
//
// 	for head.Next != nil {
// 		head = head.Next
// 	}
//
// 	head.Next = &ListNode{
// 		Val:  val,
// 		Next: nil,
// 	}
// }
//
// func (l *ListNode) appendToLast(node *ListNode) {
// 	head := l
//
// 	for head.Next != nil {
// 		head = head.Next
// 	}
//
// 	head.Next = node
// }
//
// // printAll - prints the underlying linked list in a single line.
// func (l *ListNode) printAll() {
// 	head := l
//
// 	for head.Next != nil {
// 		fmt.Print(fmt.Sprintf("%d --> ", head.Val))
//
// 		head = head.Next
// 	}
//
// 	fmt.Println(head.Val)
// }
//
// // Challenge:
