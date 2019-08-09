package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	//从标准输入读取
	//fmt.Println("please input your full name: ")
	//fmt.Scanln(&firstName, &lastName)
	//fmt.Printf("Hi %s %s\n", firstName, lastName)

	//从标准输入读取
	//ss,_ := fmt.Scanf("%s %s\n", &firstName, &lastName)
	//fmt.Println(ss,firstName, lastName)

	//从字符串读取
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from the string we read: ", f, i, s)
}
