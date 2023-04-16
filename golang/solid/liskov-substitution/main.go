package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func main() {
	var s Shape
	s = Rectangle{Width: 3, Height: 4}
	fmt.Println(s.Area()) // 12

	s = Square{Side: 5}
	fmt.Println(s.Area()) // 25
}
