package main

import "fmt"

func count(n int) int {
	return helper(n, 0)
}

func helper(n int, c int) int {
	if n == 0 {
		return c
	}
	rem := n % 10
	if rem == 0 {
		return helper(n/10, c+1)
	}
	return helper(n/10, c)
}

func main() {

	fmt.Println(count(120238720020))
}
