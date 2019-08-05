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

	//if t, ok := areaIntf.(*Square); ok {
	//	fmt.Printf("变量 areaIntf 的类型是 is %T\n", t)
	//}
	//if u, ok := areaIntf.(*Circle); ok {
	//	fmt.Printf("变量 areaIntf 的类型是 is %T\n", u)
	//} else {
	//	fmt.Printf("变量 areaIntf 的类型不是Circle\n")
	//}

	//类型判断
	switch t := areaIntf.(type) {
	case *Circle:
	case *Square:
		fmt.Printf("变量 areaIntf 的类型是 is %T\n", t)
	case nil:
		fmt.Printf("nil value nothing to check?\n")
	default:
		fmt.Printf("unexcepted type %T\n", t)

	}

}
