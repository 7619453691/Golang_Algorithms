package main

import (
	"fmt"
	"sort"
)

func printDuplicate(data []int) []int {
	x := []int{}

	for _, a := range data {
		if checkDuplicate(data, a) > 1 {
			if checkDuplicate(x, a) < 1 {
				x = append(x, a)
			}
		}
	}
	return x
}

func checkDuplicate(data []int, n int) int {
	y := 0
	for _, a := range data {
		if a == n {
			y += 1
		}
	}
	return y
}

func missingNumber(data []int) []int {
	j := 0
	missing := []int{}
	k := data[0]

	for i := k; i < k+len(data); i++ {
		if i != data[j] && checkDuplicate(data, i) < 1 {
			missing = append(missing, i)
		}
		j += 1
	}
	return missing
}

func main() {

	data := []int{28, 31, 31, 29, 30}
	sort.Ints(data)

	fmt.Println(data)

	fmt.Println("Duplicate value:", printDuplicate(data))

	fmt.Println("Missing value:", missingNumber(data))

}
