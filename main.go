// main.go
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/http"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/calculate"
)

func main() {
	calculator := calculate.NewCalculator()
	router := http.NewRouter(calculator)
	router.POST("/calculate", http.Middleware(http.CalculateHandler(calculator)))

	fmt.Println("Server is running on port 8989...")
	http.StartServer(router)
}

