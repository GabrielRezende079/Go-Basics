package concurrencyPatterns

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

// A função de Fibonacci calcula a soma dos dois numeros antecessores
// Ex: 0,1,1,2,3,5,8,13,21,34,55,...
func Fibonacci(numero uint) uint {
	if numero == 0 {
		return 0
	} else if numero == 1 {
		return 1
	} else {
		return Fibonacci(numero-1) + Fibonacci(numero-2) // (Fn = Fn-1 + Fn-2 )
	}
}

// Função worker para passar tudo em Goroutine, pegando os valores e executando de forma independente
func worker(tarefas <-chan uint, resultados chan<- uint) {
	// Vamos iterar cada numero no range de tarefas
	for numero := range tarefas {
		// e passar os valores para resultados
		resultados <- Fibonacci(numero)
	}
}

func ExecutarFibonacci() {
	clearterminal.LimparTerminal()

	// Declarando os canais
	tarefas := make(chan uint, 45)
	resultados := make(chan uint, 45)

	// Instanciando Workers com Goroutines, quanto mais, mais rápido
	// mas é limitado pelo poder do processador
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Enquanto i for menor que valor passado...
	for i := uint(0); i < 45; i++ {
		// Insere i No canal "tarefas"
		tarefas <- i
	}
	// Quando acabar fecha o canal
	close(tarefas)

	// For para printar os numeros
	for i := 0; i < 45; i++ {
		// do canal para a variavel
		resultados := <-resultados
		// printa as variaveis
		fmt.Println(resultados)
	}
}
