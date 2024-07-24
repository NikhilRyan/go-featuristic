package routes

import "net/http"

type Router interface {
	Handle(path string, handler http.Handler)
	Get(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
	Use(middleware func(http.Handler) http.Handler)
}
