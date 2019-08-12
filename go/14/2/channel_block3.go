package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //不带缓冲 同步
	go func() {
		time.Sleep(15 * time.Second)
		x := <-c
		fmt.Println("从channel中接收到了", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
