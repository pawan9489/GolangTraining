package main

import (
	"fmt"
)

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	x := []int{1, 2, 4}
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", &x)
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	z := *y
	fmt.Println(z)
	z[0] = 43
	fmt.Printf("%T \n", z)
	fmt.Println(x)
	a1 := [2]int{1, 2}
	a2 := [2]int{1, 5}
	fmt.Println(a1)
	a1 = a2
	fmt.Println(a1)
}
