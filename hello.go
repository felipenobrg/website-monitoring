package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()

	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não foi executado nenhum comando.")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	fmt.Println("O valor da variável comando é:", command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://renovecasajp.com"}

	for i, site := range sites {
		fmt.Println("Passando na posição", i, ":", site)
		testSite(site)
	}
}

func showNames() {
	names := []string{"Douglas", "Maria", "João"}
	fmt.Println(names)
}

func testSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}
