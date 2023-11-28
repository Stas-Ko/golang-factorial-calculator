package calculate

import (
	"encoding/json"
	"net/http"
	"sync"
)

// InputData - структура для представления данных в запросе.
type InputData struct {
	A int `json:"a"`
	B int `json:"b"`
}

// Middleware проверяет, что A и B являются неотрицательными целыми числами в JSON-запросе.
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

