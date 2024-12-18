package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := &LinkedList{}
	stack.Init()

	fmt.Println(stack.Size())
	fmt.Println(stack.Top())

	stack.Push(1)
	stack.Push(7)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Size())
	fmt.Println(stack.Top())

	stack.Pop()
	fmt.Println(stack.Size())
	fmt.Println(stack.Top())

	stack.Pop()
	fmt.Println(stack.Size())
	fmt.Println(stack.Top())
}

type Stack interface {
	Size() int
	Push(e int)
	Pop() (int, error)
	Top() (int, error)
}

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head     *Node
	inserted int
}

func (l *LinkedList) Init() {
	l.head = &Node{val: 0, next: nil}
	l.inserted = 0
}

func (l *LinkedList) Size() int {
	return l.inserted
}

func (l *LinkedList) Push(e int) {
	newNode := &Node{val: e, next: l.head}
	l.head = newNode
	l.inserted++
}

func (l *LinkedList) Pop() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Stack is empty")
	}

	popped := l.head

	l.head = l.head.next
	l.inserted--

	return popped.val, nil
}

func (l *LinkedList) Top() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Stack is empty")
	}

	return l.head.val, nil
}
