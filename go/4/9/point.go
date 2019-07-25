package main

import "fmt"

func main() {
	var Num int32 = 5
	var intP *int32 = &Num

	fmt.Printf("整型%d的地址是%p\n", Num, &Num)
	fmt.Printf("整型%d的地址是%p\n", Num, intP)
	fmt.Printf("整型%d的地址是%p\n", *intP, intP)

}
