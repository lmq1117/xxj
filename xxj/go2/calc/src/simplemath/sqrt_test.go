package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	r := Sqrt(16)
	if r != 4 {
		t.Error("Sqrt(16) faild.Got ", r, ",excepted 3")
	}
}
