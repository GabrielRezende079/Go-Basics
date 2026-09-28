package tipos

import(
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

func VerificarPalindromo() bool {

	clearterminal.LimparTerminal()


	word := "arara"
	fmt.Println("Palavra ", word, " é palindromo?")

	primeiro := 0
	ultimo := len(word) - 1

	for primeiro < ultimo {
		primeiraLetra := fmt.Sprintf("%c", word[primeiro])
		ultimaLetra := fmt.Sprintf("%c", word[ultimo])

		if primeiraLetra != ultimaLetra {
			return false
		}

		primeiro++
		ultimo--

	}
	return true
}
