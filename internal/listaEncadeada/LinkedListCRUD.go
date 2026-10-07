package listaEncadeada

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

type linkedList struct {
	head *produto
}

type produto struct {
	name string
	next *produto
}

func LinkedListCRUD() {
	clearterminal.LimparTerminal()

	Compra := &linkedList{
		head: &produto{
			name: "Davi",
			next: &produto{
				name: "João",
				next: nil,
			},
		},
	}

	var escolha int
	var nome string
	for {
		fmt.Println("1. listar lista | 2. Inserir no Inicio | 3. Inserir no Final | 4. Remover do Inicio | 5. Remover do Final")
		fmt.Scanln(&escolha)
		switch {
		case escolha == 1:
			Compra.printarLista()
		case escolha == 2:
			fmt.Println("Digite um Nome")
			fmt.Scanln(&nome)
			Compra.insertInicio(nome) // <--- Editar aqui ainda
		case escolha == 3:
			fmt.Println("Digite um Nome")
			fmt.Scanln(&nome)
			Compra.inserFinal(nome)
		case escolha == 4:
			Compra.removeInicio()
		case escolha == 5:
			Compra.removeFinal()
		}
	}
}

// Listar
func (ll *linkedList) printarLista() {
	current := ll.head
	for current != nil {
		fmt.Print(current.name, " --> ")
		current = current.next
	}
	fmt.Println("null")
}

// Inserir no Inicio
func (ll *linkedList) insertInicio(name string) {
	newProduto := &produto{
		name: name,
		next: ll.head,
	}
	ll.head = newProduto
	ll.printarLista()
}

// Inserir no Final
func (ll *linkedList) inserFinal(name string) {
	newProduto := &produto{
		name: name,
		next: nil,
	}
	// Se estiver vazia vai inserir no começo
	if ll.head == nil {
		ll.head = newProduto
		return
	}
	current := ll.head
	// procura pela posição que aponta pro vazio
	for current.next != nil {
		current = current.next
	}
	current.next = newProduto // insert
	ll.printarLista()
}

// Remover do Inicio
func (ll *linkedList) removeInicio() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
	ll.printarLista()
}

// Remover do Final
func (ll *linkedList) removeFinal() {
	if ll.head == nil || ll.head.next == nil {
		return
	}
	previous := ll.head
	for ll.head.next.next != nil {
		previous = previous.next
	}
	previous.next = nil
	ll.printarLista()
}
