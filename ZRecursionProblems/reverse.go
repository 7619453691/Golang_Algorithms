package main

import (
	"fmt"
	"math"
)

func reverse1(n int) {
	sum := 0
	if n == 0 {
		return
	}
	rem := n % 10
	sum = sum*10 + rem
	reverse1(n / 10)
}

func reverse2(n int) int {
	digits := (int)(math.Log10(float64(n))) + 1
	return helper(n, digits)
}

func helper(n int, digits int) int {
	if n%10 == n {
		return n
	}
	rem := n % 10
	return rem * (int)(math.Pow(10, float64(digits-1))+float64(helper(n/10, digits-1)))
}

func main() {

	// reverse1(1432)

	fmt.Println(reverse2(1234))
}
