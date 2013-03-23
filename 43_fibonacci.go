package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b, count := 0, 1, 0
	return func() int {
		if count == 0 {
			count++
			return a
		} else if count == 1 {
			count++
			return b
		}

		sum := a + b
		a = b
		b = sum
		return sum

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
