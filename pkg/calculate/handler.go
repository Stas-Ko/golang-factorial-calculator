package calculate

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)


// CalculateHandler обрабатывает запрос для расчета факториала.
func CalculateHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // ... ваш код для обработки запроса
}

