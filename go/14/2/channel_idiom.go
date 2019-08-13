package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump2()
	//协程
	go suck2(stream)
	time.Sleep(time.Second * 1)
}

func pump2() chan int {
	ch := make(chan int)
	//匿名函数使用协程
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck2(ch chan int) {
	//使用循环从ch中取值
	//for {
	//	fmt.Println(<-ch)
	//}

	//给通道使用for循环
	for v := range ch {
		fmt.Println(v)
	}
}
