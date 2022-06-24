package main

import (
	"fmt"
)

type Queue []int

func newQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("queue empty, nothing to dequeue")

		return
	}

	fmt.Println("dequeue value:", (*q)[0])

	*q = (*q)[1:]
}

func (q *Queue) front() {
	if q.isEmpty() {
		fmt.Println("queue empty")

		return
	}

	fmt.Println((*q)[0])
}

func (q *Queue) rear() {
	if q.isEmpty() {
		fmt.Println("queue empty")

		return
	}

	fmt.Println((*q)[len(*q)-1])
}

func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}

	return false
}

func (q *Queue) print() {
	for i := 0; i < len(*q); i++ {
		fmt.Println((*q)[i])
	}
}

func main() {
	queue := newQueue()

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)

	queue.print()

	queue.enqueue(40)

	queue.print()

	queue.dequeue()

	queue.print()

	queue.dequeue()
	queue.dequeue()
	queue.dequeue()

	queue.print()
}
