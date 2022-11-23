//01) write a function to return a slice of all palindrome bw 10 to 10000

package main

import "fmt"

func answer() []int {
	res := []int{}

	for i := 10; i <= 1000; i++ {
		if isPalindrome(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPalindrome(n int) bool {
	temp := n
	res := 0

	for temp > 0 {
		res = res*10 + temp%10
		temp /= 10
	}
	return map[bool]bool{true: true, false: false}[n == res]
}

func main() {
	res := answer()
	fmt.Println(res)
}
