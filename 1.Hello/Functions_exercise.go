package main

import "fmt"

func main() {
	fmt.Println(exercise1(3))
	exercise2 := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(exercise2(3))
	fmt.Println(exercise3(1, 2, 3, 5, 6, 76546, 7774, 2, 5, 5))
	exercise5(1)
	exercise5(1, 2, 3, 4)
	aSlice := []int{1, 2, 3, 4, 5}
	exercise5(aSlice...)
	exercise5()
}

func exercise1(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func exercise3(n ...int) int {
	temp := 0
	for _, v := range n {
		if temp < v {
			temp = v
		}
	}
	return temp
}

//exercise4 -> true

func exercise5(n ...int) {
	
}
