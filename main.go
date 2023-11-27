package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type InputData struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	AFactorial      string `json:"a_factorial"`
	BFactorial      string `json:"b_factorial"`
	BFactorialValue int    `json:"b_factorial_value"`
}

func calculateFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * calculateFactorial(n-1)
}

func calculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	wg.Add(1)

	var bFactorial int

	go func() {
		defer wg.Done()
		bFactorial = calculateFactorial(inputData.B)
	}()

	wg.Wait()

	result := Result{
		AFactorial:      fmt.Sprintf("%d!", inputData.A),
		BFactorial:      fmt.Sprintf("%d!", inputData.B),
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

func main() {
	router := httprouter.New()
	router.POST("/calculate", calculateHandler)

	fmt.Println("Server is running on port 8989...")
	http.ListenAndServe(":8989", router)
}
