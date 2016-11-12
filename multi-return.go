package main

import "fmt"

func value() (int, int) {
	return 4, 5
}
func main() {
	a, b := value()
	fmt.Println(a, b)
	_, c := value()
	fmt.Println(c)
}
