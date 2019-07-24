package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	a1, a2 := arr[0], arr[len(arr)-1]

	fmt.Println(arr, a1, a2)

	//使用for循环遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println("Element ", i, "of arr is ", arr[i])
	}

	//range 遍历数组 方式1 key and value
	for i, v := range arr {
		fmt.Println("range i", i, "value", v)
	}

	//只要value
	for _, v := range arr {
		fmt.Println("range v :", v)
	}

	//只要key
	for k := range arr {
		fmt.Println("range only k :", k)
	}
}
