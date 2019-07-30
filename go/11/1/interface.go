package main

import "fmt"

//接口Shaper 有一个方法Area
type Shaper interface {
	Area() float32
	//Perimeter() float32
}

//结构体Square
type Square struct {
	side float32
}

//func和方法名之间 所属类型声明|接收者声明
//方法名Area
func (sq *Square) Area() (area float32) {
	area = sq.side * sq.side
	return
}

//func (sq *Square) Perimeter() (z float32){
//	z = sq.side * 4
//	return
//}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	//var areaIntf = sq1
	areaIntf := Shaper(sq1)

	fmt.Printf("the square has area：%f\n", areaIntf.Area())
	//fmt.Printf("the square has area：%f\n", sq1.Area())
	//fmt.Printf("the square has Perimeter：%f\n", sq1.Perimeter())
}
