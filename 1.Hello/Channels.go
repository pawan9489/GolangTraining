package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Placing values in C")
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println("Printing Value of C")
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
