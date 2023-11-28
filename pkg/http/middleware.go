package http

import (
	"encoding/json" // Добавлено
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

// InputData добавлено
type InputData struct {
	A int `json:"a"`
	B int `json:"b"`
}

// Middleware перевіряє, чи a і b є не від'ємними цілими числами в JSON-запиті.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var inputData InputData
		err := decoder.Decode(&inputData)
		if err != nil || inputData.A < 0 || inputData.B < 0 {
			http.Error(w, "Incorrect input", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
