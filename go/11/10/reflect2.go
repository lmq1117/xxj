package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.4
	v := reflect.ValueOf(f)
	fmt.Println("v是否CanSet ", v.CanSet(), v.Type())

	//panic: reflect: reflect.Value.SetFloat using unaddressable value
	//v.SetFloat(3.1415)
	v = reflect.ValueOf(&f)
	fmt.Println("v是否CanSet ", v.CanSet(), v.Type())
	v = v.Elem()
	fmt.Println("v是否CanSet ", v.CanSet(), v.Type())
	v.SetFloat(3.1415)
	fmt.Println(v)

}
