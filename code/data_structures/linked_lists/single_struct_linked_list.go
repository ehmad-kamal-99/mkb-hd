package main

import (
	"errors"
	"fmt"
)

// ListNode represents a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// newListNode - initiates a new list node for linked list.
// This linked list can hold only positive numbers.
func newListNode() *ListNode {
	return &ListNode{
		Val:  -1,
		Next: nil,
	}
}

// addToBeginning - adds node to the beginning of linked list.
func (l *ListNode) addToBeginning(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: l,
	}
}

// addInOrder - add node in correct ascending order.
func (l *ListNode) addInOrder(val int) *ListNode {
	if l.Val == -1 {
		l.Val = val
		l.Next = nil

		return l
	}

	if l.Val > val {
		return l.addToBeginning(val)
	}

	if l.Val == val {
		l.append(val)

		return l
	}

	head := l

	for head != nil {
		if head.Val > val {
			// need to keep reference of previous node for this case. code refactor required. for looping, use head.Next.Next
		} else if head.Val < val {
			head = head.Next

			continue
		} else if head.Val == val {
			newListNode := &ListNode{
				Val:  val,
				Next: head.Next,
			}

			head.Next = newListNode

			break
		}

	}

	return head
}

// append - adds a value to the underlying linked list.
func (l *ListNode) append(val int) {
	if l.Val == -1 {
		l.Val = val
		l.Next = nil

		return
	}

	head := l

	for head != nil {
		head = head.Next
	}

	head = &ListNode{
		Val:  val,
		Next: nil,
	}
}

// deleteOnValue - delete node based on value.
func (l *ListNode) deleteOnValue(val int) *ListNode {
	if l.getLength() == 0 {
		return l
	}

	if l.getLength() == 1 {
		if l.Val == val {
			return l.Next
		}

		return l
	}

	if l.getLength() == 2 {
		if l.Val == val {
			return &ListNode{
				Val:  l.Next.Val,
				Next: nil,
			}
		}

		if l.Next.Val == val {
			return &ListNode{
				Val:  l.Val,
				Next: nil,
			}
		}
	}

	head := l

	for head.Next.Next != nil {
		if head.Next.Val == val {
			fmt.Println("deleting node : ", val)

			head.Next = head.Next.Next

			return head
		}

		head = head.Next
	}

	return l
}

// deleteOnPosition - delete based on position.
func (l *ListNode) deleteOnPosition(pos int) (*ListNode, error) {
	if l.getLength() < pos {
		return l, errors.New("invalid position provided")
	}

	if l.getLength() == 0 {
		return l, nil
	}

	if l.getLength() == 1 {
		if pos == 0 {
			return &ListNode{
				Val:  -1,
				Next: nil,
			}, nil
		}

		return l, errors.New("invalid position provided")
	}

	if l.getLength() == 2 {
		if pos == 0 {
			return &ListNode{
				Val:  l.Next.Val,
				Next: nil,
			}, nil
		}

		if pos == 1 {
			l.Next = nil

			return l, nil
		}
	}

	// TODO: Solvable for sure, but takes too much time and too many complications.
	// Wil comeback to this later.

	return l, nil
}

// getPosition - get node position/index based on value.
func (l *ListNode) getPosition(val int) int {
	if l.getLength() == 0 {
		return -1
	}

	if l.getLength() == 1 {
		if l.Val == val {
			return 0
		}

		return -1
	}

	head := l
	counter := 0

	for head != nil {
		if head.Val == val {
			return counter
		}

		counter++
		head = head.Next
	}

	return -1
}

// getValue - get node value based on position/index.
func (l *ListNode) getValue(pos int) int {
	if l.getLength() == 0 {
		return -1
	}

	if l.getLength() == 1 {
		if pos == 0 {
			return l.Val
		}

		return -1
	}

	if pos > l.getLength() {
		return -1
	}

	head := l

	for pos == 0 {
		head = head.Next
		pos--
	}

	return head.Val
}

func (l *ListNode) getLength() int {
	if l.Val == -1 {
		return 0
	}

	head := l
	counter := 1

	for head.Next != nil {
		counter++
	}

	return counter
}

// print - prints the underlying linked list in a single line.
func (l *ListNode) print() {
	head := l

	for head.Next != nil {
		fmt.Print(fmt.Sprintf("%d --> ", head.Val))

		head = head.Next
	}

	fmt.Println(head.Val)
}

func main() {
	list := newListNode()

	list.append(10)
	list.append(20)
	list.append(30)
	list.append(40)
	list.append(50)
	list.append(60)
	list.append(70)

	list.print()

	list = list.deleteOnValue(20)

	list.print()

	list = list.deleteOnValue(40)

	list.print()
}
