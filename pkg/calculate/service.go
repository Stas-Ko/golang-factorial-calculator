// pkg/calculate/service.go
package calculate

import "sync"

// Calculator представляє собою сервіс для обчислення факторіала.
type Calculator struct {
	mu sync.Mutex
}

// NewCalculator створює новий екземпляр Calculator.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// CalculateFactorial обчислює факторіал числа n.
func (c *Calculator) CalculateFactorial(n int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n == 0 || n == 1 {
		return 1
	}
	return n * c.CalculateFactorial(n-1)
}
