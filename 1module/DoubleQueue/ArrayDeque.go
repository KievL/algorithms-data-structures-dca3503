package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Placeholder")

	deque := &ArrayDeque{}

	deque.Init(4)

	fmt.Println(deque.Size())
	fmt.Println(deque.EnqueueFront(3))

	fmt.Println(deque.EnqueueFront(2))
	fmt.Println(deque.EnqueueFront(6))
	fmt.Println(deque.EnqueueRear(1))
	fmt.Println(deque.EnqueueRear(1))
	fmt.Println(deque.Size())

	fmt.Println(deque.DequeueFront())
	fmt.Println(deque.DequeueRear())
	fmt.Println(deque.Size())
	fmt.Println(deque.EnqueueRear(1))
	fmt.Println(deque.EnqueueRear(99))
	fmt.Println(deque.DequeueRear())

}

type ArrayDeque struct {
	v        []int
	front    int
	rear     int
	inserted int
}

type Deque interface {
	Size() int
	EnqueueFront(e int) error
	DequeueFront() (int, error)
	EnqueueRear(e int) error
	DequeueRear() (int, error)
}

func (d *ArrayDeque) Init(size int) error {
	if size < 1 {
		return errors.New("Size must be > 0")
	}

	d.v = make([]int, size)

	d.front = -1
	d.rear = -1
	d.inserted = 0

	return nil
}

func (d *ArrayDeque) Size() int {
	return d.inserted
}

func (d *ArrayDeque) EnqueueFront(e int) error {
	if d.inserted == len(d.v) {
		return errors.New("Deque is full")
	}

	if d.inserted == 0 {
		d.front = 0
		d.rear = 0
		d.v[0] = e
	} else {
		new_place := (d.front - 1 + len(d.v)) % len(d.v)

		d.v[new_place] = e

		d.front = new_place
	}
	d.inserted++

	return nil
}

func (d *ArrayDeque) EnqueueRear(e int) error {
	if d.inserted == len(d.v) {
		return errors.New("Deque is full")
	}

	if d.inserted == 0 {
		d.front = 0
		d.rear = 0
		d.v[0] = e
	} else {
		new_place := (d.rear + 1) % len(d.v)

		d.v[new_place] = e

		d.rear = new_place
	}

	d.inserted++

	return nil
}

func (d *ArrayDeque) DequeueFront() (int, error) {
	if d.inserted == 0 {
		return 0, errors.New("Deque is already empty")
	}

	temp := d.v[d.front]
	if d.inserted == 1 {
		d.front = -1
		d.rear = -1
	} else {
		d.front = (d.front + 1) % len(d.v)

	}

	d.inserted--

	return temp, nil
}

func (d *ArrayDeque) DequeueRear() (int, error) {
	if d.inserted == 0 {
		return 0, errors.New("Deque is already empty")
	}

	temp := d.v[d.rear]
	if d.inserted == 1 {
		d.front = -1
		d.rear = -1
	} else {
		d.rear = (d.front - 1 + len(d.v)) % len(d.v)

	}

	d.inserted--

	return temp, nil
}
