package main

import (
	"github.com/Stas-Ko/golang-factorial-calculator/pkg/http"
)

func main() {
	router := http.NewRouter()
	http.StartServer(router)
}
