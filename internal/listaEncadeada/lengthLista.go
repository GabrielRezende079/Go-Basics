package listaEncadeada

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

type Lista struct {
	data string
	next *Lista
}

func TamanhoLista() {
	clearterminal.LimparTerminal()

	lista := Lista{
		data: "Tofu",
		next: &Lista{
			data: "Banana",
			next: &Lista{
				data: "Arroz",
				next: nil,
			},
		},
	}

	lista.PrintarLista()
}

find

func (n *Lista) PrintarLista() {
	atual := n
	for atual != nil {
		fmt.Print(atual.data, " --> ")
		atual = atual.next
	}
	fmt.Print("null")
}
