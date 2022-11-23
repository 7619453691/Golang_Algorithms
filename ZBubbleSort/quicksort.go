// 01) i = low
// 02) j = high
// 03) pivot = low
// 04) i++ until element > pivot is found
// 05) j-- until element <= pviot is found
// 06) Swap A[i] & A[j] and Repeat 4+5 until j < i
// 07) swap pivot & A[j]

package main

import "fmt"

func partition(A []int, low int, high int) int {
	pivot := A[low]
	i := low + 1
	j := high
	temp := 0

	i = 0
	for next := true; next; next = i < j {
		for A[i] <= pivot {
			i++
		}
		for A[j] > pivot {
			j--
		}

		if i < j {
			temp = A[i]
			A[i] = A[j]
			A[j] = temp
		}
	}
	temp = A[low]
	A[low] = A[j]
	A[j] = temp
	return j
}

func quickSort(A []int, low int, high int) []int {
	partionIndex := 0

	if low < high {
		partionIndex = partition(A, low, high)
		quickSort(A, low, partionIndex-1)  // sort left subarray
		quickSort(A, partionIndex+1, high) // sort right subarray
	}
	return A
}
func main() {

	A := []int{3, 5, 2, 13, 12, 3, 2, 13, 45}
	// A := []int{2, 6, 2, 9, 3, 4, 10, 12, 16, 13, 14}
	n := 9

	fmt.Println(quickSort(A, 0, n-1))
}
