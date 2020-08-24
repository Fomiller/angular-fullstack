package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world!")
}
