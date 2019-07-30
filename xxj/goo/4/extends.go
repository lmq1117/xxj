package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) Call() string {
	return "动物叫声..."
}

//FavorFood
func (a Animal) FF() string {
	return "动物爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.name + "..."
}

//继承动物struct
type Dog struct {
	Animal
}

func (d Dog) FF() string {
	return "肉骨头..."
}

func main() {
	a := new(Animal)
	a.name = "动物i"
	fmt.Println(a)
	fmt.Println(a.Call())
	fmt.Println(a.FF())
	fmt.Println(a.GetName())

	//实例化方式1
	d := new(Dog)
	d.name = "啸天"
	fmt.Println(d)
	fmt.Println(d.Call())
	fmt.Println(d.FF())
	fmt.Println(d.GetName())

	b := Animal{"京巴"}
	bd := Dog{b}
	fmt.Println(bd, bd.GetName(), bd.Call(), bd.FF())
}
