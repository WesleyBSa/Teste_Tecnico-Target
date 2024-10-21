package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string

	fmt.Print("Digite uma ou mais palavras: ")

	// Utilizei o bufio.Scanner para ler a linha inteira, levando em consideração que no teste o usuario pode tentar entrar com uma frase ao invés de uma única palavra.
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	count := contarA(input)

	if count > 0 {
		fmt.Printf("A letra 'a' (maiúscula ou minúscula) aparece %d vez(es) na string.\n", count)
	} else {
		fmt.Println("A letra 'a' não está presente na string.")
	}
}

func contarA(s string) int {
	count := 0
	for _, char := range s {
		if char == 'a' || char == 'A' {
			count++
		}
	}
	return count
}
