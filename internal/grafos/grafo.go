package grafos

import (
	"bufio"
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
	"os"
)

func CriarVertice(g *[]int, vertice int) {
	*g = append(*g, vertice)
}

func CriarAresta(g *[]int, aresta1, aresta2 int) {

}

func Executar() {

	// Limpa o Terminal
	clearterminal.LimparTerminal()

	// Instancia o Scanner
	scanner := bufio.NewReader(os.Stdin)

	// O Grafo como um Slice simples
	grafo := []int{}

	// Laço de repetição para coletar os vertices do usuario
	for i := 0; i < 4; i++ {
		// Varievel Vertice para o grafo
		var vertice int
		fmt.Println("Digite o Vertice: ")
		_, err := fmt.Fscan(scanner, &vertice)

		if err != nil {
			break
		}
		CriarVertice(&grafo, vertice)
	}

	// Laço de repetição para coletar as Arestas do usuario
	for i := 0; i < 4; i++ {
		var aresta1 int
		var aresta2 int
		fmt.Println("Digire a Aresta 1: ")
		fmt.Fscan(scanner, &aresta1)

		fmt.Println("Digire a Aresta 2: ")
		fmt.Fscan(scanner, &aresta2)

		CriarAresta(&grafo, aresta1, aresta2)
	}

	fmt.Println(grafo)
}
