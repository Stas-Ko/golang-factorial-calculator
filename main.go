package main

import (
    "github.com/ваш_логин/github_репозиторий/pkg/http"
)

func main() {
    router := http.NewRouter()
    http.StartServer(router)
}
