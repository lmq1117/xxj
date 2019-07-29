package main

import "fmt"

func main() {
	//v1 := 3.1
	//v2 := 'a'
	//v3 := "abc"
	//获取变量类型 方式1
	//fmt.Printf("%T\n", v1)

	//获取变量类型 方式2
	//fmt.Println(reflect.TypeOf(v1))
	//fmt.Printf("%T\n", v2)
	//fmt.Println(reflect.TypeOf(v2))
	//fmt.Printf("%T\n", v3)
	//fmt.Println(reflect.TypeOf(v3))

	//
	//var int_key_1 int8
	//int_key_2 := 8
	//int_key_1 = int_key_2
	//fmt.Println(int_key_1 == int_key_2)//类型不一致

	//
	//var int_key_3 int8
	//
	//int_key_4 := 30
	//
	//int_key_3 = int8(int_key_4)
	//
	//fmt.Println(int_key_3)

	var intKeyThree int8
	intKeyFour := 31
	intKeyThree = int8(intKeyFour)
	fmt.Println(intKeyThree)

}
