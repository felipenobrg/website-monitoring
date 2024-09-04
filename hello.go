package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

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

	sites := readSites()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Passando na posição", i, ":", site)
			testSite(site)
		}
		time.Sleep(time.Second * delay)
		fmt.Println("")

	}

	fmt.Println("")
}

func showNames() {
	names := []string{"Douglas", "Maria", "João"}
	fmt.Println(names)
}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func readSites() []string {
	var site []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		site = append(site, line)
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}

	fmt.Println(site)
	return site
}
