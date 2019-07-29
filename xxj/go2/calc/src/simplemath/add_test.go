package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 3)
	if r != 4 {
		//t.Error("Add(1,2) faild. Got ",r,",excepted 4.")
		t.Errorf("Add(1,2) faild. Got %d,excepted 4.", r)
	}

}
