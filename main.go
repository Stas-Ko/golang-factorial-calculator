package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", CalculateHandler)

	fmt.Println("Server is running on port 8989...")
	http.ListenAndServe(":8989", router)
}

func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello, this is your calculate endpoint!")
}
