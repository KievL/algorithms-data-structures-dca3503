package main

import (
	"fmt"
)

func main() {
	v := []int{6, 3, 5, 12, 7, 8, 4, 7}

	fmt.Println(bubbleSort(v))
}

func bubbleSort(v []int) []int {
	for i := 0; i < len(v)-1; i++ {
		trocou := false
		for j := 0; j < len(v)-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
				trocou = true
			}
		}
		if trocou == false {
			return v
		}
	}
	return v
}

// func bubbleSort(v []int) []int {
// 	for i := 0; i < len(v); i++ {
// 		trocou := false
// 		for j := 0; j < len(v)-i-1; j++ {
// 			if v[j] > v[j+1] {
// 				v[j], v[j+1] = v[j+1], v[j]
// 				trocou = true
// 			}
// 		}
// 		if trocou == false {
// 			return v
// 		}
// 	}
// 	return v
// }
