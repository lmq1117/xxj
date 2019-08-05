package main

import (
	"fmt"
	"math"
)

//square 正方形
type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

//形状
type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("the type of areaIntf is %T\n", t)
	}

}