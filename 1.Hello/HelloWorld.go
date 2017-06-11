package main

import (
	"fmt"
	//"github.com/pawan9489/GolangTraining/1.Hello/StringUtil"
	//"github.com/pawan9489/GolangTraining/1.Hello/PKG_Folder"
	//"net/http"
	//"io/ioutil"
	"reflect"
)

const abc string = "asd"
const q = 90
const (
	_ = iota
	test = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)
const (
	A = iota
	B
	C
)


const (
	Apple, Banana = 2*iota + 1, 2*iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func main() {
	//fmt.Printf("The factorial is %d", fact(60))
	fmt.Println("Start GO")
	fmt.Println(test)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Printf("%b\n",KB)
	fmt.Printf("%b\n",MB)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(test)

	fmt.Println()
	fmt.Println(Apple)
	fmt.Println(Banana)
	fmt.Println(Cherimoya)
	fmt.Println(Durian)
	fmt.Println(Elderberry)
	fmt.Println(Fig)
	fmt.Println()

	x := 30
	y := x + 50
	fmt.Println(x,y)
	fmt.Printf("%T - %T \n", x,y)

	//a := 30
	//b := a + 50.5
	//fmt.Println(x,y)
	//fmt.Printf("%T - %T \n", a,b)

	abc := 'b'
	fmt.Printf("%T\n",abc)
	fmt.Println(reflect.TypeOf(abc))

	var uns int = -32
	fmt.Println(uns)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var str string = "My Name is Pawan Kumar"
	fmt.Println(str[3])
	fmt.Printf("%d\n",&str)
	fmt.Println(len(str))
	var myName string = "Pawan Kumar"
	fmt.Println(len(myName))
	fmt.Println(myName[:4])


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

	//res,_ := http.Get("http://www.google.com/")
	//page,_ := ioutil.ReadAll(res.Body)
	//
	//res.Body.Close()
	//fmt.Printf("%s", page)

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
