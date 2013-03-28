package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0.0 {
		return 0, ErrNegativeSqrt(f)
	}
	z := 1.0
	limit := 0.00001
	for {
		y := z
		z = y - ((y*y)-f)/(2*y)
		if math.Abs(y-z) < limit {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
