package main

import "fmt"

const (
	WIDTH  = 16
	HEIGIH = 10
)

type pixel int

var screen [WIDTH][HEIGIH]pixel

func main() {
	for y := 0; y < HEIGIH; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 1
		}
	}
	fmt.Println(screen)
}
