package calculator

func CalculateFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * CalculateFactorial(n-1)
}

type FactorialCalculator struct{}

func NewFactorialCalculator() *FactorialCalculator {
	return &FactorialCalculator{}
}

func (c *FactorialCalculator) Calculate(n int) int {
	return CalculateFactorial(n)
}
