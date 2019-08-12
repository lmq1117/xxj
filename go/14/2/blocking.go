package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	go func() { out <- 2 }()
	go f1(out)
	time.Sleep(1 * time.Second)
}

func f1(in chan int) {
	fmt.Println(<-in)
}
