package main

import "fmt"

func insertionSort(A []int) []int {
	key := 0
	j := 0

	for i := 1; i < len(A)-1; i++ {
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
	A := []int{12, 54, 65, 7, 23, 9}

	fmt.Println(insertionSort(A))

}
