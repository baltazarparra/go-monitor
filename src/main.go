package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
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
	fmt.Println("")
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

	fmt.Println("")
}

func registerLog(site string, status int) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	statusCode := strconv.Itoa(status)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Status: " + statusCode + "\n")

	file.Close()
}

func logging() {
	fmt.Println("Exibindo logs...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func entrySite() {
	fmt.Println("Digite o site que deseja monitorar...")
	var site string
	fmt.Scan(&site)
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println(site, "está de pé!")
		registerLog(site, response.StatusCode)
	}

	if response.StatusCode != 200 {
		fmt.Println(site, "está com problemas para carregar, Status:", response.StatusCode)
		registerLog(site, response.StatusCode)
	}
}

func isUp(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println(site, "está de pé!")
		registerLog(site, response.StatusCode)
	}

	if response.StatusCode != 200 {
		fmt.Println(site, "não com problema, Status:", response.StatusCode)
		registerLog(site, response.StatusCode)
	}
}
