package main

import (
	"errors"
	"fmt"
)

func main() {
	deque := &ArrayDeque{}
	deque.Init(5)

	fmt.Println(deque.Size())
	fmt.Println(deque.IsEmpty())

	deque.EnqueueFront(1)

	fmt.Println(deque.Size())
	fmt.Println(deque.IsEmpty())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	deque.EnqueueFront(2)

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	deque.EnqueueFront(3)

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	deque.EnqueueRear(4)

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	deque.EnqueueFront(5)

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	fmt.Println(deque.DequeueRear())

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	fmt.Println(deque.DequeueFront())

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())

	deque.EnqueueRear(7)

	fmt.Println(deque.Size())
	fmt.Println(deque.Front())
	fmt.Println(deque.Rear())
}

type Deque interface {
	EnqueueFront(e int) error
	DequeueFront() error
	EnqueueRear(e int) error
	DequeueRear() error
	Front() (int, error)
	Rear() (int, error)
	Size() int
	IsEmpty() bool
}

type ArrayDeque struct {
	front int
	rear  int
	v     []int
}

func (l *ArrayDeque) Init(size int) error {
	if size < 1 {
		return errors.New("Size must be higher than 0.")
	}
	l.front = -1
	l.rear = -1
	l.v = make([]int, size)

	return nil
}

func (l *ArrayDeque) Size() int {
	if l.front == -1 {
		return 0
	}

	return (l.rear-l.front+len(l.v))%len(l.v) + 1
}

func (l *ArrayDeque) EnqueueFront(e int) error {
	if len(l.v) == l.Size() {
		return errors.New("Deque is full")
	}

	if l.Size() == 0 {
		l.front++
		l.rear++
	} else {
		l.front = (l.front - 1 + len(l.v)) % len(l.v)
	}

	l.v[l.front] = e

	return nil
}

func (l *ArrayDeque) EnqueueRear(e int) error {
	if len(l.v) == l.Size() {
		return errors.New("Deque is full")
	}

	if l.Size() == 0 {
		l.front++
		l.rear++
	} else {
		l.rear = (l.rear + 1) % len(l.v)
	}

	l.v[l.rear] = e

	return nil
}

func (l *ArrayDeque) DequeueFront() (int, error) {
	if 0 == l.Size() {
		return 0, errors.New("Deque is empty")
	}

	val := l.v[l.front]
	if l.Size() == 1 {
		l.front = -1
		l.rear = -1
	} else {
		l.front = (l.front + 1) % len(l.v)
	}

	return val, nil
}

func (l *ArrayDeque) DequeueRear() (int, error) {
	if 0 == l.Size() {
		return 0, errors.New("Deque is empty")
	}

	val := l.v[l.rear]
	if l.Size() == 1 {
		l.front = -1
		l.rear = -1
	} else {
		l.rear = (l.rear - 1 + len(l.v)) % len(l.v)
	}

	return val, nil
}

func (l *ArrayDeque) Front() (int, error) {
	if 0 == l.Size() {
		return 0, errors.New("Deque is empty")
	}

	return l.v[l.front], nil
}

func (l *ArrayDeque) Rear() (int, error) {
	if 0 == l.Size() {
		return 0, errors.New("Deque is empty")
	}

	return l.v[l.rear], nil
}

func (l *ArrayDeque) IsEmpty() bool {
	if l.Size() == 0 {
		return true
	} else {
		return false
	}
}
