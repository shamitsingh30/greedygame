package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shamitsingh30/greedygame/pkg/types"
	"github.com/shamitsingh30/greedygame/pkg/validation"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Handler working")
	var requestBody types.CommandRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Command == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newReqBody, check := validation.ValidateFunc(&requestBody.Command)

	fmt.Println(newReqBody)
	if !check {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
}
