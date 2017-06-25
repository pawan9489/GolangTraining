package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("from Parallelism.go")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup
	func1 := func() {
		for i := 0; i < 10; i++ {
			fmt.Println("func1", i)
			time.Sleep(time.Duration(300 * time.Millisecond))
		}
		wg.Done()
	}
	func2 := func() {
		for i := 0; i < 10; i++ {
			fmt.Println("func2", i)
			time.Sleep(time.Duration(300 * time.Millisecond))
		}
		wg.Done()
	}
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}

//func main()  {
//	wg.Add(1)
//	fmt.Println("from main()")
//	fmt.Println(runtime.NumCPU())
//	go calculatePrimes(100000000)
//	wg.Wait()
//}
//
//func calculatePrimes(n int)  {
//	for i:= 2 ;  i < n; i ++ {
//		if isPrime(i) {
//			fmt.Println(i)
//		}
//	}
//	wg.Done()
//}
//
//func isPrime(n int) bool {
//	if n == 2 || n == 3 {
//		return true
//	}
//	if n % 2 == 0 || n % 3 == 0 {
//		return false
//	}
//	for i := 3; i < int(math.Sqrt(float64(n))); i += 2 {
//		if n % i == 0 {
//			return false
//		}
//	}
//	return true
//}
