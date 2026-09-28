package erros

import (
	"fmt"
	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

func Executar() {

	clearterminal.LimparTerminal()

	var numero1 float64
	var numero2 float64

	fmt.Print("escolha os valores n1 e n2 para divisão: \n")

	fmt.Scan(&numero1)
	fmt.Scan(&numero2)

	resposta, err := divisao(numero1, numero2)

	if err != nil {
		fmt.Println("Erro ao dividir")
	} else {
		fmt.Printf("Resposta: %f \n", resposta)
	}

}

func divisao(n1 float64, n2 float64) (float64, error) {

	if n2 == 0 {
		return 0, fmt.Errorf("Não se faz divisão por 0 seu banana")
	}

	conta := n1 / n2

	// se não der erro vai retornar o resultado da conta e um nulo (nil) para parametro 'error'
	return conta, nil
}
