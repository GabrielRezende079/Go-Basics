package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// #------------------------------#
// |  Explicação sobre Wait Group |
// #------------------------------#

// Wait Groups são uma das formas de sincronizar Goroutines, dizendo manualmente quantas são e onde estão

func escrever2(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func ExecutarWaitGroup() {
	var waitGroup sync.WaitGroup // Sync e um pacote padrão do GO
	waitGroup.Add(2)             // Esse comando dita quantas Goroutines esperar, 2

	// Função anonima para a Goroutine executar
	go func() {
		escrever2("Goroutine 1")
		waitGroup.Done() // Diz para o waitGroup que aqui acabou, "-1"
	}()

	go func() {
		escrever2("Goroutine 2")
		waitGroup.Done() // "-1"
	}()

	waitGroup.Wait() // ele espera as 2 Goroutines
}

// Wait groups são usadas apenas para saber quando um grupo de gorroutines terminou sem trocas de dados, apenas finish line
