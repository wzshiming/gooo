package gooo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Prevent() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP)
	for v := range ch {
		fmt.Println("Prevent", v, "Graceful shutdown please!")
	}
}
