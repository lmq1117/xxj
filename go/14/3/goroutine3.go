package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	//getData(ch)
	getDataAuto(ch)
}

func sendData(ch chan string) {
	ch <- "北京"
	ch <- "上海"
	ch <- "广州"
	ch <- "深圳"
	ch <- "邓州"
	close(ch)
}

//手动检查通道是否关闭了 open
func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("%s ", input)
	}
}

func getDataAuto(ch chan string) {
	for input := range ch {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("%s ", input)
	}
}
