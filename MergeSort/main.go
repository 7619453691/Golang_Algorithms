package main

import "fmt"

func merge(A []int, mid int, low int, high int) []int {
	i := 0
	j := 0
	k := 0
	B := [100]int{}

	i = low
	j = mid + 1
	k = low

	for i <= mid && j <= high {
		if A[i] < A[j] {
			B[k] = A[i]
			i++
			k++
		} else {
			B[k] = A[j]
			j++
			k++
		}
	}
	for i <= mid {
		B[k] = A[i]
		k++
		i++
	}
	for j <= high {
		B[k] = A[j]
		k++
		j++
	}
	for i := low; i <= high; i++ {
		A[i] = B[i]
	}
	return A
}

func mergeSort(A []int, low int, high int) []int {
	mid := 0
	if low < high {
		mid = (low + high) / 2
		mergeSort(A, low, mid)
		mergeSort(A, mid+1, high)
		merge(A, mid, low, high)
	}
	return A
}

func main() {
	A := []int{9, 1, 4, 14, 4, 15, 6}

	fmt.Println(mergeSort(A, 0, 6))
}
