package main

import (
	"fmt"
)

func sum(x, y int, c chan int) {
	//time.Sleep(5 * time.Second)
	c <- x + y
}

func main() {
	c := make(chan int)
	go sum(12, 13, c)

	fmt.Println(<-c) //为了知道计算何时完成，可以通过信道回报//使用channel让main程序等待协程完成----信号量模式
}
