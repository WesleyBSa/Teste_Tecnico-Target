package main

import (
	"fmt"
)

func main() {
	INDICE := 12
	SOMA := 0
	K := 1

	for K < INDICE {
		K = K + 1
		SOMA = SOMA + K
	}

	fmt.Println(SOMA)
}
