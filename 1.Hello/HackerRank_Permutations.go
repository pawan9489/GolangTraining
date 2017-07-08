package main

import (
	"fmt"
	"runtime"
)

type example struct {
	flag bool
	counter int16
	pi float32
}

func main()  {
	x := 1 << 60
	fmt.Printf("%T \n",x)
	fmt.Println(x)
	fmt.Println(runtime.GOARCH)
	var e1 example
	fmt.Println(e1)
	const (
		A = 2 * iota + 3
		B
		C
	)
	fmt.Println(A,B,C)
}