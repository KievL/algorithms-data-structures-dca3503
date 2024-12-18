package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := &ArrayList{}
	stack.Init()

	fmt.Println(stack.Size())
	fmt.Println(stack.Top())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Size())
	fmt.Println(stack.Top())

	stack.Pop()
	fmt.Println(stack.Size())
	fmt.Println(stack.Top())
}

type ArrayList struct {
	v        []int
	inserted int
}

type Stack interface {
	Size() int
	Push(e int)
	Pop() (int, error)
	Top() (int, error)
}

func (l *ArrayList) Init() {
	l.v = make([]int, 1)
	l.inserted = 0
}

func (l *ArrayList) Size() int {
	return l.inserted
}

func (l *ArrayList) doubleV() {
	newArray := make([]int, 2*len(l.v))

	for i := 0; i < l.inserted; i++ {
		newArray[i] = l.v[i]
	}

	l.v = newArray
}

func (l *ArrayList) Push(e int) {
	if l.inserted == len(l.v) {
		l.doubleV()
	}

	l.v[l.inserted] = e

	l.inserted++
}

func (l *ArrayList) Pop() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Stack is empty")
	}

	l.inserted--

	return l.v[l.inserted], nil
}

func (l *ArrayList) Top() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Stack is empty")
	}
	return l.v[l.inserted-1], nil
}
