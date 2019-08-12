package main

import (
	"fmt"
	"time"
)

//var ch1 chan string

func main() {
	//ch1 = make(chan string)
	//ch1 <- int1 //用ch通道发送int1
	//int2 := <- ch1 //变量int2 从ch1中接收数据
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	//sendData(ch)
	//getData(ch)
	time.Sleep(9 * time.Second)
}

func sendData(ch chan string) {
	ch <- "北京"
	ch <- "上海"
	ch <- "武汉"
	ch <- "深圳"
	ch <- "广州"
}
func getData(ch chan string) {
	var input string
	for {
		time.Sleep(1 * time.Second)
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
