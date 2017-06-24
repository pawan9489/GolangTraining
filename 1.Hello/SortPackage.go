package main

import (
	"sort"
	"fmt"
	"strings"
	"time"
)

type people []string

func main() {
	p := people{"Zoe", "action","Pawan", "Kennedy","Obama", "chocolate"}
	fmt.Println(p)

	start := time.Now()
	sort.Strings(p)
	elapsed := time.Since(start)
	fmt.Println("Packages method took", elapsed)

	fmt.Println(p)

	fmt.Println("My Custom Sort")
	start1 := time.Now()
	customSort(p)
	elapsed1 := time.Since(start1)
	fmt.Println("My method took", elapsed1)
	fmt.Println(p)
}

func customSort(s []string) []string {
	for k := len(s);k >= 0 ; k-- {
		for i,j := 0,1 ; i < k && j < k; i,j = i+1,j+1 {
			if !isBefore(s[i],s[j]) {
				swapIndexes(s,i,j)
			}
		}
	}
	return s
}

func swapIndexes(s []string,i,j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func isBefore(s1 string, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	if s1[0] > s2[0] {
		return false
	}else {
		var res bool // false
		var temp string
		if len(s1) < len(s2) {
			temp = s1
		}	else {
			temp = s2
		}
		for i,j:= 1,1 ; i<len(temp);i,j = i+1,j+1{
			if s1[i] < s2[i] {
				res = true
			}
		}
		return res
	}
}