package ponteiro

import (
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

// Sem o sinal de ponteiro '*' o valor da memoria é clonado e alterado, mas com ponteiro
// eu continuarei usando o valor original na memoria, sem copias

// Com ponteiro
func semAlterar(x *int) {
	*x = 5
	fmt.Print("\n Valor original de X: ", *x)
}

// Sem ponteiro
func Alterar(z int) {
	z = 5
	fmt.Print("\n Valor original de Z: ", z)
}

func TestarPonteiros() {
	clearterminal.LimparTerminal()

	y := 10

	Alterar(y)
	fmt.Print("\n Valor duplicado por Y: ", y)

	semAlterar(&y)
	fmt.Println("\n Valor de Y apontando para o endereço de X:", y)

}

// Ponteiros não são usados para imutabilidade, mas justamente o contrario
// Confere uma mudança no valor original se baseando em endereço de memoria
