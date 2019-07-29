//main 包main函数是go程序执行的起点
package main

//不得导入程序中用不到的包 否则编译错误
import "fmt"

//main函数不能带参数 不能定义返回值
func main() {
	fmt.Println("hello world")
}

//直接运行 go run hello.go
//编译成hello.exe go build hello.go
//然后运行hello.exe
