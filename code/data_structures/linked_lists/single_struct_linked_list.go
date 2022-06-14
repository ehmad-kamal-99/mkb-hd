package main

import "fmt"

// ListNode represents a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// newListNode - initiates a new list node for linked list.
func newListNode() *ListNode {
	return &ListNode{
		Val:  -1,
		Next: nil,
	}
}

// add - adds a value to the underlying linked list.
func (l *ListNode) add(val int) {
	if l.Val == -1 {
		l.Val = val
		l.Next = nil

		return
	}

	head := l

	for head.Next != nil {
		head = head.Next
	}

	head.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (l *ListNode) delete(val int) *ListNode {
	if l.Val == val {
		return l.Next
	}

	head := l

	for head.Next.Next != nil {
		if head.Next.Val == val {
			fmt.Println("deleting node : ", val)

			head.Next = head.Next.Next

			return l
		}

		head = head.Next
	}

	return l
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

	list.add(10)
	list.add(20)
	list.add(30)
	list.add(40)
	list.add(50)
	list.add(60)
	list.add(70)

	list.print()

	list = list.delete(20)

	list.print()

	list = list.delete(40)

	list.print()
}
