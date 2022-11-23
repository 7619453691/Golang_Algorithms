package main

import "fmt"

func print(A []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println("\n")
}

func merge(A []int, mid int, low int, high int) {
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
}

func mergesort(A []int, low int, high int) {
	mid := 0

	if low < high {
		mid = (low + high) / 2
		mergesort(A, low, mid)
		mergesort(A, mid+1, high)
		merge(A, mid, low, high)
	}
}

func main() {

	A := []int{9, 1, 4, 14, 4, 15, 6}
	n := 7

	print(A, n)

	mergesort(A, 0, 6)

	print(A, n)
}
