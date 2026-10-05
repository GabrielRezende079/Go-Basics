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

	lista.printarLista()
	fmt.Println("\n Tamanho: ", lista.tamanhoLista())
}

// Função Length
func (l *Lista) tamanhoLista() int {
	count := 0 // Contagem começa em zero
	current := l
	for current != nil {
		count++                // Adiciona em 'Count'
		current = current.next // Passa pro proximo
	}
	return count
}

func (n *Lista) printarLista() {
	atual := n
	for atual != nil {
		fmt.Print(atual.data, " --> ")
		atual = atual.next
	}
	fmt.Print("null")
}
