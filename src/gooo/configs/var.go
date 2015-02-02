package configs

import (
	"fmt"
	i18n "github.com/kortem/lingo"
	"gooo/addr"
	"gooo/helper"
	"log"
	"os"
	"time"
)

var Name, Port, Type, Id = helper.FlagVar()
var IP = addr.GetIP()
var I18n = i18n.New("zh_CN", "i18n")

func init() {
	if Name == "" {
		return
	}
	logfile, err := os.OpenFile(fmt.Sprintf("./log/%s_%d.log", Name, time.Now().Unix()), os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}
	//log.SetFlags(log.Ldate | log.Llongfile | log.Lmicroseconds | log.Ltime | log.LstdFlags)
	log.SetOutput(logfile)
}
