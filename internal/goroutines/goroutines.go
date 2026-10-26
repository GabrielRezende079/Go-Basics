package goroutines

import (
	"fmt"
	"time"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

var escolha int

// #------------------------------#
// |  Explicação sobre Goroutines |
// #------------------------------#

// Sem Goroutines, a linha "Tudo Bem?" nunca será executada, já que "Hello World" esta em um laço infinito

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func ExecutarGourrotines() {
	// Executar exemplos
	clearterminal.LimparTerminal()
	fmt.Print("\n 1. Exemplo Goroutine \n 2. Exemplo WaitGroup \n 3. Exemplo de Channels \n ")
	fmt.Scanln(&escolha)

	switch {
	case escolha == 1:
		// Exemplo de Goroutine basica
		go escrever("Hello World") // Mas aplicando a GOroutine, então dizemos para o computador executar sem esperar, concorrencia
		escrever("Tudo Bem?")
	case escolha == 2:
		ExecutarWaitGroup()
	case escolha == 3:
		ExecutarChannel()
	}
}
