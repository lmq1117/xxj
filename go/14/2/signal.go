package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go compute(ch)
	fmt.Println("等待协程处理完成，waiting...")
	result := <-ch
	fmt.Println(result, "处理成功！")
}

func compute(ch chan bool) {
	time.Sleep(time.Second * 5)
	ch <- true
}
