package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Olá")

	dllist := new(DoubleLinkedList)

	dllist.Add(33)
	dllist.Add(35)

	fmt.Println(dllist.Get(0))
	fmt.Println(dllist.Remove(0))
	fmt.Println(dllist.Get(0))
}

type Node struct {
	element   int
	next_node *Node
	prev_node *Node
}

type DoubleLinkedList struct {
	head     *Node
	tail     *Node
	inserted int
}

func (l *DoubleLinkedList) Add(element_ int) error {
	node := new(Node)

	node.element = element_

	if l.inserted == 0 {
		l.head = node
	} else {
		node.prev_node = l.tail
		l.tail.next_node = node
	}

	l.tail = node
	l.inserted++

	return nil
}

func (l *DoubleLinkedList) Get(index int) (int, error) {
	if index > l.inserted-1 || index < 0 {
		return -1, errors.New("Out of index.")
	}

	node_ := l.head

	for i := 1; i <= index; i++ {
		node_ = node_.next_node
	}

	return node_.element, nil
}

func (l *DoubleLinkedList) Remove(index int) error {
	if index > l.inserted-1 || index < 0 {
		return errors.New("Out of index.")
	}

	node_ := l.head

	for i := 1; i <= index; i++ {
		node_ = node_.next_node
	}

	if l.inserted == 1 {
		l.head = nil
		l.tail = nil
	} else {
		if node_ == l.head {
			l.head = node_.next_node
			node_.next_node.prev_node = nil
		} else if node_ == l.tail {
			l.tail = node_.prev_node
			node_.prev_node.next_node = nil
		} else {
			node_.next_node.prev_node = node_.prev_node
			node_.prev_node.next_node = node_.next_node
		}
	}
	l.inserted--
	return nil
}
