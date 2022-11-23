//01) Progranm

// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Print("Enter the number of rows:")
// 	fmt.Scan(&num)

// 	for i := 1; i < num+1; i++ {
// 		for j := 1; j < i+1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

//02) Program

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var num int
// 	fmt.Print("Enter the number of rows : ")
// 	fmt.Scan(&num)
// 	k := 1

// 	for i := 1; i < num+1; i++ {
// 		for j := 1; j < k+1; j++ {
// 			fmt.Print("1 ")
// 		}
// 		k = k + 2
// 		fmt.Println()
// 	}
// }

// 03) Program
// check once

// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Println("Enter the number of rows: ")
// 	fmt.Scan(&num)

// 	for i := 0; i < num; i++ {
// 		for j := 0; j < num-i-1; j++ {
// 			fmt.Print("\n")
// 			for j := 0; j < i+1; j++ {
// 				fmt.Print("*")
// 			}
// 			fmt.Println()
// 		}
// 	}
// }

//04) Program

// package main

// import "fmt"

// func pyramid(rows rune) {
// 	for i := 0; i < int(rows); i++ {
// 		fmt.Println(' '*(rows-i-1)+ "*"(i+1))
// 	}
// }

// func main() {

// }

// Program 5

// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Println("Enter the number of rows: ")
// 	fmt.Scan(&num)

// 	for i := num; i > 0; i-- {
// 		for j := 0; j < num-i; j++ {
// 			fmt.Print()
// 		}
// 		for j := 0; j < i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// program 6

// package main

// import (
// 	"fmt"

// )

// func pyramid(rows rune) {
// 	for i := 0; i < int(rows); i++ {
// 		fmt.Println(' '*(rows-i-1)+'* '*(i+1))
// 	}
// }

// func main() {
// 	var num int
// 	fmt.Println("Enter the number of rows: ");
// 	fmt.Scan(&num)
// }

/// Program 7

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var num int
// 	fmt.Println("Enter the number of rows: ")
// 	fmt.Scan(&num)

// 	for i := num; i > 0; i-- {
// 		for j := 0; j < i; j++ {
// 			fmt.Print("* ")
// 		}
// 	}
// 	fmt.Println()
// }

//08) Program

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter the number of rows: ")
	fmt.Scan(&num)

	var row rune
	var col rune

	for i := 0; i < int(row); i++ {
		for j := 0; j < int(col); j++ {
			if ((col == 0 || col == 4) && row != 0) || ((row == 0 || row == 3) && (col > 0 && col < 4)) {
				fmt.Print("*")
			} else {
				fmt.Println(" ")
			}
		}
		fmt.Println()
	}
}
