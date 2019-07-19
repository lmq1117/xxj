package main

import (
	"fmt"
	"reflect"
)

func main() {
	//字符串切片
	//var content string = "abcdefghijklmnopqrstuvwxyz1234567890"
	//strOne := content[:5]
	//strTwo := content [7:]
	//strThree := content[6:9]
	//fmt.Println(content)
	//fmt.Println(strOne)
	//fmt.Println(strTwo)
	//fmt.Println(strThree)

	//字符串遍历
	str := "Hello, 世界"

	//11个字符
	//n := len(str)
	//for i := 0; i < n; i++ {
	//		ch:= str[i]
	//		fmt.Println(i,ch)
	//}

	//9个字符
	for i, ch := range str {
		//fmt.Println(i,ch)
		//fmt.Printf("%d----%d\n",i,ch)
		fmt.Println(reflect.TypeOf(i))
		fmt.Println(reflect.TypeOf(ch))
	}
}
