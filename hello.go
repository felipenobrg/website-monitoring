package main

import "fmt"

func main() {
	showIntro()
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Iniciando Monitoramento...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Não foi executado nenhum comando.")
	}
}

func showIntro() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	fmt.Println("O valor da variável comando é:", command)

	return command
}
