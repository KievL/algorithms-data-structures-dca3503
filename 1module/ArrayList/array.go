package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("placeholder")

	array := &ArrayList{}

	array.Init(1)

	fmt.Println(array.Size())
	array.Add(3)
	fmt.Println(array.Size())
	fmt.Println(array.Get(0))
	array.Add(4)
	array.Add(5)
	array.Add(6)
	array.AddOnIndex(9, 0)
	array.Set(7, 1)
	fmt.Println(array.Size())
	fmt.Println(array.Get(0))
	fmt.Println(array.Get(1))
	array.Remove(2)
	fmt.Println(array.Get(2))

}

// ArrayList
// Add(e)
// AddOnIndex(e, i)
// Remove(i)
// Get(i)
// Size()
// Set(e, i)

type List interface {
	Size() int
	Get(i int) (int, error)
	Add(e int)
	AddOnIndex(e int, i int) error
	Remove(i int) error
	Set(e int, i int) error
}

type ArrayList struct {
	v        []int
	inserted int
}

func (l *ArrayList) Init(size int) error {
	if size <= 0 {
		return errors.New("Size should be more than 0.")
	}

	l.v = make([]int, size)
	l.inserted = 0

	return nil
}

func (l *ArrayList) Size() int {
	return l.inserted
}

func (l *ArrayList) Get(i int) (int, error) {
	if i < l.inserted && i >= 0 {
		return l.v[i], nil
	} else {
		return 0, errors.New("Index is out of range.")
	}
}

func (l *ArrayList) doubleV() {
	newSize := 2 * len(l.v)
	newV := make([]int, newSize)

	for i := 0; i < l.inserted; i++ {
		newV[i] = l.v[i]
	}

	l.v = newV
}

func (l *ArrayList) Add(e int) {
	if l.inserted == len(l.v) {
		l.doubleV()
	}

	l.v[l.inserted] = e
	l.inserted++
}

func (l *ArrayList) AddOnIndex(e int, i int) error {
	if i >= 0 && i <= l.inserted {
		if l.inserted == len(l.v) {
			l.doubleV()
		}
		for j := l.inserted; j > i; j-- {
			l.v[j] = l.v[j-1]
		}

		l.v[i] = e
		l.inserted++
		return nil
	} else {
		return errors.New("Index i out of range.")
	}
}

func (l *ArrayList) Remove(i int) error {
	if i >= 0 && i < l.inserted {

		for j := i; j < l.inserted-1; j++ {
			l.v[j] = l.v[j+1]
		}

		l.inserted--
		return nil
	} else {
		return errors.New("Index i out of range")
	}
}

func (l *ArrayList) Set(e int, i int) error {
	if i >= 0 && i < l.inserted {
		l.v[i] = e
		return nil
	} else {
		return errors.New("Index i out of range")
	}
}
