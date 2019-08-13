package main

import (
	"fmt"
)

/**
默认情况下，channel 是同步 & 无缓冲的，在有接收者接收之前，发送不会结束。阻塞
	同步：
	无缓冲：
*/

func main() {
	ch2 := make(chan int)
	go pump(ch2)
	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)
	suck(ch2)
	//time.Sleep(1 * time.Second)

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func suck(ch chan int) {
	//for {
	fmt.Print(<-ch, " ")
	//}
}
