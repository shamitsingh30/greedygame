package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shamitsingh30/greedygame/pkg/controllers"
	"github.com/shamitsingh30/greedygame/pkg/models"
	"github.com/shamitsingh30/greedygame/pkg/validation"
)

type CommandRequest struct {
	Command string `json:"command"`
}

func ApiHandler(db *models.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Handler working")
		var requestBody CommandRequest

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

		if newReqBody["querytype"] == "SET" {
			controllers.Set_controller(&newReqBody, db)
		}
	}

}
