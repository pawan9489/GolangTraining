package main

import "fmt"

type foo int

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) fullName() string {
	return p.First + " " + p.Last
}

func main() {
	var p1 Person
	fmt.Println(p1)
	p2 := Person{"Pawan", "Kumar", 23}
	p3 := Person{"Sanjay", "RamaSwamy", 40}
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p2.fullName())
}
