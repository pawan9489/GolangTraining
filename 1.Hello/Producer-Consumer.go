package main

import (
	"time"
	"fmt"
)

func main()  {
	buffer_size := 10
	consumers := 4
	c := make(chan int,buffer_size)
	//done := make(chan bool)
	done_consuning := make(chan bool)
	slicer := []int { 2,7,1,4,9,5}

	go func() {
		for _,v := range slicer {
			c <- v
		}
		//done <- true
		close(c)
	}()

	//go func() {
	//	<- done // wait for the value in the done
	//	close(c)
	//}()
	//
	//for n := range c {
	//	fmt.Println("Doing Some work -> wait", n)
	//	time.Sleep( time.Duration(n) * time.Second)
	//	fmt.Println("Done",n)
	//}

	for i := 0; i< consumers ; i++ {
		go func() {
			for n := range c {
				fmt.Println("Doing Some work -> wait", n)
				time.Sleep( time.Duration(n) * time.Second)
				fmt.Println("Done",n)
			}
			done_consuning <- true
		}()
	}

	for i:=0;i<consumers ; i++ {
		<- done_consuning
	}
}
