package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		fpp(&[3]int{i, i * i, i * i * i})
	}
}

func fpp(a *[3]int) { fmt.Println(a) }
