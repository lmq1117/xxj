package main

import "fmt"

type struct1 struct {
	age   int
	money float32
	work  string
}

func main() {
	ms := new(struct1)
	ms.age = 15
	ms.money = 100
	ms.work = "student"

	fmt.Println("the age is ", ms.age)
	fmt.Println("the money is ", ms.money)
	fmt.Println("the work is ", ms.work)
	fmt.Println(ms)

	ms2 := &struct1{99, 1000000.01, "old"}
	fmt.Println(ms2)
}
