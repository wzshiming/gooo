package gooo

import (
	"testing"
	"time"
)

func top_add(u *TopUse, s int) {
	for i := 0; i != s; i++ {
		u.Add()
	}
}

func top_sub(u *TopUse, s int) {
	for i := 0; i != s; i++ {
		u.Sub()
	}
}

func top_equ(u *TopUse, s int) bool {
	time.Sleep(time.Millisecond)
	return u.Get() == uint64(s)
}

func Test_topUse(t *testing.T) {
	u := NewTopUse()
	go u.Start()

	go top_add(u, 100)
	go top_sub(u, 99)
	if !top_equ(u, 1) {
		t.Error("step 1")
	}

	go top_sub(u, 100)
	if !top_equ(u, 0) {
		t.Error("step 2")
	}

	go top_add(u, 100)
	go top_sub(u, 100)
	go top_add(u, 100)
	go top_sub(u, 100)
	go top_add(u, 100)
	go top_sub(u, 100)
	go top_add(u, 100)
	go top_sub(u, 100)
	go top_add(u, 100)

	if !top_equ(u, 100) {
		t.Error("step 3")
	}

}
