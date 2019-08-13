package main

import "fmt"

//值类型：int float bool string array
//引用类型：slice map channel
func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a
	b[2] = 7
	//fmt.Printf("%v",a)
	m(a)
	fmt.Println(a, b)
}

func m(a []int) {
	a[3] = a[3] * a[3]
}
