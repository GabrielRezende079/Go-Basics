package goroutines

import (
	"fmt"
	"time"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

// #------------------------------#
// |   Explicação sobre Channels  |
// #------------------------------#
// Regra de ouro: Channels tem duas operações, enviar dados ou receber dados e podem estar abertos ou fechados

// A função que sera sincronização precisa receber o tipo canal, "chan", que é string nesse caso
func escrever3(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// Ao invés disso: fmt.Println(texto)
		canal <- texto // Declaramos assim, passando o texto pelo canal
		time.Sleep(time.Second)
	}
	close(canal) // Quando for imprimido 5x, então a função close vai fechar o canal
}

func ExecutarChannel() {
	clearterminal.LimparTerminal()

	canal := make(chan string)       // Declarando a Channel usando Make
	go escrever3("Olá Peter", canal) // A função vai executar em Goroutine passando o string e qual o seu canal

	for {
		mensagem, aberto := <-canal // Atribuimos o Canal para uma variavel para Imprimi-la, e uma Aberto para saber quando fechar a comunicação
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// Ou podemos aplicar um range no canal de uma forma mais delimitada:

	// for { mensagem := range canal {fmt.Println(mensagem)}}
}

// É IMPORTANTISSIMO fechar uma Channel, caso contrario pode causar DeadLock em compilação
// DeadLock: Channel aberte e sem dados para receber, e só é identificada em tempo de execução, não compilação
