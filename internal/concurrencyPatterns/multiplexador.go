package concurrencyPatterns

import (
	"fmt"
	"math/rand"
	"time"
)

func ExecutarMultiplexador() {
	canal := multiplexar(escreverr("Olá mundo"), escreverr("Eu estou encanado!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Essa função vai receber 2 canais e vai retornar como 1 só
// esse padrão é responsavel por unir canais, centralizando a comunicação de 2 canais em um só
func multiplexar(CanalEntrada1, CanalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-CanalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-CanalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escreverr(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recevido: %s", texto)
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
