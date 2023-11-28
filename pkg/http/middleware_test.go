// pkg/http/server.go
package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

// NewRouter создает новый роутер.
func NewRouter(calculator calculate.Calculator) *httprouter.Router {
	router := httprouter.New()
	router.POST("/calculate", Middleware(calculateHandler(calculator)))
	return router
}
