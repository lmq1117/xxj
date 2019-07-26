package main

import "fmt"

var months = [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}

func main() {
	summer := months[3:6]
	fmt.Println(months)
	fmt.Println(summer)
	fmt.Println(len(summer))
	fmt.Println(cap(summer))
}
