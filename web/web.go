package web

import (
	"fmt"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside helloserver handle")
	fmt.Fprintf(w, "hello,"+req.URL.Path[1:])
}

func WebMain() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("error listenAndServe: ", err.Error())
	}
}
