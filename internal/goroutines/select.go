package goroutines

import (
	"fmt"
	"time"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

// Select funciona como um Switch case das Goroutines

func ExecutarSelect() {
	clearterminal.LimparTerminal()

	// Construindo os Canais, sim, da pra fazer na mesma linha (Go é demais)
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// Aqui ele vai selecionar que estiver pronto independentemente do Loop passar pelos dois channels nesse exemplo
		select {
		// "Caso canal1 tenha passado o valor para 'mensagem'... "
		case mensagem1 := <-canal1:
			// printa mensagem
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}
}

// Explicando esse exemplo
// O For precisa recomeçar o ciclo para executar novamente, mas a channel um poderia ser printada até 4x dentro do
// tempo do loop normal. Com o Select, o for só se torna um laço de repetição e não mais o trigger dos channels
