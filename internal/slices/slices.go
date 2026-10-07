package slices

// Um Slice é montado em cima de um Array fixo e se torna flexivel

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

func EscolherSlice() {
	clearterminal.LimparTerminal()
	var escolha int

	lista0 := []int{3, 0, 1}

	fmt.Print("\n 1. Exemplo de Append \n 2. Exemplo de Indexação \n 3. Listar Impares \n 4. Inverter Lista \n 5. Mover zeros \n 6. teste \n 7. Numero faltando \n" +
		"8. Maior Numero \n")
	fmt.Scan(&escolha)

	switch {
	case escolha == 1:
		exemploSlice()
	case escolha == 2:
		perElementSlice()
	case escolha == 3:
		removerPares()
	case escolha == 4:
		invertendoSlice()
	case escolha == 5:
		moverZeros()
	case escolha == 6:
		fmt.Println(teste(lista0))
	case escolha == 7:
		fmt.Println(missingNumber(lista0))
	case escolha == 8:
		fmt.Println(maiorNumero(lista0))
	}
}

func exemploSlice() {
	clearterminal.LimparTerminal()

	lista1 := []int{1, 2, 3}

	fmt.Println("Lista original ", lista1)

	lista1 = append(lista1, 4, 5)

	fmt.Println("Lista agregada ", lista1)
}

func perElementSlice() {
	clearterminal.LimparTerminal()

	lista2 := []string{"Maçã", "Laranja", "Limão", "Tofu"}

	// i = Index &  v = Value
	for i, v := range lista2 {
		fmt.Printf("Indice: %d, Produto: %s\n", i, v)
	}
}

func removerPares() {
	clearterminal.LimparTerminal()

	lista3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	listaImpar := []int{}

	listaPar := []int{}

	fmt.Println("Lista original: ", lista3)

	for i := 0; i < len(lista3); i++ {
		if lista3[i]%2 != 0 {
			listaImpar = append(listaImpar, lista3[i])
		} else {
			listaPar = append(listaPar, lista3[i])
		}
	}
	fmt.Println("Lista impar: ", listaImpar)

	fmt.Println("Lista par: ", listaPar)
}

func invertendoSlice() {
	clearterminal.LimparTerminal()

	lista4 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(lista4)

	start := 0
	end := len(lista4) - 1

	for start < end {
		lista4[start], lista4[end] = lista4[end], lista4[start]
		start++
		end--
	}

	fmt.Println(lista4)
}

func moverZeros() {
	j := 0
	lista5 := []int{1, 0, 3, 2, 0}
	fmt.Println(lista5)

	for i := 0; i < len(lista5); i++ {
		if lista5[i] != 0 && lista5[j] == 0 {
			lista5[i], lista5[j] = lista5[j], lista5[i]
		}
		if lista5[j] != 0 {
			j++
		}
	}

	fmt.Println(lista5)
}

func teste(arr []int) []int {
	arr = append(arr, 1, 2, 3)

	return arr
}

func missingNumber(numeros []int) int {
	arrLenght := len(numeros)
	somaDosNumeros := arrLenght * (arrLenght + 1) / 2

	for _, value := range numeros {
		somaDosNumeros -= value
	}

	return somaDosNumeros
}

func maiorNumero(arr []int) int {
	maior := arr[0]
	for _, num := range arr {
		if num > maior {
			maior = num
		}
	}
	return maior
}
