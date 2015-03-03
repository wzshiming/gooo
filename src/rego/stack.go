package rego

import (
	"fmt"
	"runtime"
)

func PrintPanicStack() {
	if x := recover(); x != nil {
		PANIC(x)
		for i := 0; i < 10; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				PANIC(fmt.Sprintf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line))
			}
		}
	}
}
