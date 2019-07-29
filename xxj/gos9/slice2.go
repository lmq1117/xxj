package main

import "fmt"

func main() {
	months := [...]string{1: "january", 2: "february", 3: "march", 4: "april", 5: "may", 6: "june", 7: "july", 8: "auguest", 9: "september", 10: "october", 11: "november", 12: "december"}
	fmt.Println(len(months))

	slice1 := make([]int, 5)
	fmt.Println(slice1)
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)
	slice2[5] = 5
	fmt.Println(slice2)
}
