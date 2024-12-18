package main

import (
	"fmt"
)

func LinearSearch(v []int, e int) int {
	return -1
}

func BinarySearch(v []int, e int) int {
	ini := 0
	fim := len(v) - 1
	for ini <= fim {
		meio := (ini + fim) / 2
		if v[meio] == e { //encontrei
			return meio
		} else {
			if e > v[meio] { //estah a esquerda
				fim = meio - 1
			} else { //estah a direita
				ini = meio + 1
			}
		}
	}
	return -1
}

func BinarySearchRec(v []int, e int, ini int, fim int) int {
	if ini > fim {
		return -1
	}
	meio := (ini + fim) / 2
	if v[meio] == e {
		return meio
	} else {
		if e > v[meio] {
			fim = meio - 1
		} else {
			ini = meio + 1
		}
		return BinarySearchRec(v, e, ini, fim)
	}
}

func main() {
	v := []int{101, 100, 93, 92, 48, 41, 32, 31, 30, 20}
	val := BinarySearch(v, 20)
	fmt.Println(val)
	val = BinarySearch(v, 101)
	fmt.Println(val)
	val = BinarySearch(v, 41)
	fmt.Println(val)
	val = BinarySearch(v, 0)
	fmt.Println(val)
	val = BinarySearch(v, 1000)
	fmt.Println(val)
}
