package main

import (
	"fmt"
	"math"
	"sort"
)

func main()  {
	var n int
	fmt.Scan(&n)
	var c []int
	var temp int
	for i := 0; i< n ; i++ {
		fmt.Scan(&temp)
		c = append(c,temp)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(c)))
	const max_INT = 1  << 31 -1
	sum := 0
	for i := 0; i < n ; i++ {
		if max_INT - (int(math.Pow(2,float64(i))) * c[i]) > sum {
			sum += int(math.Pow(2,float64(i)))  * c[i]
		}
	}
	fmt.Println(sum)
}
