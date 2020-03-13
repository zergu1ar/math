package math

// Calc is a function that return an operation function
func Calc(operation test.CalcOperation) func(a, b float64) float64 {
	return func(a, b float64) float64 {
		switch operation {
			case "+":
			return a + b
			case "-":
			return a - b
			case "*":
			return a * b
			case "/":
			return a/b
		}
		return 0
	}
}
