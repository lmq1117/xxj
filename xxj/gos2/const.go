package main

import "fmt"

func main() {
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0
	const (
		size int64 = 1024
		eof        = -1
	)

	const u, v float32 = 0, 2.1
	const a, b, c = 3, 4, "foo"
	const d, e int64 = 43, 45

	fmt.Println(Pi, zero, size, eof, u, v, a, b, c, d, e)

	//常量赋值是编译期行为 不能把运行期才能得出结果的表达式赋值给常量
	//const n = GetNum()

}

func GetNum() int {
	return 100
}
