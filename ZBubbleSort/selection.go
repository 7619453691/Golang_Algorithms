package main

import "fmt"

func selection(A []int, n int) []int {
	indexOfMin := 0
	temp := 0

	fmt.Printf("Running selection sort...\n")
	for i := 0; i < n-1; i++ {
		indexOfMin = i
		for j := i + 1; j < n; j++ {
			if A[j] < A[indexOfMin] {
				indexOfMin = j
			}
		}
		// swap A[i] and A[indexOfMin]
		temp = A[i]
		A[i] = A[indexOfMin]
		A[indexOfMin] = temp
	}
	return A
}

func main() {
	A := []int{1, 2, 5, 6, 7, 3, 9, 8}
	n := 8

	fmt.Println(selection(A, n))
}
