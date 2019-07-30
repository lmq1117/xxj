package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

//func NewStudent(id uint, name string, male bool, score float64) *Student {
//	return &Student{id: 1, name: "张三丰", male: true, score: 99.8}//初始化指定字段
//}

//增加接收者声明的方式定义了函数所属类型
func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

//toString 类似php魔术方法 当输出时自动调用
func (s Student) String() string {
	return fmt.Sprintf("------%d-----%s-----\n", s.id, s.name)
}

func main() {
	s := NewStudent(1, "张三丰", true, 107.3)
	s.SetName("张无忌")
	//fmt.Println(s.GetName())
	fmt.Println(s)
}
