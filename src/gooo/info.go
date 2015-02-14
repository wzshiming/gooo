package gooo

import (
	"log"
	"runtime"
)

func Recover() {
	if err := recover(); err != nil {
		pc, file, line, ok := runtime.Caller(4)
		f := runtime.FuncForPC(pc).Name()
		if ok {
			log.Printf("Error: %v, File: %v, Line: %v, Func: %v.", err, file, line, f)
		} else {
			log.Printf("Error: %v", err)
		}
	}
}

func RecoverInfo() {
	pc, file, line, ok := runtime.Caller(4)
	f := runtime.FuncForPC(pc).Name()
	if ok {
		log.Printf("file: %v, line: %v, func: %v", file, line, f)
	}
}
