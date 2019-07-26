package main

import "fmt"

//数组的声明
//var 数组标识符 [len]type
var arr1 [5]int

func main() {
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("Array at index ", i, " is ", arr1[i])
	}

	fmt.Println("#######################")

	for i, _ := range arr1 {
		fmt.Println("Array range at index ", i, " is ", arr1[i])
	}
	fmt.Println("#######################")
	for i, v := range arr1 {
		fmt.Println("Array range i v at index ", i, " is ", v)
	}

	fmt.Println("#######################")
	for i := range arr1 {
		fmt.Println("Array range i only at index ", i, " is ", arr1[i])
	}

	fmt.Println()

	a := [...]string{"aa", "bb", "cc", "dd", "ee"}
	for i := range a {
		fmt.Println("arr a i ", i, " is ", a[i])
	}

	//var arr2 = new([5]int)

}
