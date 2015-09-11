package main

import (
	"fmt"
	"net/http"
)

//func initVar() (dir, port string) {
//	inputDir := flag.String("d", "./static", "Static Directory")
//	inputPort := flag.Int("p", 8000, "Public Port")
//	flag.Parse()
//	dir = fmt.Sprintf("%s", *inputDir)
//	port = fmt.Sprintf(":%d", *inputPort)
//	return
//}

func filesServer(dir string, port int) (err error) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

//func run() (err error) {
//	dir, port := initVar()
//	fmt.Printf("Start the http files server\n\tfor dir:%s port%s\n", dir, port)
//	defer func() {
//		fmt.Printf("Stop the http files server\n")
//		if err != nil {
//			fmt.Printf("\tfor %s\n", err.Error())
//		}
//	}()
//	err = filesServer(dir, port)
//	return
//}
