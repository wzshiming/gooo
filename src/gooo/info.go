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

func MsgInfo(s *Configs, msg []byte) {
	//c0 := msg[0]
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	b1 := s.Rc[c1]
	b2 := b1.Map[c2]
	b3 := b2.Map[c3]
	log.Printf("%s(%d).%s(%d).%s(%d) <0x%08X^0x%08X>",
		b1.Name, c1,
		b2.Name, c2,
		b3, c3,
		b1.Allow|b2.Allow,
		b1.Unallow|b2.Unallow)
	log.Printf(string(msg[4:]))
}
