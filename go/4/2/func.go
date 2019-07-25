package main

import "fmt"

func main() {
	fmt.Println(LoveWho("lucy", "lily"))
	fmt.Println(LoveWho("mary", "jack"))
	fmt.Println(Sum(1, 5))
}

func LoveWho(first string, second string) string {
	return first
}

func Sum(a, b int) int {
	return a + b
}
