package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hello world v1, url: %#v\n", r.URL)
	fmt.Fprintf(w, "hello world v1")
}

type helloworld struct {
}
func NewHelloworld() http.Handler {
	return &helloworld{}
}
func (h *helloworld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hello world v2, url: %#v\n", r.URL)
	fmt.Fprintf(w, "hello world v2")
}

func main() {
	http.HandleFunc("/helloworld-v1", helloHandler)
	http.Handle("/helloworld-v2", NewHelloworld())
	http.ListenAndServe(":60000", nil)
}
