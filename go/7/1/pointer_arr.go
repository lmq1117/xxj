package main

import "fmt"

var arr4 [3]int

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func main() {
	f(arr4)
	fp(&arr4)
	arr4[1] = 7
	fmt.Println(arr4[1])
}
