package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", index)

	server := http.Server{
		Addr:    "127.0.0.1:1500",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := template.Must(template.ParseFiles("view/index.html"))

	t.Execute(w, "Hello, Template")
}
