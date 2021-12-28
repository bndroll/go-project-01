package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World !")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8000", nil)
}
