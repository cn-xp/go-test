package web

import (
	"io"
	"log"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit" />
		</form>
	</body></html>
`

func simpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello world.</h1>")
}

func formServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, r.FormValue("in"))
		//var b = 1
		//b --
		//var a = 1/b
		//a ++
	}
}

func WebSimpleServerMain() {
	http.HandleFunc("/test1", logPanics(simpleServer))
	http.HandleFunc("/test2", logPanics(formServer))
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}

func logPanics(function HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		function(w, r)
	}
}
