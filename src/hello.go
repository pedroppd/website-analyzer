package main

import (
	"fmt"
	"net/http"
)

func main() {
	showOptions()
	callMethod()
}

func startMonitoring() {
	sites := []string{"https://www.google.com", "https://www.alura.com.br", "https://www.facebook.com", "https://www.twitter.com"}
	for _, site := range sites {
		fmt.Println("Starting monitoring: ", site)
		resp, err := http.Get(site)
		if err != nil {
			fmt.Println(site, "Error: ", err)
		} else {
			if resp.StatusCode == 200 {
				fmt.Println(site, "is up")
			} else {
				fmt.Println(site, "Error Status: ", resp.StatusCode)
			}
		}
	}
}

func callMethod() {
	var command int
	fmt.Scan(&command)

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Case 2")
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default")
	}
}

func showOptions() {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
}
