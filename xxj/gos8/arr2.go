package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := [...]int{1, 2, 3}

	c := [5]int{1, 2, 3}

	d := [5]int{1: 11, 3: 44}

	arrLength := len(d)

	fmt.Println(a, b, c, d, arrLength)
}
