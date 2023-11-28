// pkg/calculate/service.go
package calculate

// Calculator определяет интерфейс для вычисления факториала.
type Calculator interface {
	CalculateFactorial(n int) int
}

// FactorialService реализует вычисление факториала.
type FactorialService struct{}

// CalculateFactorial вычисляет факториал числа n.
func (s *FactorialService) CalculateFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * s.CalculateFactorial(n-1)
}
