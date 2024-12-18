package main

import "fmt"

func main() {
	arr := []int{5, 4, 2, 8, 9, 6, 5, 7}

	sort(arr, 0, 7)

	fmt.Println(arr)
}

func merge(v []int, left int, mid int, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	vl := make([]int, n1)
	vr := make([]int, n2)

	for i := 0; i < n1; i++ {
		vl[i] = v[left+i]
	}

	for i := 0; i < n2; i++ {
		vr[i] = v[mid+1+i]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if vl[i] <= vr[j] {
			v[k] = vl[i]
			i++
		} else {
			v[k] = vr[j]
			j++
		}
		k++
	}

	for i < n1 {
		v[k] = vl[i]
		i++
		k++
	}

	for j < n2 {
		v[k] = vr[j]
		j++
		k++
	}
}

func sort(v []int, ini int, end int) {
	if ini >= end {
		return
	}

	mid := ini + (end-ini)/2

	sort(v, ini, mid)
	sort(v, mid+1, end)
	merge(v, ini, mid, end)
}

// func merge(v []int, left int, mid int, right int) {
// 	n1 := mid - left + 1
// 	n2 := right - mid

// 	l := make([]int, n1)
// 	r := make([]int, n2)

// 	for i := 0; i < n1; i++ {
// 		l[i] = v[left+i]
// 	}
// 	for i := 0; i < n2; i++ {
// 		r[i] = v[mid+1+i]
// 	}

// 	i := 0
// 	j := 0

// 	k := left

// 	for i < n1 && j < n2 {
// 		if l[i] <= r[j] {
// 			v[k] = l[i]
// 			i++
// 		} else {
// 			v[k] = r[j]
// 			j++
// 		}
// 		k++
// 	}

// 	for i < n1 {
// 		v[k] = l[i]
// 		i++
// 		k++
// 	}

// 	for j < n2 {
// 		v[k] = r[j]
// 		j++
// 		k++
// 	}
// }

// func sort(arr []int, left int, right int) {
// 	mid := left + (right-left)/2

// 	if left < right {
// 		sort(arr, left, mid)
// 		sort(arr, mid+1, right)

// 		merge(arr, left, mid, right)
// 	}
// }
