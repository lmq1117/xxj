package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main开始", time.Now())
	go longWait()
	go shortWait()
	//longWait()
	//shortWait()
	fmt.Println("main sleep 开始", time.Now())
	time.Sleep(10 * time.Second)
	fmt.Println("main结束", time.Now())
}

func longWait() {
	fmt.Println("longWait 开始", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("longWait 结束", time.Now())
}

func shortWait() {
	fmt.Println("shortWait 开始", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println("shortWait 结束", time.Now())
}
