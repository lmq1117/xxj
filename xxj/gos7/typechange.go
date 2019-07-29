package main

import "fmt"

func main() {

	v1 := uint(16)
	v2 := int8(v1)
	v3 := uint16(v2)
	fmt.Println(v3)
}
