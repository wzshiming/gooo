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

func MsgInfo(msg []byte) {
	c0 := msg[0]
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	log.Println("flag:", c0)
	log.Println("type:", c1)
	log.Println("sub:", c2)
	log.Println("code:", c3)
	log.Println("info:", string(msg[4:]))
}
