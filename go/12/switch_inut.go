package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入您的姓名：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("error ")
		return
	}

	fmt.Printf("你输入的姓名是：%s\n", input)

	switch input {
	case "liming\r\n", "wangerxiao\r\n":
		fallthrough
	case "limin\r\n":
		fallthrough
	case "limi\r\n":
		fmt.Printf("welcome %s\n", input)
	default:
		fmt.Printf("you are not welcome here good bye!\n")
	}

}
