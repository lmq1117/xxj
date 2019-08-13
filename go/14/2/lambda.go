package main

import (
	"fmt"
	sort2 "sort"
	"time"
)

func main() {
	s := []int{7, 3, 6, 8, 2, 1, 9}
	fmt.Println(time.Now())
	fmt.Println(s)
	i := 3
	done := make(chan bool)
	doSort := func(s []int) {
		time.Sleep(time.Second * 2)
		sort(&s)
		done <- true
	}
	go doSort(s[:i])
	go doSort(s[i:])
	<-done //接收第一个信号
	<-done //接收第二个信号
	fmt.Println(time.Now())
	fmt.Println(s)
}

func sort(s *[]int) {
	sort2.Ints(*s)
}
