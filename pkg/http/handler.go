// pkg/http/handler.go
package http

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

// Result представляє структуру для JSON-відповіді.
type Result struct {
	AFactorial      string `json:"a_factorial"`
	BFactorial      string `json:"b_factorial"`
	AFactorialValue int    `json:"a_factorial_value"`
	BFactorialValue int    `json:"b_factorial_value"`
}

// CalculateHandler обробляє запит для обчислення факторіала.
func CalculateHandler(calculator *calculate.Calculator) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// ... ваш код для обробки запиту
	}
}
