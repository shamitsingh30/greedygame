package handlers

import (
	"encoding/json"
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

		var requestBody CommandRequest

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil || requestBody.Command == "" {
			// http.Error(w, "Invalid request body", http.StatusBadRequest)
			resp["error"] = "invalid command"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}

		newReqBody, err := validation.ValidateFunc(&requestBody.Command)
		if err != nil {
			// http.Error(w, "Invalid request body", http.StatusBadRequest)
			resp["error"] = err.Error()
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}

		if newReqBody["querytype"] == "SET" {
			controllers.Set_controller(&newReqBody, db)
		} else if newReqBody["querytype"] == "GET" {
			key, val, err := controllers.Get_controller(&newReqBody, db)

			if err == nil {
				resp[key] = val
			} else {
				resp["error"] = err.Error()
			}

			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}
	}
}
