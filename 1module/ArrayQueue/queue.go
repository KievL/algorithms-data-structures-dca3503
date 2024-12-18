package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	q := &ArrayList{}
	q.Init(6)
	//q := &LinkedListQueue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	val, _ := q.Dequeue()
	fmt.Println(val)
	val, _ = q.Dequeue()
	fmt.Println(val)
	val, _ = q.Dequeue()
	fmt.Println(val)
	val, _ = q.Front()
	fmt.Println(val)
	q.Enqueue(7)
	q.Enqueue(8)
	val, _ = q.Dequeue()
	fmt.Println(val)
	val, _ = q.Dequeue()
	fmt.Println(val)
	val, _ = q.Dequeue()
	fmt.Println(val)
	val, _ = q.Front()
	fmt.Println(val)
	val, _ = q.Dequeue()
	fmt.Println(val)
	val, _ = q.Dequeue()
	fmt.Println(val)
	fmt.Println(q.Size())
}

type Queue interface {
	Enqueue(e int) error
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayList struct {
	v        []int
	front    int
	rear     int
	inserted int
}

func (l *ArrayList) Init(size int) error {
	if size < 1 {
		return errors.New("Size must be higher of equal to 1.")
	}

	l.v = make([]int, size)
	l.front = 0
	l.rear = 0
	l.inserted = 0
	return nil
}

func (l *ArrayList) Enqueue(e int) error {
	if l.inserted == len(l.v) {
		return errors.New("Queue is full")
	}

	if l.inserted == 0 {
		l.front++
		l.rear++
	} else {
		l.rear = (l.rear + 1) % len(l.v)
	}

	l.v[l.rear] = e
	l.inserted++

	return nil

}

func (l *ArrayList) Dequeue() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Queue is empty")
	}

	val := l.v[l.front]
	if l.inserted == 1 {
		l.front = -1
		l.rear = -1
	} else {
		l.front = (l.front + 1) % len(l.v)
	}

	l.inserted--
	return val, nil
}

func (l *ArrayList) Front() (int, error) {
	if l.inserted == 0 {
		return 0, errors.New("Queue is empty")
	}

	return l.v[l.front], nil
}

func (l *ArrayList) IsEmpty() bool {
	if l.inserted == 0 {
		return true
	} else {
		return false
	}
}

func (l *ArrayList) Size() int {
	if l.front == -1 {
		return 0
	}
	return int(math.Abs(float64(l.rear)-float64(l.front)) + 1)
}
