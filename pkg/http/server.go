package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
	"net/http"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/calculate", calculate.Middleware(CalculateHandler))
	return router
}

func StartServer(router *httprouter.Router) {
	fmt.Println("Server is running on port 8989...")
	http.ListenAndServe(":8989", router)
}

func CalculateHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintln(w, "Hello, this is your calculate handler!")
}

