// pkg/http/middleware.go
package http

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

// Middleware перевіряє, чи a і b є невід'ємними цілими числами в JSON-запиті.
func Middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// ... ваш код middleware
	}
}
