package main

import (
	"os"
	"os/signal"
	"rego"
	"syscall"
)

func over(funcs func()) {
	for {
		ch := make(chan os.Signal)

		signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP)
		<-ch
		funcs()
		rego.NOTICE("exit")
		os.Exit(1)
	}
}
