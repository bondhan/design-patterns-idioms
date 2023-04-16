package main

import "fmt"

func apply(f func(int) int, x int) int {
	return f(x)
}

func double(x int) int {
	return x * 2
}

func main() {
	result := apply(double, 5)
	fmt.Println(result)
}
