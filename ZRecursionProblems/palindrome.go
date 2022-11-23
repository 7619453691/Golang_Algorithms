package main

import (
	"fmt"
	"math"
)

func reverse(n int) int {
	digits := (int)(math.Log10(float64(n))) + 1
	return helper1(n, digits)
}

func helper1(n int, digits int) int {
	if n%10 == n {
		return n
	}
	rem := n % 10
	return rem * (int)(math.Pow(10, float64(digits-1))+float64(helper1(n/10, digits-1)))
}

func palindrome(n int) bool {
	return n == reverse(n)
}

func main() {
	fmt.Println(palindrome(1234321))
}
