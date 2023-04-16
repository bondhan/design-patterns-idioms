package main

import "fmt"

type Integer interface {
	int | int32 | int64 | uint | uint32 | uint64
}

type shape2D[T Integer] struct {
	length T
	width  T
}

type shape3D[T Integer] struct {
	shape2D[T]
	height T
}

func (s *shape2D[T]) area() T {
	return calcArea(s.length, s.width)
}

func (s *shape3D[T]) volume() T {
	return s.area() * s.height
}

func calcArea[T Integer](length, width T) T {
	return length * width
}

func main() {
	s2d := shape2D[int]{
		length: 10,
		width:  12,
	}

	s3d := shape3D[int]{
		shape2D: s2d,
		height:  13,
	}

	fmt.Printf("properties: %v, area: %d\n", s2d, s2d.area())
	fmt.Printf("properties: %v, volume: %d\n", s3d, s3d.volume())
}
