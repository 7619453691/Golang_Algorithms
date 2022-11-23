package main

import "fmt"

func selectionSort(A []int) []int {
	indexOfMin := 0
	temp := 0
	fmt.Printf("Running selection sort...\n")

	for i := 0; i < len(A)-1; i++ {
		indexOfMin = i
		for j := i + 1; j < len(A); j++ {
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
	A := []int{3, 5, 2, 13, 12}

	fmt.Println(selectionSort(A))
}
