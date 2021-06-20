package http

import (
	"fmt"
	"net/http"
)

type HelloWrold interface {
	http.Handler
	GetName() string
}
type helloWorld struct {
	name string
}

func NewHelloWorld(name string) HelloWrold {
	return &helloWorld{name: name}
}

func (h *helloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hellowrold, name: %s, url: %v\n", h.name, r.URL)
	fmt.Fprintf(w,"helloworld, name: %s, url: %v\n", h.name, r.URL)
}

func (h *helloWorld) GetName() string {
	return h.name
}
