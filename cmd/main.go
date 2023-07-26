package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	err := http.ListenAndServe(":8000", nil) // set listen port and handle
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Server started on :8000")
	}
}
