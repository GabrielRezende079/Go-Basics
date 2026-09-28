package clearterminal

import (
	"os"
	"os/exec"
	"runtime"
)

// Limpa o terminal antes de escrever novas coisas
func LimparTerminal() {

	var cmd *exec.Cmd

	// Verifica o sistema operacional
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	// Define a saída do comando para o terminal
	cmd.Stdout = os.Stdout
	cmd.Run()

}
