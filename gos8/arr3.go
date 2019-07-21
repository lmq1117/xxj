package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	a1, a2 := arr[0], arr[len(arr)-1]

	fmt.Println(arr, a1, a2)

	for i := 0; i < len(arr); i++ {
		fmt.Println("Element ", i, "of arr is ", arr[i])
	}

	for i, v := range arr {
		fmt.Println("range i", i, "value", v)
	}
}
