// Você tem um array de nums inteiros e um target inteiro.
// Retorne os indices de dois dos números que somados resultam no inteiro target
// Exemplo: INPUT: nums = [2,7,11,15], target = 9  | OUTPUT: [0,1]  (2+7 = 9)
// OBS: Cada input tem que ter uma solução!

package easyleets

import "fmt"

func TwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	current := 0

	for i, v := range nums {
		before := i - 1
		current += v

		if current == target {
			fmt.Println("receba ai")
			fmt.Println(current, " foi somado usando o indice", before, " + ", i)
			break
		}
	}
}
