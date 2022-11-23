package main

import "fmt"

func insertion(A []int, n int) []int {
	j := 0
	key := 0

	for i := 1; i <= n-1; i++ {
		key = A[i]
		j = i - 1
		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
	return A
}

func main() {

	A := []int{1, 2, 5, 6, 7, 3, 9, 8}
	n := 8

	fmt.Println(insertion(A, n))
}
