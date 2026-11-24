package main

import "fmt"

// Pow computes x raised to the integer power n using exponentiation by
// squaring. Negative exponents go through the reciprocal; the magnitude
// conversion is written to survive the most negative int as well.
func Pow(x float64, n int) float64 {
	m := uint64(n)
	if n < 0 {
		x = 1 / x
		m = uint64(-(int64(n) + 1)) + 1
	}
	res := 1.0
	for m > 0 {
		if m&1 == 1 {
			res *= x
		}
		x *= x
		m >>= 1
	}
	return res
}

func main() {
	fmt.Println(Pow(2, 10), Pow(2, -2), Pow(1.0001, 123456))
}
