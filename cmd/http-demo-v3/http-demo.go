package main

import (
	"fmt"
	demo "github.com/277631272/cloud-native/pkg/demo/http"
	"net/http"
	"time"
)

func HelloWrold(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello richardgu"))
}

func timeFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		timeElapsed := time.Since(startTime)

		fmt.Printf("time elapsed: %d | %d\n", timeElapsed.Microseconds(), timeElapsed.Nanoseconds())
	})
}

func main() {
	r := demo.NewRouter()

	r.AddFilter(timeFilter)
	r.AddRoute("/", http.HandlerFunc(HelloWrold))
	r.Run(":60000")
}
