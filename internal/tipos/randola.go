package tipos

import (
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
	"math/rand"
)

func Randomizador() {

	clearterminal.LimparTerminal()

	for i := 0; i < 5; i++ {

		numeroAleatorio := rand.Intn(10)

		if numeroAleatorio > 5 {
			fmt.Printf("O Faustão morreu: %d\n", numeroAleatorio)
		} else {
			fmt.Printf("O Faustão viverá mais um dia: %d\n", numeroAleatorio)
		}

	}

}
