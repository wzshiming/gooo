package gooo

import (
	"testing"
)

func Test_Used(t *testing.T) {
	u := NewUsed(9)
	for i := 0; i != 10; i++ {
		u.Join()
	}
	for i := 2; i != 7; i++ {
		if !u.IsUse(i) {
			t.Error("error")
		}
		u.Leave(i)
	}

	if u.IsUse(10) {
		t.Error("error")
	}
	t.Log(u.List(), "[0 1 7 8]")

	for i := 0; i != 2; i++ {
		u.Join()
	}

	t.Log(u.List(), "[0 1 2 3 7 8]")
}
