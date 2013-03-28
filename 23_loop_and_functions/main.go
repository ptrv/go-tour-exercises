package main

import (
	"fmt"
	"math"
)

func NewtonsSqrt(x float64) float64 {
	z := 1.0
	limit := 0.00001
	for {
		y := z
		z = y - ((y*y)-x)/(2*y)
		if math.Abs(y-z) < limit {
			break
		}
	}
	return z
}

func main() {
	s1 := math.Sqrt(2)
	s2 := NewtonsSqrt(2)
	fmt.Println(s1, s2, s2-s1)
}
