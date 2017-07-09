package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var d []string
	fmt.Println(d)

	d1 := []string{}
	fmt.Println(d1)

	z1, _ := json.Marshal(d)
	fmt.Println(string(z1))

	z2, _ := json.Marshal(d1)
	fmt.Println(string(z2))
	fmt.Println()

	slice1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println()
	slice2 := []string{}
	slice2 = append(slice2, "a")
	slice2 = append(slice2, "b")
	slice2 = append(slice2, "c")
	slice2 = append(slice2, "d")
	slice2 = append(slice2, "e")
	slice2 = append(slice2, "f")
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println()
	slice3 := slice2[2:4:5]
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println(slice3[1])
	slice3[0] = "Changed"
	fmt.Println()

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println()

	slice3 = append(slice3, "added")
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(slice3[0:1])
}
