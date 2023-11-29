package main

import (
	"fmt"
	"net/http"

	"github.com/Stas-Ko/golang-factorial-calculator/pkg/server"
)

func main() {
	router := server.NewRouter()
	router.AddRoutes()

	addr := ":8989"
	httpServer := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Printf("Server is running on %s...\n", addr)
	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
