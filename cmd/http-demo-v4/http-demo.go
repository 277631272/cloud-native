package main

import (
	"fmt"
	"net/http"
)

func middleWare1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleWare1, hello world")
		next.ServeHTTP(w, r)
	})
}

func middleWare2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleWare2, hello world")
		next.ServeHTTP(w, r)
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", middleWare1(middleWare2(http.HandlerFunc(HelloWorld))))
	http.ListenAndServe(":60000", mux)
}
