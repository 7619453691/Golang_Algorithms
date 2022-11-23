package main

import "fmt"

func search(arr []int, target int, s int, e int) int {
	if s > e {
		return -1
	}

	m := s + (e-s)/2
	if arr[m] == target {
		return m
	}
	if target < arr[m] {
		return search(arr, target, s, m-1)
	}
	return search(arr, target, m+1, e)
}

func main() {
	arr := []int{1, 2, 3, 4, 55, 66, 78}
	target := 55

	fmt.Println(search(arr, target, 0, len(arr)-1))
}
