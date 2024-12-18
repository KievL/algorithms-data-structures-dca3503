package main

import "fmt"

func main() {
	v := []int{6, 3, 5, 12, 7, 8, 4, 7}

	fmt.Println(selectionSort(v))
}

func selectionSort(v []int) []int {
	for i := 0; i < len(v)-1; i++ {
		menorId := i

		for j := i+1; j < len(v); j++ {
			if v[j] < v[menorId] {
				menorId = j
			}
		}

		v[i], v[menorId] = v[menorId], v[i]
	}
	return v
}
