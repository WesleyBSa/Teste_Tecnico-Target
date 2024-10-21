package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	if n < 0 {
		return []int{}
	}

	fib := []int{0, 1}
	for i := 2; i <= n; i++ {
		next := fib[i-1] + fib[i-2]
		fib = append(fib, next)
	}
	return fib
}

func pertenceAFibonacci(num int) bool {
	fibSeq := fibonacci(num)

	for _, v := range fibSeq {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	var numero int

	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	if pertenceAFibonacci(numero) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", numero)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", numero)
	}
}
