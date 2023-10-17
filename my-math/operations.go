package math

// ***Functions must be in Uppercase to be exported*** //
// It's recomended to put comments over exported functions

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Div(a int, b int) float64 {
	return float64(a) / float64(b)
}

func Mult(a int, b int) int {
	return a * b
}

// Functions in lowercase can't be exported
