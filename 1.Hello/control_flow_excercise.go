package main

import (
	"fmt"
)

func main() {
	//var name string = "Pawan Kumar"
	//var fname string
	//var lname string
	//fmt.Scanf("%s%s",&fname,&lname)
	//fmt.Println("Hello,",fname,lname)
	//var smallNumber int
	//var bigNumber int
	//fmt.Print("Enter Small Number: ")
	//fmt.Scan(&smallNumber)
	//fmt.Print("Enter Big Number: ")
	//fmt.Scan(&bigNumber)
	//fmt.Println(bigNumber % smallNumber)
	//for i := 0; i< 100; i++ {
	//	if i % 2 == 0 {
	//		fmt.Println(i)
	//	}
	//}
	//for i := 1; i< 100; i++ {
	//	if i % 3 == 0 && i % 5 == 0{
	//		fmt.Println("FizzBuzz")
	//	}else if i % 3 == 0 {
	//		fmt.Println("Fizz")
	//	}else if i % 5 == 0 {
	//		fmt.Println("Buzz")
	//	}else{
	//		fmt.Println(i)
	//	}
	//}

	fmt.Printf("%T \n", temp)

}

func temp(n int, _ int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
