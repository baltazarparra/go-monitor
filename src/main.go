package main

import "fmt"

func main() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
	fmt.Println("...")

	var comando int

	fmt.Scan(&comando)

	fmt.Println("Comando escolhido:", comando)
}
