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

func ApiHandler(db *models.Datastore, qb *models.Queuestore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestBody CommandRequest

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil || requestBody.Command == "" {
			resp["error"] = "invalid command"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}

		newReqBody, err := validation.ValidateFunc(&requestBody.Command)
		if err != nil {
			resp["error"] = err.Error()
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}

		querytype := newReqBody["querytype"]

		switch querytype {
		case "SET":
			controllers.Set_controller(&newReqBody, db)
		case "GET":
			key, val, err := controllers.Get_controller(&newReqBody, db)
			// fmt.Println(key, val, err)
			if err == nil {
				resp[key] = val
			} else {
				resp["error"] = err.Error()
			}
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)

		case "QPUSH":
			controllers.Push_controller(&newReqBody, qb)

		case "QPOP":
			x, err := controllers.Pop_controller(&newReqBody, qb)
			if err == nil {
				resp["value"] = x
			} else {
				resp["error"] = err.Error()
			}
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
		}
	}
}
