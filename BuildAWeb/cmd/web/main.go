package main

import (
	"fmt"
	"net/http"

	"github.com/tsawler/go-course/pkg/handlers"
)

const portNumber = ":8080"

//main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting applications on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
