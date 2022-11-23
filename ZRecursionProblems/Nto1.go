package main

import "fmt"

func fun(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	fun(n - 1)
}

func funcRev(n int) {
	if n == 0 {
		return
	}
	funcRev(n - 1)
	fmt.Println(n)
}

func funcBoth(n int) {
	if n == 0 {
		return
	}
	funcBoth(n - 1)
	fmt.Println(n)
}

func main() {
	funcBoth(5)
}
