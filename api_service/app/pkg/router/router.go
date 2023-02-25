package router

import (
	"net/http"
)

type Router struct {
	methodHandler  map[string]*http.ServeMux
	patternMethods map[string]map[string]bool
}

var (
	notFound         func(w http.ResponseWriter) = notFoundDefault
	methodNotAllowed func(w http.ResponseWriter) = methodNotAllowedDefault
)

func NewRouter() *Router {
	return &Router{
		methodHandler:  make(map[string]*http.ServeMux),
		patternMethods: make(map[string]map[string]bool),
	}
}

func (mux *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv, ok := mux.methodHandler[r.Method]
	if !ok {
		methodNotAllowed(w)
		return
	}

	handler, pattern := serv.Handler(r)
	if !mux.patternMethods[mux.longestMatch(r)][r.Method] {
		methodNotAllowed(w)
		return
	}

	if pattern != r.URL.Path {
		notFound(w)
		return
	}

	handler.ServeHTTP(w, r)
}

func (mux *Router) handler(method, pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	if _, ok := mux.methodHandler[method]; !ok {
		mux.methodHandler[method] = new(http.ServeMux)
	}
	if _, ok := mux.patternMethods[pattern]; !ok {
		mux.patternMethods[pattern] = map[string]bool{
			http.MethodGet:     false,
			http.MethodPost:    false,
			http.MethodPut:     false,
			http.MethodPatch:   false,
			http.MethodDelete:  false,
			http.MethodHead:    false,
			http.MethodConnect: false,
			http.MethodOptions: false,
			http.MethodTrace:   false,
		}
	}

	mux.patternMethods[pattern][method] = true
	mux.methodHandler[method].HandleFunc(pattern, f)
}

func (mux *Router) GET(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodGet, pattern, f)
}

func (mux *Router) POST(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodPost, pattern, f)
}

func (mux *Router) PUT(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodPut, pattern, f)
}

func (mux *Router) PATCH(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodPatch, pattern, f)
}

func (mux *Router) DELETE(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodDelete, pattern, f)
}

func (mux *Router) HEAD(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodHead, pattern, f)
}

func (mux *Router) CONNECT(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodConnect, pattern, f)
}

func (mux *Router) OPTIONS(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodOptions, pattern, f)
}

func (mux *Router) TRACE(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handler(http.MethodTrace, pattern, f)
}

func (mux *Router) NotFound(f func(w http.ResponseWriter)) {
	notFound = f
}

func (mux *Router) MethodNotAllowed(f func(w http.ResponseWriter)) {
	methodNotAllowed = f
}

func notFoundDefault(w http.ResponseWriter) {
	w.Write([]byte("not found\n"))
}

func methodNotAllowedDefault(w http.ResponseWriter) {
	w.Write([]byte("method not allowed\n"))
}

func (mux *Router) longestMatch(r *http.Request) (s string) {
	for _, serv := range mux.methodHandler {
		if _, pattern := serv.Handler(r); pattern > s {
			s = pattern
		}
	}
	return s
}
