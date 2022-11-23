package main

import "fmt"

func print(n int) {
	if n == 5 {
		fmt.Println(5)
		return
	}
	fmt.Println(n)
	print(n + 1)
}

func main() {
	print(1)
}
