package main

import "fmt"

func main() {
	//声明1
	var v1 int

	var (
		v2 int
		v3 string
		v4 bool
		v5 [10]int
		v6 struct {
			f float64
		}
		v7 *int
		v8 map[string]int
		v9 func(a int) int
	)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9)

}
