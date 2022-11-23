package main

import "fmt"

func main()  {
	s := "python"
	firstNonRepeated(s)
}

func firstNonRepeated(s string) {
	chars := []byte(s)

	for i := 0; i < len(chars); i++ {
		count := 1

		if chars[i] != '$' {
			for j := i + 1; j < len(chars); j++ {
				if chars[j] != '$' {
					if chars[i] == chars[j] {
						count++
						chars[j] = '$'
					}
				}
			}
			if count == 1 {
				fmt.Println(chars[i])
				break
			}
		}
	}
}

