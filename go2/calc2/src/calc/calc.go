package main

import (
	"fmt"
	"strconv"
)

func main() {

	// string to int
	i, err := strconv.Atoi("445")
	if err != nil {
		fmt.Println("err found")
	}
	fmt.Println("i", i)

	// int to string
	vi := 123
	v2i := 756
	vs := strconv.Itoa(vi)
	v2s := strconv.Itoa(v2i)
	fmt.Println(vs + v2s)
}
