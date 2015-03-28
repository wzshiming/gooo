package test

import (
	"testing"
	"ygo"
	"ygo/define"
)

func Test_play(t *testing.T) {
	y := ygo.NewYGO(define.DefaultTest(nil, 1), define.DefaultTest(nil, 2))
	t.Log(y)
	y.Loop()
}
