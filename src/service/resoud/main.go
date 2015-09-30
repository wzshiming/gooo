package main

import "flag"

func start() {
	//exec.Command("web").Start()
	//addr := flag.String("host", "www.mohegame.com:3007", "server host")
	addr := flag.String("host", "127.0.0.1:3007", "server host")
	port := flag.Uint("port", 3006, "listen port")
	flag.Parse()
	if *addr == "" {
		return
	}
	runWeb(*addr, *port)
}

func main() {
	start()
}
