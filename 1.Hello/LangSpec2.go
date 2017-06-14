package main

import (
	"fmt"
)

func main()  {
	//a := 42
	//fmt.Println(&a)
	//var b* int = &a
	//fmt.Println(b)
	//var c *int = &a
	//fmt.Println(c)
	//
	//fmt.Println(*b)
	//fmt.Println(*c)
	//*b = 55
	//fmt.Println(*b)
	//fmt.Println(*c)
	//fmt.Println(a)

	//for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
	//	fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	//}
	////}

	switch "Pawan" {
	case "kumar":
	case "Pawan":
		fmt.Println("First Name")
		fmt.Println("Multi Line without braces are allowed")
		fallthrough
	default:
		fmt.Println("default")
	}

	switch "OO" {
	case "PK", "OK":
		fmt.Println("PK or OK")
	case "OO":
		fmt.Println("OO")
	}

}