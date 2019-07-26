package main

import "fmt"

//表达式 由值运算符组合而成
//一元运算符
//二元运算符 两个类型相同的值才可以用二元运算符结合 GO是强类型 不会隐式转换，任何不同类型的值的转换都必须显式说明

var f bool = false
var t bool = true
var s string = "true and false are not eq"

func main() {
	if f != t {
		fmt.Println(s)
	}
}
