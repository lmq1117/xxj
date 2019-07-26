package main

import "fmt"

//把一个大数组传递给函数 会消耗很多内存
//	a 传递数组的指针
//	b 传切片

func main() {
	arr5 := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&arr5)
	fmt.Println(arr5)
	fmt.Println(x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
