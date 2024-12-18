package main

import "fmt"

func main() {
	v := []int{6, 3, 5, 12, 7, 8, 4, 7}

	fmt.Println(insertion(v))
}

func insertion(v []int) []int {
	for i := 1; i < len(v); i++ {
		key := v[i]

		j := i - 1
		for j >= 0 && key < v[j] {
			v[j+1] = v[j]

			j--
		}

		v[j+1] = key
	}

	return v
}

// func insertion(v []int) []int {
// 	index := 1
// 	for index < len(v) {
// 		key := v[index]
// 		j := index - 1

// 		for j >= 0 && key < v[j] {
// 			v[j+1] = v[j]
// 			j--
// 		}

// 		v[j+1] = key
// 		index++
// 	}
// 	return v
// }
