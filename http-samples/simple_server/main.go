package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.Handle("/foo", MyHandler("Hello Wolrd!"))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type MyHandler string

func (handler MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler)
}
