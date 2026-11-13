package concurrencyPatterns

import (
	"fmt"
	"time"
)

// O padrão generator é basicamente a forma de encapsular canais
// para abstraí-los de funções principais. É o generator de Goroutines

// Função principal que instancia "escrever" e passa um valor ...
func ExecutarGenerator() {
	canal := escrever("Ola mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Mas toda a logica de Goroutines esta na propria função "escrever"
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recevido: %s", texto)
			time.Sleep(time.Microsecond * 500)
		}
	}()
	return canal
}
