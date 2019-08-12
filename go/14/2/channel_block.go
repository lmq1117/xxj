package main

import (
	"fmt"
	"time"
)

func main() {
	ch2 := make(chan int)
	go pump(ch2)
	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)
	go suck(ch2)
	time.Sleep(1 * time.Second)

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func suck(ch chan int) {
	for {
		fmt.Print(<-ch, " ")
	}
}
