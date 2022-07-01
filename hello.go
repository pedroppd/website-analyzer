package main

import "fmt"

func main() {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
	var command int
	fmt.Scan(&command)

	switch command {
	case 1:
		fmt.Println("Case 01")
	case 2:
		fmt.Println("Case 2")
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default")
	}
}
