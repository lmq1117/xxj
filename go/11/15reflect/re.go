package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello type user")
}

func Info(o interface{}) {
	//
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s:=%v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{1, "jack", 88}
	Info(u)
}
