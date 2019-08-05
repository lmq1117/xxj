package main

import (
	"fmt"
	"reflect"
)

//const (
//	Invalid Kind = iota
//	Bool
//	Int
//	Int8
//	Int16
//	Int32
//	Int64
//	Uint
//	Uint8
//	Uint16
//	Uint32
//	Uint64
//	Uintptr
//	Float32
//	Float64
//	Complex64
//	Complex128
//	Array
//	Chan
//	Func
//	Interface
//	Map
//	Ptr
//	Slice
//	String
//	Struct
//	UnsafePointer
//)

type MyInt int

func main() {
	var x float64 = 3.4
	//fmt.Println(reflect.TypeOf(x),reflect.ValueOf(x))

	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println(v.Kind()) //Kind()总是返回底层类型
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Kind() == reflect.Float32)

	var m MyInt = 5
	vm := reflect.ValueOf(m)
	fmt.Println(vm, reflect.TypeOf(vm), reflect.ValueOf(vm), vm.Kind(), vm.Interface())
}
