package main

import "fmt"

func main() {
	arr := []int{4, 3, 12, 1, 5, 5, 3, 9}

	arr = countingsort(arr)

	fmt.Println(arr)
}

func countingsort(v []int) []int {
	n := len(v)

	high := v[0]
	for i := 1; i < n; i++ {
		if v[i] > high {
			high = v[i]
		}
	}

	ca := make([]int, high+1)

	for i := 0; i < n; i++ {
		ca[v[i]]++
	}

	for i := 1; i < high+1; i++ {
		ca[i] += ca[i-1]
	}

	oa := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		oa[ca[v[i]]-1] = v[i]
		ca[v[i]] -= 1
	}

	return oa
}

// func countingsort(arr []int) []int {
// 	max := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] > max {
// 			max = arr[i]
// 		}
// 	}

// 	countarr := make([]int, max+1)

// 	for i := 0; i < len(countarr); i++ {
// 		countarr[i] = 0
// 	}

// 	for i := 0; i < len(arr); i++ {
// 		countarr[arr[i]]++
// 	}

// 	for i := 1; i < len(countarr); i++ {
// 		countarr[i] += countarr[i-1]
// 	}

// 	outputArray := make([]int, len(arr))

// 	for i := 0; i < len(arr); i++ {
// 		outputArray[countarr[arr[i]]-1] = arr[i]
// 		countarr[arr[i]]--
// 	}

// 	return outputArray

// }
