package main

import "fmt"

//切片 ： 指向相关数组的指针 、切片长度、切片容量

var months = [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
var slice []int

func main() {
	//summer := months[3:6]
	//fmt.Println(months)
	//fmt.Println(summer)
	//fmt.Println(len(summer))
	//fmt.Println(cap(summer))

	//未初始化之前的切片
	//fmt.Printf("%T\n",slice) //[]int
	//fmt.Printf("%p\n",slice) //0x0
	//fmt.Printf("%d\n",len(slice))//0
	//fmt.Printf("%d\n",cap(slice))//0

	slice1 := months[:]
	slice2 := &months
	fmt.Println(slice1)
	fmt.Println(*slice2)

	slice3 := months[2:]
	slice4 := months[2:len(months)]
	fmt.Println(slice3)
	fmt.Println(slice4)

	slice5 := []int{123}[:]
	fmt.Println(slice5)
	slice6 := slice5[:]
	fmt.Println(slice6)

}
