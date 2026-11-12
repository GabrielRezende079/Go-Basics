package easyleets

import "fmt"

func InverterArryar() {
	numeros := [5]int{10, 20, 30, 40, 50}

	fmt.Println("E ordem", numeros)

	tamanho := len(numeros) - 1

	for i := tamanho; i >= 0; i-- {
		fmt.Println(numeros[i])
	}
}
