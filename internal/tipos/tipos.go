package tipos

import (
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

func Executar() {

	clearterminal.LimparTerminal()

	var Resposta float64 = 2.00

	fmt.Printf("valor: %f\n", Resposta)

}
