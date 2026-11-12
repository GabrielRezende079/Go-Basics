package easyleets

// Multiplicar somente os numeros Pares por 3 dentro de suas posições do Slice

import (
	"fmt"
	"math/rand/v2"
)

func SumIndex() {
	numeros := []int{}

	for i := 0; i < 10; i++ {
		n := rand.IntN(20)
		numeros = append(numeros, n)
	}
	fmt.Println(numeros)

	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			numeros[i] = numeros[i] * 3
		}
	}
	fmt.Println(numeros)
}
