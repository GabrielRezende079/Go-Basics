package main

import (
	"fmt"

	"github/GabrielRezende079/teste/internal/buscas"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
	"github/GabrielRezende079/teste/internal/erros"
	"github/GabrielRezende079/teste/internal/goroutines"
	"github/GabrielRezende079/teste/internal/grafos"
	"github/GabrielRezende079/teste/internal/interfaces"
	"github/GabrielRezende079/teste/internal/listaEncadeada"
	"github/GabrielRezende079/teste/internal/ponteiro"
	"github/GabrielRezende079/teste/internal/slices"
	"github/GabrielRezende079/teste/internal/tipos"
)

func main() {
	clearterminal.LimparTerminal()

	var escolha int

	fmt.Print("Escolha uma função:\n 1. tiposExecutar \n 2. Randomizador \n 3. Checador \n 4. errosExecutar \n 5." +
		" Ponteiro \n 6. Slices \n 7. Grafo \n 8. Palindromo \n 9. Lista Encadeada \n 10. Tamanho da Lista \n 11. Inset na Lista" +
		"\n 12. CRUD \n 13. Buscas em Slice \n 14. Interface \n 15. Gourotines \n")
	fmt.Scanln(&escolha)

	switch {
	case escolha == 1:
		tipos.Executar()

	case escolha == 2:
		tipos.Randomizador()

	case escolha == 3:
		tipos.Checador()

	case escolha == 4:
		erros.Executar()

	case escolha == 5:
		ponteiro.TestarPonteiros()

	case escolha == 6:
		slices.EscolherSlice()
	case escolha == 7:
		grafos.Executar()
	case escolha == 8:
		fmt.Println(tipos.VerificarPalindromo())
	case escolha == 9:
		listaEncadeada.NomeEncadeada()
	case escolha == 10:
		listaEncadeada.TamanhoLista()
	case escolha == 11:
		listaEncadeada.Convidados("Otavio")
	case escolha == 12:
		listaEncadeada.LinkedListCRUD()
	case escolha == 13:
		buscas.BuscarNumero()
	case escolha == 14:
		interfaces.ExecutarInterfaces()
	case escolha == 15:
		goroutines.ExecutarGourrotines()
	}
}
