package http

import "net/http"

type filter func(http.Handler) http.Handler

type Router struct {
	filterChain []filter
	mux         map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		filterChain: []filter{},
		mux:         make(map[string]http.Handler),
	}
}

func (r *Router) AddFilter(f filter) {
	r.filterChain = append(r.filterChain, f)
}

func (r *Router) AddRoute(route string, h http.Handler) {
	wh := h
	for i := 0; i < len(r.filterChain); i++ {
		wh = r.filterChain[i](wh)
	}
	r.mux[route] = wh
}

func (r *Router) Run(addr string) {
	for route, handler := range r.mux {
		http.Handle(route, handler)
	}
	http.ListenAndServe(addr, nil)
}
