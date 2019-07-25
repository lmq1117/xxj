package main

//更短形式，但使用gofmt后会被强制换行
import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(os.Args[0])
}

//当导入多个包时 最好按字母顺序排列报名 更清晰易读

//如果包名 不是以 . | / | "container/list"  开头,则Go会在全局文件进行查找；
// ./ GO在相对目录查找
// / 在系统绝对路径中查找

// _

//可见性顾泽 Pi name

//包别名 解决包名冲突 import fm "fmt"

//导入包但没使用 构建程序时报错 build？

//包的分级声明和初始化
//
