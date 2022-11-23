package main

import "fmt"

type binarySearch struct {
	nums   []int
	key    int
	length int
}

func (bs *binarySearch) BinarySearch(nums []int, key int) int {
	low := 0
	high := bs.length - 1
	for low <= high {
		mid := (low + high) / 2
		if bs.nums[mid] == key {
			return mid
		}
		if key < bs.nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	bs := binarySearch{}

	nums := []int{1, 10, 20, 47, 59, 65, 75, 88, 99}

	fmt.Println(bs.BinarySearch(nums, 65))
}
