package used

import (
	"testing"
)

func Test_Used(t *testing.T) {
	u := NewUsed(10)
	u.Join()
	u.Join()
	u.Join()
	u.Leave(0)
	t.Log(u.List())
	u.Join()
	u.Leave(2)
	t.Log(u.List())
	u.Join()
	u.Leave(1)
	t.Log(u.List())
}
