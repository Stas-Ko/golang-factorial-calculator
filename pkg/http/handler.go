package http

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Handler handles HTTP requests.
type Handler struct {
	// Add any dependencies or services here
}

// NewHandler creates a new instance of Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// CalculateHandler handles the /calculate endpoint.
func (h *Handler) CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		aFactorial = calculateFactorial(inputData.A)
	}()

	go func() {
		defer wg.Done()
		bFactorial = calculateFactorial(inputData.B)
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
