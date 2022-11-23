package main

import "fmt"

func fibon(n int) int {

	if n < 2 {
		return n
	}
	return fibon(n-1) + fibon(n-2)
}

func main() {

	ans := fibon(4)
	fmt.Println(ans)
}
