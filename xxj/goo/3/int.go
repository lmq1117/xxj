package main

import "fmt"

type Integer int //对类型再定义
//type Integer = int //别名

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) LessThan(b Integer) bool {
	return a < b
}

func (a Integer) MoreThan(b Integer) bool {
	return a > b
}

func (a *Integer) Incr(i Integer) {
	*a = *a + i
}

func (a *Integer) Decr(i Integer) {
	*a = *a - i
}

func main() {
	var a Integer = 5
	var b Integer = 3
	if a.Equal(b) {
		fmt.Println("a eq b")
	} else if a.LessThan(b) {
		fmt.Println("a less than b")
	} else if a.MoreThan(b) {
		fmt.Println("a more than b")
	}
	a.Decr(2)
	fmt.Println(a)
	a.Incr(4) //Go底层自动转换成了(&a).Incr(4)
	fmt.Println(a)
	(&a).Incr(5)
	fmt.Println(a)
}
