package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println("appd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c: ", c)

	l := s[2:5]
	fmt.Println("l1: ", l)

	l = s[:5]
	fmt.Println("l2: ", l)

	l = s[2:]
	fmt.Println("l3: ", l)

	t := []string{"d", "h", "f"}
	fmt.Println("t: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD: ", twoD)
}
