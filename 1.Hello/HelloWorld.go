package main

import "fmt"

func main() {
	fmt.Printf("The factorial is %d", fact(60))
}

func factorial(num int) int {
	if num <= 1 {
		return num
	}
	return num * factorial(num - 1)
}

func fact(num int) int{
	//ans := 1
	//func helper(acc int, y int) int {
	//	if y == 1{
	//		return acc
	//	}
	//	return helper(acc * y, y-1)
	//}
	return helper2(1,num)
}
func helper2(acc int, y int) int {
	if y == 1{
		return acc
	}
	return helper2(acc * y, y-1)
}