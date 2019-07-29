package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

//垃圾回收
func main() {
	// 切片指向的数组，容量可能大于切片容量，只有在没有任何切片指向该数组的时候，底层数组的内存才会被释放，
	// 这种特性，可能导致占用多余的内存

	//s对应的底层数组是整个文件的数组，s不释放，则指向数组不释放
	s := FindDigits("aa.txt")
	fmt.Printf("%c%c%c\n", s[0], s[1], s[2])

}

func FindDigits(fileName string) []byte {
	b, _ := ioutil.ReadFile(fileName)
	return digitRegexp.Find(b)
}
