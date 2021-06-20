package main

import (
	demo "github.com/277631272/cloud-native/pkg/demo/http"
	"net/http"
)

func main() {
	name := "richardgu"
	server := http.Server{
		Addr: ":60000",
		Handler: demo.NewHelloWorld(name),
	}
	server.ListenAndServe()
}
