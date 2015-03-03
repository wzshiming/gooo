package rego

import (
	"time"
)

var x0 uint32 = uint32(time.Now().UnixNano())
var a uint32 = 1664525
var c uint32 = 1013904223

var LCG chan uint32

func init() {
	LCG = make(chan uint32, 1024)
	go func() {
		for {
			x0 = a*x0 + c
			LCG <- x0
		}
	}()
}
