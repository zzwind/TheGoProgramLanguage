package main

import (
	"fmt"
	"math"
	"reflect"
)

//
//func main() {
//	//const abc = 111
//	//fmt.Printf("%T\n", abc)
//	//fmt.Println(math.Pi * abc)
//	//fmt.Printf("%T", math.Pi)
//	////var i3 int = 3
//	////fmt.Println(math.Pi * i3)
//
//	const hello = "Hello, 世界"
//	const typedHello string = "Hello, 世界"
//
//	var s string
//	s = typedHello
//	fmt.Println(s)
//
//	//type Mystring string
//	//var m Mystring
//	//m = typedHello
//	//fmt.Println(m)
//
//	fmt.Printf("%T : %v", hello, hello)
//}

func main() {
	const abc = 8 //this is untyped int
	//const abc int = 8 //this is typed int
	fmt.Printf("%t\n", abc)
	fmt.Println(math.Pi * abc)
	//fmt.Printf("%t", abc)

	type Myint int

	const test1 int = 10
	const test2 = 10

	var mmint Myint

	fmt.Printf("test1 type:%T,test2 type:%T,test1 mmint:%T,", test1, test2, mmint)

	mmint = test2

	fmt.Println(mmint)

	fmt.Println(reflect.TypeOf(test2))

	fmt.Println(mmint)

	//var a int
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)
}
