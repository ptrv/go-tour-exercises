package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
	z := 1.0 + 0i
	limit := 0.0000000000001
	for {
		y := z
		z = y - (((y * y * y) - x) / (3 * (y * y)))
		d := y - z
		if cmplx.Abs(d) < limit {
			break
		}
	}
	return z
}

func main() {
	cr1 := cmplx.Pow(2.0+3i, 1.0/3.0)
	cr2 := Cbrt(2 + 3i)
	fmt.Println("cmplx.Pow: ", cr1)
	fmt.Println("Cbrt: ", cr2)
	fmt.Println("Diff: ", cr2-cr1)
}
