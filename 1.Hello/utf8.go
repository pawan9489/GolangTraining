package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	x := "నా పేరు పవన్ కుమార్"
	for i, r := range x {
		rl := utf8.RuneLen(r)
		i += rl
		fmt.Println(rl)
		fmt.Println(i, r)
	}
}
