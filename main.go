package main

import (
	"fmt"
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/http"
	
)

func main() {
	server := http.NewServer(8989)
	fmt.Println("Server is running on port 8989...")
	server.Start()
}

