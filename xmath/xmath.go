package xmath

import "math"

func XMathRoundN(val float64, n int) float64 {
	pow := math.Pow10(n)
	return math.Round(val*pow) / pow
}

func XMathFoolN(val float64, n int) float64 {
	pow := math.Pow10(n)
	return math.Floor(val*pow) / pow
}

func XMathCeilN(val float64, n int) float64 {
	pow := math.Pow10(n)
	return math.Ceil(val*pow) / pow
}
