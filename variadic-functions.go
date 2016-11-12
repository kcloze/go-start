package main

import "fmt"

func sum(nums ...int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	sum(1, 2, 33)
	sum(3, 333, 324, 1, 2, 33)
	nums := []int{323, 234, 45, 232020, 233}
	sum(nums...)
}
