package tipos

import (
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
	"math/rand"
)

func Checador() {

	clearterminal.LimparTerminal()

	numerosPar := []int{}
	numerosImpar := []int{}

	for i := 0; i < 10; i++ {
		numero := rand.Intn(10)

		if numero%2 == 0 {
			// fmt.Printf("Numero %d é PAR\n", numero)
			numerosPar = append(numerosPar, numero)

		} else {
			// fmt.Printf("Numero %d é IMPAR\n", numero)
			numerosImpar = append(numerosImpar, numero)
		}
	}

	fmt.Printf("Lista de Numeros Par: %d\n", numerosPar)
	fmt.Printf("Lista de Numeros Impar: %d\n", numerosImpar)

	quantidadeImpar := len(numerosImpar)
	quantidadePar := len(numerosPar)

	switch {
	case quantidadeImpar > quantidadePar:
		fmt.Print("Caiu mais Impar do que Par\n")

	case quantidadeImpar < quantidadePar:
		fmt.Print("Caiu mais Par do que Impar\n")

	case quantidadeImpar == quantidadePar:
		fmt.Print("Caiu quantidade iguais\n")
	}

	/*
		if quantidadeImpar > quantidadePar {
			fmt.Print("Caiu mais numeros Impar\n")
		} else {
			fmt.Print("Caiu main numero Par\n")
		}
	*/

}
