package main

import "fmt"

func bubbleSort(A []int) []int {
	temp := 0
	for i := 0; i < len(A)-1; i++ {
		// fmt.Printf("Working on pass number %d\n", i+1)
		for j := 0; j < len(A)-1-i; j++ {
			if A[j] > A[j+1] {
				temp = A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}
	return A
}

func main() {
	A := []int{12, 54, 65, 7, 23, 9}

	fmt.Println(bubbleSort(A))
}
