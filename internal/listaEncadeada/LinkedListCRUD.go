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

	Compra := &linkedList{}

	var escolha int
	for {
		fmt.Println("1. listar lista | 2. Inserir no Inicio | 3. Inserir no Final | 4. Remover do Inicio | 5. Remover do Final")
		switch {
		case escolha == 1:
			Compra.printarLista()
		case escolha == 2:
			// Do Something
		case escolha == 3:
			// Do Something
		case escolha == 4:
			// Do Something
		case escolha == 5:
			// Do Something
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
}

// Remover do Inicio
func (ll *linkedList) removeInicio() {
}

// Remover do Final
func (ll *linkedList) removeFinal() {
}
