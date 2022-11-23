package main

import "fmt"

func concept1(n int) {

	if n == 0 {
		return
	}
	fmt.Println(n)
	concept1(n - 1)
}

func main() {

	concept1(5)
}
