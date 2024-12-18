package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Placeholder")

	ll := &LinkedList{}
	ll.Init()
	fmt.Println(ll.Size())
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.AddOnIndex(99, 1)
	ll.Remove(3)
	fmt.Println(ll.Size())
	fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(1))
	fmt.Println(ll.Get(2))

}

type List interface {
	Size() int
	Get(i int) (int, error)
	Add(e int)
	AddOnIndex(e int, i int) (int, error)
	Remove(i int) error
	Set(e int, i int) error
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
	l.head = nil
	l.inserted = 0
}

func (l *LinkedList) Size() int {
	return l.inserted
}

func (l *LinkedList) Get(i int) (int, error) {
	if i >= 0 && i < l.inserted {
		current := l.head
		for j := 0; j < i; j++ {
			current = current.next
		}

		return current.val, nil
	} else {
		return 0, errors.New("Index i is out of range.")
	}
}

func (l *LinkedList) AddOnIndex(e int, i int) error {
	if i >= 0 && i <= l.inserted {
		newNode := &Node{val: e, next: nil}

		if l.head == nil {
			l.head = newNode
			l.inserted++
			return nil
		}

		current := l.head

		for j := 0; j < i-1; j++ {
			current = current.next
		}

		newNode.next = current.next
		current.next = newNode
		l.inserted++
		return nil
	} else {
		return errors.New("Index i is out or range.")
	}
}

func (l *LinkedList) Add(e int) {

	// newNode := &Node{val: e, next: nil}

	// if l.head == nil {
	// 	l.head = newNode
	// 	l.inserted++
	// 	return
	// }
	// current := l.head

	// for i := 0; i < l.inserted-1; i++ {
	// 	current = current.next
	// }

	// current.next = newNode
	// l.inserted++

	l.AddOnIndex(e, l.inserted)
}

func (l *LinkedList) Remove(i int) error {
	if i >= 0 && i < l.inserted {

		if i == 0 {
			l.head = l.head.next
			l.inserted--
			return nil
		}

		current := l.head

		for j := 0; j < i-1; j++ {
			current = current.next
		}

		current.next = current.next.next
		l.inserted--
		return nil

	} else {
		return errors.New("Index i is out of range.")
	}
}
