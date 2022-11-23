package main

import "fmt"

func partition(A []int, low int, high int) int {
	pivot := 0
	pivot = A[low]
	i := low + 1
	j := high
	temp := 0

	for A[i] <= pivot {
		i++
	}
	for A[j] > pivot {
		j--
	}
	 
		temp = A[i]
		A[i] = A[j]
		A[j] = temp
	

	temp = A[low]
	A[low] = A[j]
	A[j] = temp

	return j
}

func quickSort(A []int, low int, high int) []int {
	partitionIndex := 0
	if low < high {
		partitionIndex = partition(A, low, high)
		quickSort(A, low, partitionIndex-1)
		quickSort(A, partitionIndex+1, high)
	}
	return A
}

func main() {
	A := []int{3, 5, 2, 13, 12, 3, 2, 13, 45}

	fmt.Println(quickSort(A, 0, len(A)-1))
}
