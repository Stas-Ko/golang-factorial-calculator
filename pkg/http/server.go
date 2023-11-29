package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	router := httprouter.New()

	return &Router{
		Router: router,
	}
}

func (r *Router) AddRoutes() {
	calculateHandler := NewCalculateHandler(calculator.NewFactorialCalculator())

	r.POST("/calculate", calculateHandler.ServeHTTP)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Router.ServeHTTP(w, req)
}
