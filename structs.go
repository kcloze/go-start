package main

import "fmt"

type person struct {
	name string
	age  int
}

type oop struct {
	attr string
	age  int
}

func main() {
	fmt.Println(person{"kcloze", 29})
	fmt.Println(person{name: "kcloze", age: 29})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
