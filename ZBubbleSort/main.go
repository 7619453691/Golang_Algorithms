package main

import "fmt"

func bubbleSort(A []int, n int) []int {
	temp := 0
	// isSorted := 0

	for i := 0; i < n-1; i++ {
		fmt.Printf("Working on pass number %d\n", i+1)
		// isSorted = 1
		for j := 0; j < n-1-i; j++ {
			if A[j] > A[j+1] {
				temp = A[j]
				A[j] = A[j+1]
				A[j+1] = temp
				// isSorted = 0
			}
		}
		// if isSorted == 1 {
		// 	return
		// }
	}
	return A
}
func main() {

	A := []int{1, 2, 5, 6, 7, 3, 9, 8}
	// A := []int{1, 2, 3, 4, 5, 6}
	n := 8

	fmt.Println(bubbleSort(A, n))

}
