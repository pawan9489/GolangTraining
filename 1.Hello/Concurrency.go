package main

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	fmt.Println("from Concurrency.go")
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1", i)
		time.Sleep(time.Duration(300 * time.Millisecond))
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2", i)
		time.Sleep(time.Duration(300 * time.Millisecond))
	}
	wg.Done()
}
