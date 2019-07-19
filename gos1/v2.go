package main

import (
	"fmt"
)

func main() {

	var age int64 = 10
	var size = 18
	colorSize := 95

	fmt.Println(age, size, colorSize)
	//os.Args

	// := 左边的变量不能是已经声明过的
	//var i int32
	//i := 2
	//fmt.Println(i)

	var num int64
	num = 12
	fmt.Println(num)
	num = 24
	fmt.Println(num)

	//多重赋值
	var left int64 = 10
	var right int64 = 20
	fmt.Println(left, right)
	left, right = right, left
	fmt.Println(left, right)

	//匿名变量
	userName, guardianName := GetName()
	fmt.Println("USERNAME:", userName, ",GURADIANNAME:", guardianName)
	_, name := GetName()
	fmt.Println("NAME:", name)

	xingming, _ := GetName()
	fmt.Println("XINGMING:", xingming)

	//变量的作用域
	//name Name
	//代码块内层使用与外层相同的变量名 外部暂时不可见 执行完后外层又可见
}

func GetName() (userName, guardianName string) {
	return "阿斗", "刘备"

}
