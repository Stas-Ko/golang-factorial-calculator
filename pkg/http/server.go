package http

import (
	"fmt"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
	"github.com/julienschmidt/httprouter"
)

// NewRouter создает новый роутер.
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/calculate", calculate.Middleware(CalculateHandler))
	return router
}

// StartServer запускает сервер.
func StartServer(router *httprouter.Router) {
	fmt.Println("Server is running on port 8989...")
	http.ListenAndServe(":8989", router)
}
