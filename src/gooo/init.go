package gooo

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Name, Port, Type, Id = FlagVar()
var IP = GetIP()

func init() {
	//SetLogFile()
}

func SetLogFile() {
	if Name == "" {
		return
	}
	logfile, err := os.OpenFile(fmt.Sprintf("./log/%s_%d.log", Name, time.Now().Unix()), os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		return
	}
	//log.SetFlags(log.Ldate | log.Llongfile | log.Lmicroseconds | log.Ltime | log.LstdFlags)
	log.SetOutput(logfile)
}

func SetLogStd() {
	log.SetOutput(os.Stdout)
}
