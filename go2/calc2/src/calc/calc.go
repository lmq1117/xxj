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
	fmt.Println(vs + v2s) //字串拼接

	//convert strings to values
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseBool("false")
	c, err := strconv.ParseFloat("3.1415", 64)
	d, err := strconv.ParseInt("-42", 10, 64)
	e, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, c, d, e)

	if b {
		fmt.Println("b is true")
	}
	if !f {
		fmt.Println("f is false")
	}

}
