package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func pluss(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("res: ", res)
	res = pluss(1, 3, 33)
	fmt.Println("res: ", res)
}
