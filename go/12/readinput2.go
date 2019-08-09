package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input2      string
	err         error
)

func main() {
	//创建读取器，并与标准输入绑定
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please in put some words:")

	//返回读取到的字符串
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("the input was: %s\n", input2)
	}
}
