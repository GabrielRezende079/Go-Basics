package arvores

import "fmt"

// Estrutura da Arvore
type Node struct {
	Valor    int
	Esquerda *Node
	Direita  *Node
}

// Inserir na Arvore binária
func (n *Node) inserir(valor int) {
	// Se valor passado for menor doque o Valor do nó, Esquerda...
	if valor < n.Valor {
		// Se na Esquerda estiver vazio...
		if n.Esquerda == nil {
			// valor passado agora é o Valor da Esquerda
			n.Esquerda = &Node{Valor: valor}
		} else {
			// Se não estiver vazio, chama a função novamente para inserir na Esquerda
			n.Esquerda.inserir(valor)
		}
	}

	// Se valor passado for menor doque Valor do nó, Direita...
	if valor > n.Valor {
		// Se na Direita estiver vazio...
		if n.Direita == nil {
			// valor passado agora é o Valor da Direita
			n.Direita = &Node{Valor: valor}
		} else {
			// Se não estiver vazio, chama a função novamente para inserir na Direita
			n.Direita.inserir(valor)
		}
	}
}

// Imprimir em ordem, da esquerda para direita
func (n *Node) emOrdem() {
	// Se Nó não estiver vazio
	if n != nil {
		// Chama a função novamente para imprimir a Esquerda
		n.Esquerda.emOrdem()
		// Imprime o Valor do Nó
		fmt.Printf("%d ", n.Valor)
		// Chama a função novamente para imprimir a Direita
		n.Direita.emOrdem()
	}
}

func ExecutarArvoreBinaria() {
	// Criando a Raiz da Arvore
	Raiz := &Node{Valor: 10}

	Raiz.inserir(21)
	Raiz.inserir(4)
	Raiz.inserir(80)
	Raiz.inserir(18)
	Raiz.inserir(6)

	Raiz.emOrdem()
}
