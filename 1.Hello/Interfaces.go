package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s shape
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	s = square{4}
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Printf("%T\n", s)
	s = circle{10}
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Printf("%T\n", s)
}
