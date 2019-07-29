package main

import "fmt"

func main() {
	var a [8]byte //长度为8的数组，每个元素为一个字节

	var b [3][3]int // 二维数组

	var c [3][3][3]float64

	var d = [3]int{1, 2, 3}

	var e = new([3]string)

	fmt.Println(a, b, c, d, e)
}
