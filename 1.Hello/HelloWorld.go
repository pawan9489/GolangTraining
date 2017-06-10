package main

import (
	"fmt"
	//"github.com/pawan9489/GolangTraining/1.Hello/StringUtil"
	//"github.com/pawan9489/GolangTraining/1.Hello/PKG_Folder"
	"net/http"
	"io/ioutil"
)

//import (
//	"fmt"
//	"github.com/pawan9489/GolangTraining/1.Hello/StringUtil"
//)

const abc string = "asd"
const q = 90


func main() {
	//fmt.Printf("The factorial is %d", fact(60))
	fmt.Println("Start GO")
	//StringUtil.ConcatString(StringUtil.MyName)
	//fmt.Printf("%-d %b %o %#x % X %q", 333, 333, 333, 333, 333, 333)
	//var abc int = 33
	//fmt.Printf("%d",abc)
	//
	//for i:=0; i<200; i++{
	//	fmt.Printf("%d \t%b \t%o \t%#x \t%X \t%c \t%q \n",i,i,i,i,i,i,i)
	//}
	//fmt.Println(pkg1.Var1)
	//var x = StringUtil.MyName
	//fmt.Printf("%T",x)
	//x := 0
	//increment := func () int {
	//	x++
	//	return x
	//}
	//
	//
	//fmt.Println(increment())
	//fmt.Println(increment())
	//fmt.Println(increment())
	//
	//{ // Scopes
	//	x = 22
	//	fmt.Println(x)
	//	y := 33
	//	fmt.Println(y)
	//	fmt.Println("-------")
	//	{
	//		fmt.Println(x)
	//		fmt.Println(y)
	//		y := 44
	//		x = 55
	//		fmt.Println(y)
	//		fmt.Println(x)
	//		fmt.Println("-------")
	//	}
	//	fmt.Println(y)
	//	fmt.Println(x)
	//	fmt.Println("-------")
	//}
	//fmt.Println(x)
	//fmt.Println("-------")

	res,_ := http.Get("http://www.google.com/")
	page,_ := ioutil.ReadAll(res.Body)

	res.Body.Close()
	fmt.Printf("%s", page)

}

func factorial(num int) int {
	if num <= 1 {
		return num
	}
	return num * factorial(num-1)
}

func fact(num int) int {
	//ans := 1
	//func helper(acc int, y int) int {
	//	if y == 1{
	//		return acc
	//	}
	//	return helper(acc * y, y-1)
	//}
	return helper2(1, num)
}
func helper2(acc int, y int) int {
	if y == 1 {
		return acc
	}
	return helper2(acc*y, y-1)
}
