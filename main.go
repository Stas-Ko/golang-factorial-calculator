package main

import (
	"fmt"
	"github.com/your-username/golang-factorial-calculator/pkg/http"
)

func main() {
	server := http.NewServer()

	fmt.Println("Server is running on port 8989...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
