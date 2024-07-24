package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiRouter struct {
	chi.Router
}

func NewChiRouter() *ChiRouter {
	return &ChiRouter{chi.NewRouter()}
}

func (r *ChiRouter) Handle(path string, handler http.Handler) {
	r.Router.Handle(path, handler)
}

func (r *ChiRouter) Get(path string, handler http.HandlerFunc) {
	r.Router.Get(path, handler)
}

func (r *ChiRouter) Post(path string, handler http.HandlerFunc) {
	r.Router.Post(path, handler)
}

func (r *ChiRouter) Put(path string, handler http.HandlerFunc) {
	r.Router.Put(path, handler)
}

func (r *ChiRouter) Delete(path string, handler http.HandlerFunc) {
	r.Router.Delete(path, handler)
}

func (r *ChiRouter) Use(middleware func(http.Handler) http.Handler) {
	r.Router.Use(middleware)
}
