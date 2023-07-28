package main

import (
	"fmt"
	"net/http"

	"github.com/shamitsingh30/greedygame/pkg/handlers"
	"github.com/shamitsingh30/greedygame/pkg/models"
)

func main() {

	database := models.NewDatastore()
	fmt.Println(database.Data)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/check", handlers.ApiHandler(database))

	err := http.ListenAndServe(":8000", nil) // set listen port and handle
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Server started on :8000")
	}
}
