// main.go
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/http"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

// ...

// CalculateHandler обрабатывает запрос для расчета факториала.
func calculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params, calculator calculate.Calculator) {
	// ...
}

// ...

func main() {
	router := http.NewRouter(calculator)
	router.POST("/calculate", calculateHandler)

	fmt.Println("Server is running on port 8989...")
	http.StartServer(router)
}
