package main

import "fmt"

//以下两种方式生成相同的切片
//make([]int,50,100)	适合切片 map channel
//new([100]int)[:50] //new数组后切片  适合值类型如数组和结构体 相当于&T{}

func main(){
	var slice1 []int = make([]int,10)
	for i:=0;i<len(slice1);i++ {
		slice1[i] = 5 * i
		fmt.Println(slice1[i])
	}
	fmt.Println(slice1);
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
