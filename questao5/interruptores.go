package main

import (
	"fmt"
	"time"
)

func main() {
	interruptores := []string{"I1", "I2", "I3"}
	lampadas := make(map[string]string)

	lampadas["I1"] = "quente"
	time.Sleep(10 * time.Second)

	lampadas["I1"] = "desligada"
	lampadas["I2"] = "acesa"

	lampadas["I3"] = "desligada"

	fmt.Println("Resultados após a análise das lâmpadas:")
	for _, interruptor := range interruptores {
		fmt.Printf("Interruptor %s controla a lâmpada %s\n", interruptor, lampadas[interruptor])
	}

	validarPossibilidades(lampadas)
}

func validarPossibilidades(lampadas map[string]string) {
	fmt.Println("\nValidação das possibilidades:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				fmt.Printf("Se a lâmpada %s é controlada pelo interruptor I%d e a lâmpada %s pelo interruptor I%d\n",
					lampadas["I1"], 1, lampadas["I2"], 2)
			}
		}
	}
}
