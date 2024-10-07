package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis{
		checkApi(api)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())
}

func checkApi(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("ERROR: ¡%s está caído! \n", api)
		return
	}

	fmt.Printf("SUCCESS: ¡%s está en funcionamiento \n", api)	
}