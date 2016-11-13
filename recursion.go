package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res := fact(7)
	fmt.Println(res)
	res = fact(20)
	fmt.Println(res)
}
