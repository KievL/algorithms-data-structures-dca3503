package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Placeholder")

	test := &LinkedListQueue{}
	test.Init()
	fmt.Println(test.Size())
	test.Enqueue(123)
	test.Enqueue(122)
	test.Enqueue(16)
	fmt.Println(test.Size())
	fmt.Println(test.Front())
	fmt.Println(test.Dequeue())
	fmt.Println(test.Front())
	fmt.Println(test.Size())
}

type Queue interface {
	Size() int
	Enqueue(e int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
}

type Node struct {
	val  int
	next *Node
}

type LinkedListQueue struct {
	front    *Node
	rear     *Node
	inserted int
}

func (l *LinkedListQueue) Init() {
	l.front = nil
	l.rear = nil
	l.inserted = 0
}

func (l *LinkedListQueue) Size() int {
	return l.inserted
}

func (l *LinkedListQueue) Enqueue(e int) {
	newNode := &Node{val: e, next: nil}

	if l.inserted == 0 {
		l.front = newNode

	} else {
		l.rear.next = newNode
	}

	l.rear = newNode

	l.inserted++

}

func (l *LinkedListQueue) Dequeue() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Queue is empty")
	}

	val := l.front.val

	if l.inserted == 1 {
		l.rear = nil
	}

	l.front = l.front.next
	l.inserted--

	return val, nil

}

func (l *LinkedListQueue) Front() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Queue is empty")
	} else {
		return l.front.val, nil
	}
}

func (l *LinkedListQueue) IsEmpty() bool {
	if l.inserted == 0 {
		return true
	} else {
		return false
	}
}
