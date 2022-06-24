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

func (l *LinkedList) AddToBeginning(val int) {
	node := &Node{
		Val:  val,
		Next: l.Head,
	}

	l.Head = node
	l.Length++
}

func (l *LinkedList) AddInMiddle(val, pos int) {
	if pos == 1 {
		l.AddToBeginning(val)

		return
	}

	if pos > l.Length {
		l.Append(val)

		return
	}

	insertPoint := pos - 1
	pointer := l.Head

	for insertPoint != 1 {
		pointer = pointer.Next
		insertPoint--
	}

	temp := pointer.Next
	pointer.Next = &Node{
		Val:  val,
		Next: temp,
	}
	l.Length++
}

func (l *LinkedList) Append(val int) {
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
	l.Length++
}

func (l *LinkedList) DeleteHead() {
	l.Head = l.Head.Next
	l.Length--
}

func (l *LinkedList) DeleteOnValue(val int) {
	if l.Head.Val == val {
		l.DeleteHead()

		return
	}

	pointer := l.Head

	for pointer.Next.Next != nil {
		if pointer.Next.Val == val {
			fmt.Println("value found, deleting node: ", pointer.Next.Val)

			pointer.Next = pointer.Next.Next
			l.Length--

			return
		}

		pointer = pointer.Next
	}

	if pointer.Next.Val == val {
		fmt.Println("value found, deleting node: ", pointer.Next.Val)

		l.DeleteLast()

		return
	}

	fmt.Println("value:", val, "not found")
}

func (l *LinkedList) DeleteOnPosition(pos int) {
	if pos == 1 {
		fmt.Println("deleting first node:", l.Head.Val)

		l.DeleteHead()

		return
	}

	if pos > l.Length {
		fmt.Println("position does not exist")

		return
	}

	if pos == l.Length {
		fmt.Println("deleting last node")

		l.DeleteLast()

		return
	}

	pointer := l.Head
	pos = pos - 1

	for pos != 1 {
		pointer = pointer.Next
		pos--
	}

	fmt.Println("deleting node:", pointer.Next.Val)

	temp := pointer.Next.Next
	pointer.Next = temp
	l.Length--
}

func (l *LinkedList) DeleteLast() {
	pointer := l.Head

	for pointer.Next.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = nil
	l.Length--
}

func (l *LinkedList) Search(val int) {
	pointer := l.Head
	counter := 1

	for pointer.Next != nil {
		if pointer.Val == val {
			fmt.Println("node: ", pointer.Val, "found at position:", counter)

			return
		}

		pointer = pointer.Next
		counter++
	}

	if pointer.Val == val {
		fmt.Println("node: ", pointer.Val, "found at position:", counter)

		return
	}

	fmt.Println("node:", val, "not found!")
}

// Sort - sort items in the linked list.
// Start one loop with pointer at head.
// Start another loop with pointer at head.next
// Compare both and swap if need be.
// Move second pointer till linked list empty.
// Then repeat the process.
func (l *LinkedList) Sort() {
	ptr := l.Head

	for i := 0; i < l.Length; i++ {
		loopPtr := ptr

		for loopPtr.Next != nil {
			if loopPtr.Val > loopPtr.Next.Val {
				temp := loopPtr.Next.Val
				loopPtr.Next.Val = loopPtr.Val
				loopPtr.Val = temp
			}

			loopPtr = loopPtr.Next
		}
	}
}

func (l *LinkedList) Print() {
	firstNode := l.Head

	fmt.Print(firstNode.Val)
	firstNode = firstNode.Next

	for firstNode != nil {
		fmt.Print(fmt.Sprintf("--> %d", firstNode.Val))
		firstNode = firstNode.Next
	}

	fmt.Println()
}

func main() {
	newList := NewLinkedList()

	newList.Append(20)
	newList.Append(30)
	newList.Append(40)
	newList.Append(50)

	newList.Print()

	newList.AddToBeginning(10)

	newList.Print()

	newList.AddInMiddle(25, 3)
	newList.AddInMiddle(0, 1)   // 1 indicates first position or head. Same as AddToBeginning.
	newList.AddInMiddle(60, 10) // if position > linkedList.length, insert at end
	newList.AddInMiddle(55, 8)
	newList.AddInMiddle(70, 10)

	newList.Print()

	newList.DeleteHead()

	newList.Print()

	newList.DeleteLast()

	newList.Print()

	newList.DeleteOnValue(55)

	newList.Print()

	newList.DeleteOnValue(60)

	newList.Print()

	newList.DeleteOnValue(20)

	newList.Print()

	newList.DeleteOnValue(99) // not found.

	newList.Print()

	newList.DeleteOnPosition(3)
	newList.Print()
	newList.DeleteOnPosition(3)
	newList.Print()
	newList.DeleteOnPosition(3)
	newList.Print()
	newList.DeleteOnPosition(3)
	newList.Print()

	newList.AddInMiddle(250, 100)
	newList.Print()

	newList.Search(10)
	newList.Print()
	newList.Search(25)
	newList.Print()
	newList.Search(250)
	newList.Print()
	newList.Search(260) // not found
	newList.Print()

	newList.Append(1)
	newList.Append(100)
	newList.Append(66)
	newList.Append(30)
	newList.Append(45)
	newList.Append(15)
	newList.Print()

	newList.Sort()
	newList.Print()
}
