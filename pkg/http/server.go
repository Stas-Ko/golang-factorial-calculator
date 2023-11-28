package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// NewServer создает новый HTTP-сервер с обработчиками маршрутов.
func NewServer() *http.Server {
	router := httprouter.New()
	router.POST("/calculate", CalculateHandler)

	port := 8989
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return server
}
