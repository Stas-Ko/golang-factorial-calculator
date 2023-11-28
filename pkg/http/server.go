package http

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
)

// NewRouter создает новый роутер.
func NewRouter() *httprouter.Router {
    router := httprouter.New()
    router.POST("/calculate", Middleware(CalculateHandler))
    return router
}

// StartServer запускает сервер.
func StartServer(router *httprouter.Router) {
    fmt.Println("Server is running on port 8989...")
    http.ListenAndServe(":8989", router)
}
