package listaEncadeada

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

// Para fazermos um insert verdadeiro, sem ficar recriando a lista,
// precisamos alterar o ponteiro para o começo da lista (Head), ou pro final da lista (Tail).
// Ao invés de alterar o ponteiro da Struct original, podemos fazer uma struct separada só para manipular Head ou Tail da lista.
// Atribuir outro ponteiro para o reciver do método, faz com que que chamou o metodo não conheça o novo ponteiro gerado.

// Conseguimos utlizar essa nova struct como novo receiver no método e alterar o ponteiro do campo
// Head e não do objeto inteiro.
type LinkedList struct {
	head *Pessoas
}

type Pessoas struct {
	name string
	next *Pessoas
}

func Convidados(name string) {
	clearterminal.LimparTerminal()

	convidados := &LinkedList{}

	convidados.insertAtBeginning("Robson")
	convidados.insertAtBeginning("Davi")
	convidados.printarConvidados()
	convidados.insertAtBeginning(name)
	convidados.insertAtEnd(name)
	convidados.printarConvidados()
}

// Insert Head
func (ll *LinkedList) insertAtBeginning(name string) {
	newPessoa := &Pessoas{
		name: name,
		next: ll.head,
	}
	ll.head = newPessoa
}

// Insert Tail
func (ll *LinkedList) insertAtEnd(name string) {
	newPessoa := &Pessoas{
		name: name,
		next: nil,
	}
	// Ele procura por nil porque significa que é o ultimo
	if ll.head == nil {
		ll.head = newPessoa
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newPessoa // Insert
}

// Printar
func (ll *LinkedList) printarConvidados() {
	current := ll.head
	for current != nil {
		fmt.Print(current.name, " --> ")
		current = current.next
	}
	fmt.Println("null")
}
