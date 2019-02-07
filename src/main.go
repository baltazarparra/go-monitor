package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		menu()

		switch userEntry() {
		case 1:
			monitoring()
		case 2:
			logging()
		case 3:
			entrySite()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço esse comando")
		}
	}

}

func menu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Monitorar outro site")
	fmt.Println("0 - Sair")
	fmt.Println("...")
}

func userEntry() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func monitoring() {
	fmt.Println("Monitorando...")

	sites := []string{"https://baltazarparra.github.io/error", "https://google.com", "https://twitter.com", "https://youtube.com"}

	for _, site := range sites {
		isUp(site)
	}
}

func logging() {
	fmt.Println("Exibindo logs...")
}

func entrySite() {
	fmt.Println("Digite o site que deseja monitorar...")
	var site string
	fmt.Scan(&site)
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println(site, "está de pé!")
	}

	if response.StatusCode != 200 {
		fmt.Println(site, "está com problemas para carregar, Status:", response.StatusCode)
	}
}

func isUp(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println(site, "está de pé!")
	}

	if response.StatusCode != 200 {
		fmt.Println(site, "não com problema, Status:", response.StatusCode)
	}
}
