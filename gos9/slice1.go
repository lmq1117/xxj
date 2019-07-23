package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "january", 2: "february", 3: "march", 4: "april", 5: "may", 6: "june", 7: "july", 8: "auguest", 9: "september", 10: "october", 11: "november", 12: "december"}

	//创建slice1 基于数组
	q2 := months[3:6]     //第二季度
	summer := months[5:8] //夏季
	fmt.Println(months, "\n", q2, "\n", summer)
	halfUp := months[:7]
	halfDown := months[7:]

	//array[first:last] 基于slice
	all := months[:]
	fmt.Println(all, halfUp, halfDown)

	fmt.Println(len(q2), cap(q2))
	q1 := halfUp[:4]
	fmt.Println(q1)
	fmt.Println(months[0])
	fmt.Println(months[1])
	qianSanQ := halfUp[:10]
	fmt.Println(qianSanQ)
}
