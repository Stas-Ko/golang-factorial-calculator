package http

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

type Result struct {
	AFactorial      string `json:"a_factorial"`
	BFactorial      string `json:"b_factorial"`
	AFactorialValue int    `json:"a_factorial_value"`
	BFactorialValue int    `json:"b_factorial_value"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var inputData InputData
	err := decoder.Decode(&inputData)
	if err != nil {
		http.Error(w, "Incorrect input", http.StatusBadRequest)
		return
	}

	if inputData.A < 0 || inputData.B < 0 {
		http.Error(w, "Incorrect input", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var aFactorial, bFactorial int

	go func() {
		defer wg.Done()
		aFactorial = calculate.CalculateFactorial(inputData.A)
	}()

	go func() {
		defer wg.Done()
		bFactorial = calculate.CalculateFactorial(inputData.B)
	}()

	wg.Wait()

	result := Result{
		AFactorial:      fmt.Sprintf("%d!", inputData.A),
		BFactorial:      fmt.Sprintf("%d!", inputData.B),
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

