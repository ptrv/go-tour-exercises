package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for j := range a {
		b := make([]uint8, dx)
		for i := range b {
			z := (i + j) / 2
			b[i] = uint8(z)
		}
		a[j] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
