package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{5, 43, 6, 8, 54, 3, 6, 86, 543, 33}

	quickSort(arr, 0, 9)

	fmt.Println(arr)
}

func partition(v []int, ini int, end int) int {
	pivot := rand.Intn(end-ini+1) + ini

	v[pivot], v[end] = v[end], v[pivot]

	i := ini - 1
	for j := ini; j < end; j++ {
		if v[end] > v[j] {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[end] = v[end], v[i+1]

	return i + 1
}

func quickSort(v []int, ini int, end int) {
	if ini >= end {
		return
	}

	pivot := partition(v, ini, end)

	quickSort(v, ini, pivot-1)
	quickSort(v, pivot+1, end)
}

// func partition(arr []int, low int, high int) int {
// 	pivot := high

// 	i := low - 1

// 	for j := low; j <= pivot-1; j++ {
// 		if arr[j] < arr[pivot] {
// 			i++
// 			swap(arr, i, j)
// 		}
// 	}

// 	swap(arr, pivot, i+1)
// 	return i + 1
// }

// func swap(arr []int, i int, j int) {
// 	temp := arr[i]
// 	arr[i] = arr[j]
// 	arr[j] = temp
// }

// func quickSort(arr []int, low int, high int) {
// 	if low < high {
// 		pivot := partition(arr, low, high)

// 		quickSort(arr, low, pivot-1)
// 		quickSort(arr, pivot+1, high)
// 	}
// }
