package main

import (
	"fmt"
)

func main() {
	// a) 1, 3, 5, 7, ___
	a := 7 + 2
	fmt.Printf("Próximo elemento da sequência a: %d\n", a)

	// b) 2, 4, 8, 16, 32, 64, ____
	b := 64 * 2
	fmt.Printf("Próximo elemento da sequência b: %d\n", b)

	// c) 0, 1, 4, 9, 16, 25, 36, ____
	c := 6 * 6 // 7^2
	fmt.Printf("Próximo elemento da sequência c: %d\n", c)

	// d) 4, 16, 36, 64, ____
	d := 10 * 10 // 10^2
	fmt.Printf("Próximo elemento da sequência d: %d\n", d)

	// e) 1, 1, 2, 3, 5, 8, ____
	e := 5 + 8
	fmt.Printf("Próximo elemento da sequência e: %d\n", e)

	// f) 2, 10, 12, 16, 17, 18, 19, ____
	f := 20
	fmt.Printf("Próximo elemento da sequência f: %d\n", f)
}
