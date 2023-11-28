package calculate

// CalculateFactorial рассчитывает факториал.
func CalculateFactorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * CalculateFactorial(n-1)
}
