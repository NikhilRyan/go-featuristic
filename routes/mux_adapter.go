package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouter struct {
	*mux.Router
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{mux.NewRouter()}
}

func (r *MuxRouter) Handle(path string, handler http.Handler) {
	r.Router.Handle(path, handler)
}

func (r *MuxRouter) Get(path string, handler http.HandlerFunc) {
	r.Router.HandleFunc(path, handler).Methods("GET")
}

func (r *MuxRouter) Post(path string, handler http.HandlerFunc) {
	r.Router.HandleFunc(path, handler).Methods("POST")
}

func (r *MuxRouter) Put(path string, handler http.HandlerFunc) {
	r.Router.HandleFunc(path, handler).Methods("PUT")
}

func (r *MuxRouter) Delete(path string, handler http.HandlerFunc) {
	r.Router.HandleFunc(path, handler).Methods("DELETE")
}

func (r *MuxRouter) Use(middleware func(http.Handler) http.Handler) {
	r.Router.Use(middleware)
}
