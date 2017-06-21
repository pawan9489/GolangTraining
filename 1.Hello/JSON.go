package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int //`json:"-"`
	notExported int
}

func main() {
	p1 := person{"Pawan","Kumar",23,55}
	fmt.Println(p1)
	bs,_ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)
	fmt.Println(string(bs))

	x := 7
	fmt.Println(x)

	str := "My Name is Pawan Kumar"
	bs1,_ := json.Marshal(str)
	fmt.Println(bs1)
	fmt.Println(string(bs1))
	json.NewEncoder(os.Stdout).Encode(p1)
}
