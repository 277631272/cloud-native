package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":60000", nil)
	}()

	select {}
	fmt.Println("helloworld")
}
