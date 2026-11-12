package concurrencyPatterns

import (
	"fmt"

	clearterminal "github/GabrielRezende079/teste/internal/clearTerminal"
)

func Fibonacci(numero int) int {
	if numero == 0 {
		return 0
	} else if numero == 1 {
		return 1
	} else {
		return Fibonacci(numero-1) + Fibonacci(numero-2) // (Fn = Fn-1 + Fn-2 )
	}
}

func ExecutarFibonacci() {
	clearterminal.LimparTerminal()

	var numero int = 10

	fmt.Println(Fibonacci(numero))
}
