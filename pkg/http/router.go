// pkg/http/router.go

package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

// NewRouter создает новый маршрутизатор HTTP с зарегистрированными обработчиками.
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/calculate", calculate.CalculateHandler)
	// Add more routes as needed

	return router
}
