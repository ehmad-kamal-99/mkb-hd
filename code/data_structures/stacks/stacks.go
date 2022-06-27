package main

import (
	"fmt"
)

// Use stack struct with pointer to top + size defined.
// Also, try some infix/prefix/postfix problems.

type Stack []int

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() {
	fmt.Println("popping value")

	if len(*s) == 0 {
		fmt.Println("stack is empty")
		return
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	fmt.Println(val, ": popped")

	return
}

func (s *Stack) IsEmpty() {
	if len(*s) == 0 {
		fmt.Println("stack is empty")

		return
	}
}

func (s *Stack) Print() {
	i := len(*s) - 1

	for i >= 0 {
		fmt.Println((*s)[i])
		i--
	}
}

func main() {
	stack := NewStack()

	stack.IsEmpty()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	stack.Print()

	stack.Pop()

	stack.Print()
}
