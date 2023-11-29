package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CalculateHandler struct {
	FactorialCalculator *calculator.FactorialCalculator
}

func NewCalculateHandler(factorialCalculator *calculator.FactorialCalculator) *CalculateHandler {
	return &CalculateHandler{
		FactorialCalculator: factorialCalculator,
	}
}

func (h *CalculateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var inputData struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "Incorrect input", http.StatusBadRequest)
		return
	}

	if inputData.A < 0 || inputData.B < 0 {
		http.Error(w, "Incorrect input", http.StatusBadRequest)
		return
	}

	aFactorial := h.FactorialCalculator.Calculate(inputData.A)
	bFactorial := h.FactorialCalculator.Calculate(inputData.B)

	result := struct {
		AFactorial      string `json:"a_factorial"`
		BFactorial      string `json:"b_factorial"`
		AFactorialValue int    `json:"a_factorial_value"`
		BFactorialValue int    `json:"b_factorial_value"`
	}{
		AFactorial:      strconv.Itoa(inputData.A) + "!",
		BFactorial:      strconv.Itoa(inputData.B) + "!",
		AFactorialValue: aFactorial,
		BFactorialValue: bFactorial,
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
