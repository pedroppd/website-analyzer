package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const DELAY = 30 * time.Second
const FILE_PATH = "sites.txt"

func main() {
	showOptions()
	callMethod()
}

func startMonitoring() {
	sites := readFile(FILE_PATH)
	for {
		for _, site := range sites {
			fmt.Println("Starting monitoring: ", site)
			resp, err := http.Get(site)
			if err != nil {
				fmt.Println(site, "Error: ", err)
			} else {
				statusCode := resp.StatusCode
				if resp.StatusCode == 200 {
					registerLogger(site, statusCode)
					fmt.Println(site, "is up")
				} else {
					registerLogger(site, statusCode)
					fmt.Println(site, "Error Status: ", resp.StatusCode)
				}
			}
		}
		fmt.Println("----------------------------------------------------------------")
		time.Sleep(DELAY)
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

func readFile(filePath string) []string {
	var eventList []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	bufferValue := bufio.NewReader(file)
	for {
		fileLine, err := bufferValue.ReadString('\n')
		fileLine = strings.TrimSpace(fileLine)
		fmt.Println(fileLine)
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		eventList = append(eventList, fileLine)
	}
	file.Close()
	return eventList
}

func registerLogger(site string, status int) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(site + "-" + strconv.Itoa(status) + "\n")
	fmt.Println(file)
}
