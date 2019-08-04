package main

import (
	"fmt"
	"reflect"
)

type UserP struct {
	Id   int
	Name string
	Age  int
}

func (u UserP) Hello() {
	fmt.Println("type user.func Hello")
}

func (u UserP) Go() {
	fmt.Println("type user.func Go")
}

func InfoP(o interface{}) {
	//
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name()) //User 结构体类型名称 t.Field(0).Name 则是对应成员的类型

	//fmt.Println("Name:", t.Field(0).Name) //成员名字 Id
	//fmt.Println("Type:", t.Field(0).Type) //成员类型 int

	//获取成员属性 及成员属性的类型 成员属性的值
	//v := reflect.ValueOf(o)
	//for i := 0; i < t.NumField(); i++ {
	//	f := t.Field(i)
	//	val := v.Field(i).Interface() //成员属性的值
	//	fmt.Printf("字段名称：%6s， 字段的类型：%v， 字段的值：%v\n", f.Name, f.Type, val)
	//	//fmt.Println(f, v)
	//}

	//获取成员方法的名称 及定义
	//for i := 0; i < t.NumMethod(); i++ {
	//	m := t.Method(i)
	//	fmt.Printf("%s:=%v\n", m.Name, m.Type)
	//}

	//获取匿名|嵌入字段

}

func main() {
	u := UserP{1, "jack", 88}
	InfoP(&u)
}
