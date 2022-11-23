package main

import "fmt"

func sumOfProduct(n int) int {
	if n%10 == n {
		return n
	}
	return (n % 10) * sumOfProduct(n/10)
}

func main() {
	ans := sumOfProduct(1432)
	fmt.Println(ans)
}
