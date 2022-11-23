package main

import "fmt"

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumOfDigits(n/10)
}

func main() {
	ans := sumOfDigits(1432)
	fmt.Println(ans)
}
