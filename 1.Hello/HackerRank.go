package main
import (
	"fmt"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	i := 20
	j := 23
	k := 6
	//fmt.Scan(&i)
	//fmt.Scan(&j)
	//fmt.Scan(&k)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	count := 0
	for z:= i ; z <= j; z++ {
		if calcDays(z,k) {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(reverseString(strconv.Itoa(23)))

	s := generatePattern(50)
	fmt.Println(s)
	n := 50
	switch n {
	case 0 : fmt.Println(0)
	case 1 : fmt.Println(2)
	default :
		s := generatePattern(n)
		fmt.Println(sumSlice(s))
	}
	//2068789129 -> 50
}

func generatePattern(n int) []int {
	aSlice := []int {2,3}
	for i:=1; i<n; i++ {
		temp := aSlice[i] + i
		aSlice = append(aSlice, temp)
	}
	return aSlice
}

func sumSlice(n []int)int {
	sum := 0
	for _,v := range n {
		sum += v
	}
	return sum
}

func calcDays(n int, k int)bool{
	rev,_ := strconv.Atoi(reverseString(strconv.Itoa(n)))
	fmt.Println(rev)
	if (n-rev) % k == 0 {
		return true
	}
	return false
}

func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}


