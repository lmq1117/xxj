package main

//导入方式3 分别导入 但写在一行
import "fmt"
import "os"

func main() {
	fmt.Println(os.Args[0])
}
