package main

import (
	"flag"
	"fmt"
	"gooo"
	"log"
	"net/http"
)

func initVar() (name, dir, port string) {
	inputName := flag.String("n", "files", "Name")
	inputDir := flag.String("d", "./static", "Static Directory")
	inputPort := flag.Int("p", 8000, "Public Port")

	flag.Parse()
	name = fmt.Sprintf("%s", *inputName)
	dir = fmt.Sprintf("%s", *inputDir)
	port = fmt.Sprintf(":%d", *inputPort)
	return
}

func filesServer(dir, port string) (err error) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	return http.ListenAndServe(port, nil)
}

func run() (err error) {
	_, dir, port := initVar()
	log.Printf("%s Server starts from port %s\n", gooo.GetBaseName(), port)
	//fmt.Printf("Start the http files server\n\tfor dir:%s port%s\n", dir, port)
	defer func() {
		log.Printf("Stop the http files server\n")
		if err != nil {
			log.Printf("\tfor %s\n", err.Error())
		}
	}()
	err = filesServer(dir, port)
	return
}

func main() {
	run()
	return
}
