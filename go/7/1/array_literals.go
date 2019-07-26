package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 20, 21, 22, 24}
	var arrLazy = [...]int{19, 20, 21, 22, 23}
	var arrLazySlice = []int{5, 6, 7, 8, 9}
	var arrKeyValue = [5]string{3: "zhangsan", 4: "lisi"}
	var arrKeyValueSlice = []string{3: "zhaosi", 4: "wangwu"}
	fmt.Println(arrAge)
	fmt.Println(arrLazy)
	fmt.Println(arrLazySlice)
	fmt.Println(arrKeyValue)
	fmt.Println(arrKeyValueSlice)

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

	fmt.Println("--------------------------------")
	var arrAgeB = [5]int{1, 2, 3, 4, 5}
	//前三个元素有值
	var arrAgeC = [10]int{1, 2, 3}
	fmt.Println(arrAgeB)
	fmt.Println(arrAgeC)
	var arrAgeD = [6]int{2: 9, 4: 16}
	fmt.Println(arrAgeD)
}
