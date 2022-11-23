// 01) Program
// *
// * *
// * * *
// * * * *
// * * * * *

// package main

// import "fmt"

// func pattern(n int) {
// 	for row := 1; row <= n; row++ {
// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {

// 	pattern(5)
// }

//02) Program
// * * * * *
// * * * * *
// * * * * *
// * * * * *
// * * * * *

// package main

// import "fmt"

// func pattern(n int) {
// 	for row := 1; row <= n; row++ {
// 		for col := 1; col <= n; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	pattern(5)
// }

//03) Program
// * * * * *
// * * * *
// * * *
// * *
// *
// package main

// import "fmt"

// func pattern(n int) {
// 	for row := 1; row <= n; row++ {
// 		for col := 1; col <= n-row+1; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	pattern(5)
// }

//04) Program
// *
// * *
// * * *
// * * * *
// * * * * *
// package main

// import "fmt"

// func pattern(n int) {
// 	for row := 1; row <= n; row++ {
// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	pattern(5)
// }

//05) Program

// *
// * *
// * * *
// * * * *
// * * * * *
// * * * *
// * * *
// * *
// *

// package main

// import "fmt"

// func print(n int) {

// 	for row := 0; row < 2*n; row++ {
// 		totalColInRow := 0
// 		if row > n {
// 			totalColInRow = 2*n - row
// 		} else {
// 			totalColInRow = row
// 		}
// 		for col := 0; col < totalColInRow; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	print(5)
// }

//06) Program

// *
// * *
// * * *
// * * * *
// * * * * *
// * * * *
// * * *
// * *
// *
// package main

// import "fmt"

// func print(n int) {
// 	for row := 0; row < 2*n; row++ {
// 		totalColInRow := 0

// 		if row > n {
// 			totalColInRow = 2*n - row
// 		} else {
// 			totalColInRow = row
// 		}

// 		noSpaces := n - totalColInRow

// 		for s := 0; s < noSpaces; s++ {
// 			fmt.Print(" ")
// 		}

// 		for col := 0; col < totalColInRow; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	print(5)
// }

//07) Program
//        1
//      2  1 2
//    3 2 1 2 3
//  4 3 2 1 2 3 4
//5 4 3 2 1 2 3 4 5

// package main

// import "fmt"

// func print(n int) {
// 	for row := 0; row <= n; row++ {
// 		for space := 0; space < n-row; space++ {
// 			fmt.Print("  ")
// 		}

// 		for col := row; col >= 1; col-- {
// 			fmt.Print(col, " ")
// 		}

// 		for col := 2; col <= row; col++ {
// 			fmt.Print(col, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {

// 	print(5)
// }

//Program 8
//         1
//       2 1 2
//     3 2 1 2 3
//   4 3 2 1 2 3 4
// 5 4 3 2 1 2 3 4 5
//   4 3 2 1 2 3 4
//     3 2 1 2 3
//       2 1 2
//         1
// package main

// import "fmt"

// func print(n int) {
// 	for row := 1; row <= 2*n; row++ {
// 		c := 0
// 		if row > n {
// 			c = 2*n - row
// 		} else {
// 			c = row
// 		}

// 		for space := 0; space < n-c; space++ {
// 			fmt.Print("  ")
// 		}

// 		for col := c; col >= 1; col-- {
// 			fmt.Print(col, " ")
// 		}

// 		for col := 2; col <= c; col++ {
// 			fmt.Print(col, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	print(5)
// }

//09) program

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func print(n int) {
// 	originalN := n
// 	n = 2 * n

// 	for row := 0; row <= n; row++ {
// 		for col := 0; col <= n; col++ {
// 			atEveryIndex := originalN - math.Min(math.Min(row, col), math.Min(n-row, n-col))
// 			fmt.Print(atEveryIndex, " ")
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	print(5)
}
