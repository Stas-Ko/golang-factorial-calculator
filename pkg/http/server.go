// pkg/http/server.go

package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/your-username/golang-factorial-calculator/pkg/calculate"
)

// NewServer создает новый HTTP-сервер с обработчиками маршрутов.
func NewServer() *http.Server {
	router := NewRouter()
	port := 8989
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return server
}

