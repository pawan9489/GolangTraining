package DataStructures

import "fmt"

func Maps(){
	x := map[string]map[string]string {
		"He" : {
			"name" : "Helium",
			"state" : "gas",
		},
	}

	y := map[string]map[string]string {"He" : {"name" : "Helium","state" : "gas"}}
	fmt.Println(x)
	fmt.Println(y)

	a := map[string]string {
		"Ann" : "one",
		"Joe" : "two",

	}
	xx := a["Ann"]
	fmt.Println(xx)
	yy := a["Annsdfsd"]
	fmt.Println(yy)
	fmt.Println(a["Pawan"] == "")

	b := map[string]map[string]int{}
	fmt.Println(b["PK"] == nil)
}
