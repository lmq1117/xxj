package main

import (
	"fmt"
	"reflect"
)

func main() {
	v1 := 3.1
	v2 := 'a'
	v3 := "abc"
	//获取变量类型 方式1
	fmt.Printf("%T\n", v1)
	//获取变量类型 方式2
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(reflect.TypeOf(v2))
	fmt.Println(reflect.TypeOf(v3))
}
