package Poly

import (
	"fmt"
	"reflect"
)

type Point struct {
	x int
	y int
}

func Assert() {
	//var a interface{}
	//var point Point = Point{1,2}
	//a = point
	//fmt.Println(reflect.TypeOf(a))
	//var b Point
	//b = a.(Point)  //类型断言，前提是a是一个Point类型，否则会panic
	//fmt.Println(b)

	//如何做一个带检查的类型断言
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	fmt.Println(reflect.TypeOf(a))
	b, ok := a.(float64) //带判断的类型断言
	//等价 if b ,ok:= a.(float64) ; ok
	if ok == true {
		fmt.Println("convert ok")
		fmt.Println(b)
	} else {
		fmt.Println("Convert fail")
	}
	fmt.Println("GO ON")
	fmt.Println("GO ON")
	
}
