package listaEncadeada

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

// Lista encadeada é um Objeto que se comporta como lista, sempre apontando para o proximo item
type Node struct {
	data string
	next *Node // Normalmente chamado de "next"  já que ele é usado apenas para apontar para o proximo
}

// A Lista encadeada ser para armazenar objetos com mais de uma caracteristica, se for só uma como
// no exemplo a baixo, não vale a pena, é melhor uma lista normal

func NomeEncadeada() {
	clearterminal.LimparTerminal()

	nodes := Node{
		data: "Gabriel",
		next: &Node{
			data: "Rezende",
			next: &Node{
				data: "Barbosa",
				next: &Node{
					data: "Florencio",
					next: nil,
				},
			},
		},
	}

	nodes.printarLista()
}

// Função que trabalha apontando para *Node unsando n e depois passando para "current"
func (n *Node) printarLista() {
	current := n
	for current != nil {
		fmt.Print(current.data, " --> ") // Printa 'data' do atual
		current = current.next           // current igual ao Next...
	}
	fmt.Print("null") // Quando acabar imprime vazio
}
