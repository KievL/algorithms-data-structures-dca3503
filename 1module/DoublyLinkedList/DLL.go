package main

import (
	"errors"
	"fmt"
)

func main() {
	dll := &DoublyLinkedList{}
	dll.Init()

	fmt.Println(dll.Size())

	dll.Add(1)
	dll.Add(2)
	dll.Add(3)
	dll.Add(4)
	dll.AddOnIndex(66, 1)
	dll.Remove(1)

	fmt.Println(dll.Size())
	fmt.Println(dll.Get(0))
	fmt.Println(dll.Get(1))
	fmt.Println(dll.Get(2))
	fmt.Println(dll.Get(3))
	//fmt.Println(dll.Get(4))

}

type List interface {
	Size() int
	Get(i int) (int, error)
	Add(e int)
	AddOnIndex(e int, i int) (int, error)
	Remove(i int) error
}

type Node struct {
	val  int
	next *Node
	back *Node
}

type DoublyLinkedList struct {
	inserted int
	head     *Node
	tail     *Node
}

func (l *DoublyLinkedList) Init() {
	l.inserted = 0
	l.head = nil
	l.tail = nil
}

func (l *DoublyLinkedList) Size() int {
	return l.inserted
}

func (l *DoublyLinkedList) Get(i int) (int, error) {
	if i >= 0 && i < l.inserted {
		current := l.head

		for j := 0; j < i; j++ {
			current = current.next
		}

		return current.val, nil
	} else {
		return 0, errors.New("Index i is out of range")
	}
}

func (l *DoublyLinkedList) Add(e int) {
	// newNode := &Node{val: e, back: l.tail, next: nil}
	// if l.inserted == 0 {
	// 	l.head = newNode

	// } else {
	// 	l.tail.next = newNode
	// }

	// l.tail = newNode

	// l.inserted++

	_ = l.AddOnIndex(e, l.inserted)
}

func (l *DoublyLinkedList) AddOnIndex(e int, i int) error {
	if i >= 0 && i <= l.inserted {
		newNode := &Node{val: e, back: nil, next: nil}

		if l.inserted == 0 {
			l.head = newNode
			l.tail = newNode
			l.inserted++

			return nil
		}

		if i == 0 {
			newNode.next = l.head
			l.head.back = newNode
			l.head = newNode

		} else if i == l.inserted {
			newNode.back = l.tail
			l.tail.next = newNode
			l.tail = newNode
		} else {
			current := l.head

			for j := 0; j < i-1; j++ {
				current = current.next
			}

			newNode.back = current
			newNode.next = current.next
			current.next = newNode
			newNode.next.back = newNode
		}

		l.inserted++
		return nil
	} else {
		return errors.New("Index i is out of range.")
	}
}

func (l *DoublyLinkedList) Remove(i int) error {
	if i >= 0 && i < l.inserted {

		if l.inserted == 1 {
			l.head = nil
			l.tail = nil
			l.inserted--
			return nil
		}

		if i == 0 {
			l.head.back = nil
		} else if i == l.inserted-1 {
			l.tail.next = nil
		} else {
			current := l.head

			for j := 0; j < i; j++ {
				current = current.next
			}

			current.back.next = current.next
			current.next.back = current.back
		}

		l.inserted--
		return nil
	} else {
		return errors.New("Index i is out of range")
	}
}
