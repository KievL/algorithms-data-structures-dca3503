package main

import (
	"fmt"
)

func main() {
	v := []int{1}

	fmt.Println(binarySearch(v, 1))
}

func binarySearch(v []int, e int) int {
	ini := 0
	fim := len(v) - 1

	for ini <= fim {
		meio := int((ini + fim) / 2)

		if e == v[meio] {
			return meio
		} else if e < v[meio] {
			fim = meio - 1

		} else {
			ini = meio + 1
		}

	}

	return -1
}
