package goroutines

import "fmt"

// #-----------------------------------#
// | Explicação sobre Bufferd Channels |
// #-----------------------------------#

// Um Buffer na channel torna ela como uma lista. O Channel se torna assíncrono, não sendo necessario alguém esperar
// "Remetente" -> "Channel", "Buffer" -> "Destinatario"
// Ele ganha uma fila de execução

func ExecutarBufferedChannel() {
	canal := make(chan string, 2) // Simplemente defina o tamanho do buffer com um numero, sem ele, nesse exmplo, daria DeadLock

	canal <- "Ola "
	canal <- "Mundo"

	palavra1 := <-canal
	palavra2 := <-canal

	fmt.Print(palavra1)
	fmt.Println(palavra2)
}
