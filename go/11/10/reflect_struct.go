package main

import (
	"fmt"
	"reflect"
)

type NotKnownType struct {
	s1, s2, s3, S4 string
}

func (n NotKnownType) String() string {
	return n.s1 + "----" + n.s2 + "----" + n.s3 + "----" + n.S4
}

var secret interface{} = NotKnownType{"ada", "go", "oberon", "haha"}

func main() {
	secretValue := reflect.ValueOf(secret)
	secretType := reflect.TypeOf(secret)
	fmt.Println(secret) //自动调用 NotKnownType.String
	fmt.Println(secretValue, secretType, secretValue.Kind())

	for i := 0; i < secretValue.NumField(); i++ {
		fmt.Printf("%d:%v\n", i, secretValue.Field(i))
	}
	//secretValue.Field(3).SetString("c#") //panic: reflect: reflect.Value.SetString using unaddressable value
	fmt.Println(secretValue.Method(0).Call(nil))

}
