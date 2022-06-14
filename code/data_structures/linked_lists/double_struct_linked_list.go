package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Length int
	Head   *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Length: 0,
		Head:   nil,
	}
}

func (l *LinkedList) Add(val int) {
	if l.Head == nil {
		l.Length++
		l.Head = &Node{
			Val:  val,
			Next: nil,
		}

		return
	}

	currentNode := l.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = &Node{
		Val:  val,
		Next: nil,
	}
}

func (l *LinkedList) prettyPrint() {
	firstNode := l.Head

	for firstNode.Next != nil {
		fmt.Println(fmt.Sprintf("node value: %d, next node value: %d", firstNode.Val, firstNode.Next.Val))
		firstNode = firstNode.Next
	}
}

func main() {
	newList := NewLinkedList()

	newList.Add(10)
	newList.Add(20)
	newList.Add(30)
	newList.Add(40)
	newList.Add(50)

	newList.prettyPrint()
}
