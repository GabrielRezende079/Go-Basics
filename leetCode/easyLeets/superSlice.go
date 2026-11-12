package easyleets

import (
	"fmt"
	"math/rand/v2"
)

func SuperSlice() {
	// Gera um slice aleatori0
	comprimento := rand.IntN(11)
	numeros := make([]int, comprimento)

	for i := 0; i < comprimento; i++ {
		numeros[i] = rand.IntN(20)
	}
	fmt.Println("Slice aleatorio: ", numeros, "\n")
	// Validando slice vazio
	if len(numeros) == 0 {
		fmt.Println("Slice Vazio")
	} else {
		totalImpar := 0
		totalPar := 0
		somaTotal := 0
		maiorValor := numeros[0]
		menorValor := numeros[0]
		// Checando impares e pares
		for i := 0; i < len(numeros); i++ {
			if numeros[i]%2 != 0 {
				totalImpar++
			} else {
				totalPar++
			}
		}

		// Somando Total
		for _, valor := range numeros {
			somaTotal += valor
		}

		// Media
		var mediaValores float32 = float32(somaTotal) / float32(len(numeros))

		// Maior e menor valor
		for _, valor := range numeros {
			if valor >= maiorValor {
				maiorValor = valor
			}
		}

		for _, valor := range numeros {
			if valor <= menorValor {
				menorValor = valor
			}
		}

		fmt.Println("Informações Sobre o Slice:")
		fmt.Println("Tamanho do Slice: ", len(numeros))
		fmt.Println("Total de numeros impares: ", totalImpar)
		fmt.Println("Total de numeros pares: ", totalPar)
		fmt.Println("Maior valor: ", maiorValor)
		fmt.Println("Menor valor: ", menorValor)
		fmt.Println("A Soma de todos os numeros é de: ", somaTotal)
		fmt.Printf("A media dos numeros é de: %.2f \n", mediaValores)

	}
}
