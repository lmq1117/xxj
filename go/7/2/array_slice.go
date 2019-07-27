package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * i * i
	}
	fmt.Println(arr1)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice[i])
	}

	fmt.Println("length arr1 ", len(arr1))
	fmt.Println("length slice ", len(slice))
	fmt.Println("cap arr1 ", cap(slice))

	slice = slice[0:4] //[0:7]越界panic: runtime error: slice bounds out of range

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice[i])
	}
	fmt.Println("length slice ", len(slice))
	fmt.Println("cap arr1 ", cap(slice))

	//切片可以右移，相当于从切片开头删除元素
	fmt.Println(slice)
	slice = slice[1:]
	fmt.Println(slice)

	//再右移2个
	slice = slice[2:]
	fmt.Println(slice)

}
