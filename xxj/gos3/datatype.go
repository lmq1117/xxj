package main

import "fmt"

func main() {
	//基本数据类型
	// 布尔 bool
	// 整型 int8 byte int16 int32 int64 uint8 uint16 uint32 uint64 uintptr
	// 浮点型 float32 float64
	// 复数   complex64 complex128
	// 字符串 string
	// 字符 rune
	// 错误类型 error

	//复合类型
	//指针 point
	//数组 array
	//切片 slice
	//字典 map
	//通道 chan
	//结构体 struct
	//接口 interface

	var v1 bool
	v1 = true
	v2 := (1 == 2)
	fmt.Println(v1, v2)
}
