package main

import "fmt"

func main(){
	slFrom := []int{1,2,3}
	slTo := make([]int,10)

	fmt.Println(slTo,slFrom)

	//返回值 n copied 元素个数
	n := copy(slTo,slFrom)
	fmt.Println(n)
	fmt.Println(slTo,slFrom)

	slThree := []int{1,2,3}
	fmt.Println(slThree)
	slThree = append(slThree,4,5,6,7,8)
	fmt.Println(slThree)
	for i:=9;i<100000000;i++ {
		slThree = append(slThree,i)
	}

	//从1开始 所以减2
	fmt.Println(slThree[100000000 - 2])
}