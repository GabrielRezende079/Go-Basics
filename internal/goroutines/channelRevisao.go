package goroutines

import "fmt"

func Principal() {
	canal := make(chan string) // 1. Constroi o channel sem buffer. Unbuffered Channel

	go func() {
		canal <- "Hello World" // 2. Manda a mensagem para o channel

		close(canal) // 3. Fecha a channel
	}()

	mensagem := <-canal // 4. A mensagem saí do Channel para uma variavel

	fmt.Println(mensagem) // 5. Printa a mensagem que passou pelo channel
}

// Unbuffered Channels são canais síncronos sem um buffer que diz a capacidade da channel.
// "Só consegue inserir no canal se tiver alguém para esperar e receber dados do outro lado"
