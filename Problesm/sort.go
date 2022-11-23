package main

import "fmt"

func Sort(A []int) []int {
	temp := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			temp = A[i]
			A[i] = A[i+1]
			A[i+1] = temp

			i = -1
		}
	}
	return A
}

func main() {
	A := []int{1, 0, -1, -1, -1, 0, 1, 0}

	fmt.Println(Sort(A))
}
