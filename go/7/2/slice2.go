package main

import "fmt"

//将切片通过参数传入函数

func sum(a []int)int{
	s := 0
	for i := range a {
		s += a[i]
	}
	return s
}

func main(){
	var slice = []int{1,2,3,4,5}
	fmt.Println(sum(slice))
}
