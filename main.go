package main

import (
	"fmt"
	"net/http"

	"github.com/suyashmishra007/golang-web/pkg/handlers"
)

const portNumber string = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting application on port %s\n", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
