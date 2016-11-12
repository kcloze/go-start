package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 4 {
		fmt.Println("8 in divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is newgative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has mutiple digits")
	}
}
