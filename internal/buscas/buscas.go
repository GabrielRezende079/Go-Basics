package buscas

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

var (
	numeros          = []int{1, 5, 7, 0, 9, 8, 17}
	numerosOrdenados = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	escolha          int
)

func BuscarNumero() {
	fmt.Println("1. Busca Linear | 2. Busca Binaria")
	fmt.Scanln(&escolha)
	switch {
	case escolha == 1:
		buscaLinear()
	case escolha == 2:
		buscaBinaria()
	}
}

func buscaLinear() {
	clearterminal.LimparTerminal()
	target := 9

	for i, v := range numeros {
		if v == target {
			fmt.Println(target, "Foi encontrado na posição ", i)
		}
	}
}

func buscaBinaria() {
	clearterminal.LimparTerminal()
	inicio := 0
	fim := len(numerosOrdenados) - 1
	target := 16

	for inicio <= fim {
		meio := inicio + (fim-inicio)/2
		if target == meio {
			fmt.Println("Numero ", meio, " achado")
			break
		} else if meio < target {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
}
