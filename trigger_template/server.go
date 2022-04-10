package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./tmpl.html")
	t.Execute(w, "Hello World!")
}

func first(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/", first)
	server.ListenAndServe()
}
