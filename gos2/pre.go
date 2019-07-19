package main

import (
	"fmt"
)

func main() {
	//预定义常量有以下三个
	//true false iota

	//iota 可以被编译器修改的常量
	const (
		c0 = iota
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c0, c1, c2, c3)

	const (
		d0 = iota * 2
		d1 = iota * 2
		d2 = iota * 3
		d3 = iota * 3
	)
	fmt.Println(d0, d1, d2, d3)

	const (
		e0 = iota
		e1
		e2
		e3
	)
	fmt.Println(e0, e1, e2, e3)

	const (
		f0 = iota * 5
		f1
		f2
		f3
	)
	fmt.Println(f0, f1, f2, f3)

	//枚举
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)

	//常量作用域
	//Name name
	//函数提体内定义的常量 只能在函数体内使用
}
