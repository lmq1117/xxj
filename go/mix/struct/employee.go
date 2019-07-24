package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Id        int64
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerId int
	IsWait    bool
}

func main() {

	var dilbert Employee
	dilbert.Id = 1
	dilbert.Name = "李莉莉"
	dilbert.Address = "深圳市南山区兴海大道"
	dilbert.DoB = time.Now()
	dilbert.Position = "经理"
	dilbert.Salary = 50000
	dilbert.ManagerId = 1
	dilbert.IsWait = false

	fmt.Println(dilbert)
	fmt.Println(dilbert.DoB)
	fmt.Println(dilbert.IsWait)
	fmt.Println(dilbert.Position)

	//指针
	position := &dilbert.Position
	*position = "高级 " + *position
	fmt.Println(dilbert.Position)

	//指向结构体的指针
	var pDilbert *Employee = &dilbert
	pDilbert.Position = "资深 " + pDilbert.Position
	(*pDilbert).Position = "专家级 " + (*pDilbert).Position
	fmt.Println(dilbert.Position)

	//推导出类型
	p2 := &dilbert
	(*p2).Position = "古董 " + (*p2).Position
	fmt.Println(dilbert.Position)
}

//func EmployeeById(id int) *Employee {
//	return
//}
