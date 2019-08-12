package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50) //带缓冲 异步 啥时候取都行
	go func() {
		time.Sleep(15 * time.Second)
		x := <-c
		fmt.Println("从channel c中读得：", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent 10")
}
