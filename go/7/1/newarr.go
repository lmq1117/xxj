package main

import "fmt"

/**
	GO 中数组是值类型
   C C++ 数组是指向首元素的指针
*/
var arr3p = new([5]int)

func main() {
	arr3v := *arr3p
	arr3v[4] = 88
	fmt.Println(arr3v)
	fmt.Printf("%p", arr3p)
}
