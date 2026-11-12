package easyleets

import "fmt"

func ImparPar() {
	numeros := []int{9, 5, 2, 3, 6, 4, 1, 7, 8, 10}
	listaPar := []int{}
	soma := 0

	for _, valor := range numeros {
		if valor%2 == 0 {
			soma += valor
			listaPar = append(listaPar, valor)
		}
	}

	fmt.Println(listaPar)
	fmt.Println(soma)
}
