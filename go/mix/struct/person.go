package main

import (
	"fmt"
	"strings"
)

type Person struct {
	fName string
	lName string
}

func UpPerson(p *Person) {
	p.fName = strings.ToUpper(p.fName)
	p.lName = strings.ToUpper(p.lName)
}

func main() {

	//value
	var p1 Person
	p1.fName = "Chris"
	p1.lName = "Woodward"
	UpPerson(&p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.lName)
	fmt.Println(p1)

	//point
	p2 := new(Person)
	p2.fName = "jack"
	p2.lName = "ma"

	fmt.Println(p2.fName)
	fmt.Println(p2.lName)
	fmt.Println(p2)
	UpPerson(p2)
	fmt.Println(p2.fName)
	fmt.Println(p2.lName)
	fmt.Println(p2)

	//字面量
	p3 := &Person{"tony", "ma"}
	fmt.Println("############")
	fmt.Println(p3.fName)
	fmt.Println(p3.lName)
	fmt.Println(p3)

	(*p3).lName = "mm"
	fmt.Println((*p3).fName)
	fmt.Println((*p3).lName)
	fmt.Println(*p3)

}
