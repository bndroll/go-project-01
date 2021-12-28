package main

import (
	"fmt"
	"github.com/bndroll/go-project-01/pkg/handlers"
	"net/http"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
