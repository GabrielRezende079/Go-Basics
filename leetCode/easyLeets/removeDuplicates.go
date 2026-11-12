package easyleets

import "fmt"

func RemoveDuplicates() {
	numerosDuplicados := []int{4, 2, 4, 2, 7, 9, 7}
	numerosUnicos := []int{}

	// Vai passar por todos os numeros da Lista Original
	for i := 0; i < len(numerosDuplicados)-1; i++ {
		if numerosDuplicados[i] == numerosDuplicados[i+1] {
			numerosUnicos = append(numerosUnicos, numerosDuplicados[i])
		}
	}

	fmt.Println(numerosUnicos)
}

// {4, }

// Redes é legal

// VLSM

// Gateuai
