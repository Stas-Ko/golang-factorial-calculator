// pkg/http/middleware.go
package http

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

// Middleware перевіряє, чи a і b є невід'ємними цілими числами в JSON-запиті.
func Middleware(next httprouter.Handle, calculator *calculate.Calculator) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// ... ваш код middleware
	}
}
