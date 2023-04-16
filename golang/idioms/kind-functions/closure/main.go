package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c1 := counter()
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	c2 := counter()
	fmt.Println(c2()) // 1
	fmt.Println(c2()) // 2
}
